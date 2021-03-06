package dict

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"

	"hiank.net/translate/core"
	"hiank.net/translate/tool"
)

// RoutXML realize FileFilter and Routine interface
type RoutXML struct {
	Dict core.Dict

	// T *tool.Data
}

// Match is the function realize for FileFilter interface
func (r *RoutXML) Match(name string) bool {

	rlt := false
	switch {

	case strings.HasSuffix(name, ".xml"):
		rlt = true
	}

	return rlt
}

// Filter is the function realize for Routine interface
func (r *RoutXML) Filter() core.Filter {

	return r
}

// Run is the function realize for Routine interface
func (r *RoutXML) Run(path string) core.RoutineChan {

	content, err := ioutil.ReadFile(path)
	if err != nil {

		fmt.Println("open xml file error : " + err.Error())
		return r
	}

	data := new(tool.Data)
	err = xml.Unmarshal(content, data)
	if err != nil {
		fmt.Println("format xml content error : " + err.Error())
		return r
	}

	// r.T = data
	return data
}

// End is the function realize for Routine interface
func (r *RoutXML) End(ch core.RoutineChan) {

	var data *tool.Data
	var ok bool
	if data, ok = ch.(*tool.Data); !ok {

		fmt.Println("dict-xml.go chan error")
		return
	}

	if data == nil {
		fmt.Println("there's no xml dictionary found")
		return
	}

	for _, i := range data.Data {

		item := new(tool.Item)
		*item = i
		// r.Dict.AddItem(i.Desc, item)
		r.Dict.AddValue(i.Desc, item)
	}
}

// LoadXML used to load dictionary from xml format file
func LoadXML(dict core.Dict, path string) {

	r := new(RoutXML)
	r.Dict = dict

	core.RoutineLoadDir(r, path)

}
