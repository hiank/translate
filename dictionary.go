package translate

import (

	"os"
	"fmt"

	"encoding/xml"
	"io/ioutil"
	"strconv"
	"bufio"
	"strings"
	"hiank.net/translate/tool"

)


// type Item struct {
// 	Key 	int		`xml:"key,attr"`
// 	Desc 	string	`xml:"desc,attr"`	// simplified chinese
// 	Value 	string 	`xml:"value"`

// }

type Data struct {

	XMLName xml.Name    `xml:"Data"`
	Item 	[]tool.Item		`xml:"Item"`
}


type Dict struct {

	_dictFile string   // dictionary name
	_dict     map[string]*tool.Item


}


func NewDict(dictName string) *Dict {

	d := new(Dict)
	d.reset(dictName)

	return d
}

func (d *Dict) reset(dictName string) {

	*d = Dict {
		_dictFile: 	"default.plist",
		_dict:		make(map[string]*tool.Item),

	}

	d.InitDict(dictName)
}

func (d *Dict) InitDict(xmlName string) {

	d._dictFile = xmlName

	content, err := ioutil.ReadFile(xmlName)
	if err != nil {
		fmt.Println("open xml file error : " + err.Error())
	}

	var data Data
	err = xml.Unmarshal(content, &data)
	for _, i := range data.Item {

		item := new(tool.Item)
		*item = i
		d._dict[i.Desc] = item

//		fmt.Printf("__initDict : %v\n", i.Desc, d._dict[i.Desc])
	}

}


func (d *Dict) Flush() {

	d.format()

	data := Data {
		Item: 	make([]tool.Item, 0),
	}
	for _, value := range d._dict {
//		fmt.Printf("why cann't work %v \n", value)
		data.Item = append(data.Item, *value)
	}

	file, err := os.OpenFile(d._dictFile, os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("cann't write the result to " + d._dictFile + " : " + err.Error())
	}
	defer file.Close()



	out, err := xml.MarshalIndent(data, "", "	")
	if err != nil {
		fmt.Println("error : " + err.Error())
	}

	w := bufio.NewWriter(file)
	w.Write([]byte(xml.Header))
	w.Write(out)

	w.Flush()


	file, e := os.OpenFile(strings.Replace(d._dictFile, ".xml", "l.info", -1), os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
	if e != nil {
		fmt.Println("cann't write the result to dic.info, it cann't be opened " + e.Error())
		return
	}
	defer file.Close()

	bufW := bufio.NewWriter(file)
	content := ""
	for _, value := range d._dict {

		desc := strings.Replace(value.Desc, "\x0A", "&#xA;", -1)

		content += strconv.Itoa(value.Key) + "," + desc + "\n"
//		bufW.WriteString(strconv.Itoa(value.Key) + "," + desc + "\n")
//		bufW.Flush()
	}
	bufW.WriteString(content)
	bufW.Flush()
}

func (d *Dict) format() {

	maxKey := -1

//	fmt.Printf("dict format len : %v, %v\n", d._dict["1"], d._dict["2"])

	dictEx := NewDictEx(strings.Replace(d._dictFile, ".xml", "r.info", -1))
	fmt.Printf("dict len : %v\n", len(d._dict))
	for _, value := range d._dict {


		if value != nil {

			if maxKey < value.Key {
				maxKey = value.Key
			}

			str := dictEx._dict[value.Key]
			if str != "" {
				value.Value = dictEx._dict[value.Key]
			}
		}
	}

	for key, value := range d._dict {

		if value == nil {

			maxKey++

			value = new(tool.Item)
			*value = tool.Item {
				Key: 	maxKey,
				Desc: 	key,

				Value: 	dictEx._dict[maxKey],
			}
			d._dict[key] = value
		}
	}
}





