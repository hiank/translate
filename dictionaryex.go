package translate

import (

	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
	"bytes"
	"strconv"
)


type DictEx struct {

//	_dictFile string   // dictionary name
	_dict     map[int]string


}


func NewDictEx(dictName string) *DictEx {

	d := new(DictEx)
	d.reset(dictName)

	return d
}

func (d *DictEx) reset(dictName string) {

	*d = DictEx {
//		_dictFile: 	"default.plist",
		_dict:		make(map[int]string),

	}

	d.InitDictEx(dictName)
}

func (d *DictEx) InitDictEx(name string) error {

//	d._dictFile = xmlName
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("when initDictEx : " + name + err.Error())
		return err
	}

	defer file.Close()

	r := bufio.NewReader(file)

	loop:	for {

		lineByte, err := r.ReadBytes('\n')
		switch err {
		case nil:
		case io.EOF:
			//			fmt.Println("end of file")
			break loop
		default:
			fmt.Println("there's an error when read line : " + err.Error())
			break loop
		}
//		frame.pushLine(lineByte)
		d.readItem(lineByte)
	}

	return nil
}

func (d *DictEx) readItem(lineByte []byte) {

//	lineStr := string(lineByte)
	idx := bytes.IndexByte(lineByte, ',')
	key, err := strconv.Atoi(string(lineByte[0:idx]))
	if err != nil {
		fmt.Printf("why error %v...%v", lineByte[0:idx], err.Error())
	}
	value := string(lineByte[idx+1:len(lineByte)])

	value = strings.TrimSuffix(value, "\x0A")
	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r\n")

	if d._dict[key] != "" {
		fmt.Print("why why : %v", key)

	}
	d._dict[key] = strings.Replace(value, "&#xA;", "\x0A", -1)

//	fmt.Printf("__ key : %v, __ value : %v\n", key, value)
}



