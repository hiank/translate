package translate


import (
	"strings"

	"fmt"
	"hiank.net/translate/tool"
)






type ClassTool struct {


	_dict 		*Dict

	_focus 		*FileInfo 		/// the file's classinfo what is working
	_info 		*FlushInfo

	_hide     	bool     // when some line start with "/*", _hide should set to true, and
	_fileArr 	*FileArr

}


func NewClass(root string) *ClassTool {

	tool := new(ClassTool)
	*tool = ClassTool{


		_dict: 		nil,//NewDict(dictName),
		_focus: 	nil,
		_info: 		nil,
		_hide: 		false,
		_fileArr:	new(FileArr),
	}

//	tool._dict.InitDict(dictName)
	tool._fileArr.InitArr(root)
	fmt.Printf("files len : %v\n", len(tool._fileArr._fileArr))
	return tool
}


func (t *ClassTool) InitDictWithPath(path string) {
	
	t._dict = NewDict(path)
}

func (t *ClassTool) InitDictWithDict(dict map[string]*tool.Item) {
	
	t._dict = new(Dict)
	t._dict._dictFile = ""
	t._dict._dict = dict
	
	
}


func (tool *ClassTool) pushLine(line []byte) {


	lineStr := strings.TrimSpace(string(line))
	switch {
	case strings.HasPrefix(lineStr, "/*") || tool._hide:
		tool._hide = !strings.HasSuffix(lineStr, "*/")
		return
	case classB(lineStr):
		return
	}

	formatB := lineStr[strings.Index(lineStr, "\"")+1:]
	arrStr := strings.Split(formatB, "\"")

	m := tool._focus._dict
	idx := 0
	for _, str := range arrStr {

		if idx % 2 == 0 && strings.IndexFunc(str, unicodeF) != -1 && !classBEx(str) {

			item := tool._dict._dict[str]
			if item == nil {

				tool._dict._dict[str] = nil
			} else {

				m[str] = item
			}
		}
		idx++
	}


}


func (tool *ClassTool) nextFile() *FileInfo {

	var name string
	fileArr := tool._fileArr
	if len(fileArr._fileArr) <= fileArr._fileIdx {
		return nil
	}
	name = fileArr._fileArr[fileArr._fileIdx]
	fileArr._fileIdx++

//	println("next file " + name)
	info := NewFileInfo(tool._info, name)
	tool._focus = info
//	fmt.Printf("info : %v", tool._focus)
//	tool._content = ""	///
	return info
}



func (tool *ClassTool) SetFlushInfo(info *FlushInfo) {

	tool._info = info
}

func (tool *ClassTool) GetFlushInfo() *FlushInfo {

	return tool._info
}

func (tool *ClassTool) flush() {

	// tool._focus.save(func(content string, info *FileInfo) string {

	// 	for key, value := range info._dict {

	// 		key = "\"" + key + "\""
	// 		content = strings.Replace(content, key, "L(" + strconv.Itoa(value.Key) + ")/*" + key + "*/", -1)

	// 	}
	// 	return content
	// })
	
	
	tool._focus.save(func(lineStr string, info *FileInfo) string {

		key := strings.Trim(lineStr, " ")
		key = strings.Trim(key, "\t")
		if strings.HasPrefix(key, "<value>") {
			
			item := tool._dict._dict[key]
			if item == nil {
				fmt.Printf("++++no value #%v#\n", key)

			} else {
				
				lineStr = strings.Replace(lineStr, key, item.Value, -1)
				fmt.Printf(">>>>>>>>>#%v#\n", item)
			}
		}
		fmt.Println("....<<<<" + lineStr)
		return lineStr
	})

}

func (tool *ClassTool) Flush() {

	tool._dict.Flush()
}
