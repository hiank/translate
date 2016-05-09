package translate


import (
	"strings"

	"hiank.net/translate/tool"
	"fmt"
)



type CSVTool struct {
								   //	_dict		map[string]int

	_dict 		*Dict

	_focus 		*FileInfo 		/// the file's classinfo what is working
	_info 		*FlushInfo

//	_content 	string
	_fileArr 	*FileArr
	
	NilArr 		[]string
}


func NewCSV(root string) *CSVTool {

	tool := new(CSVTool)
	*tool = CSVTool{


		_dict: 		nil,//NewDict(dictName),
		_focus: 	nil,
		_info: 		nil,

		_fileArr:	new(FileArr),
	}

	//	tool._dict.InitDict(dictName)
	tool._fileArr.InitArr(root)

	return tool
}

func (t *CSVTool) InitDictWithPath(path string) {
	
	t._dict = NewDict(path)
}

func (t *CSVTool) InitDictWithDict(dict map[string]*tool.Item) {
	
	t._dict = new(Dict)
	t._dict._dictFile = ""
	t._dict._dict = dict
	
	
}


func (tool *CSVTool) pushLine(line []byte) {


	lineStr := string(line)

	if strings.IndexFunc(lineStr, unicodeF) == -1 {

		return
	}
	// fmt.Printf(">>>%v\n", lineStr)
	arr := strings.Split(lineStr, ",")
	for _, str := range arr {

		if strings.IndexFunc(str, unicodeF) != -1 {

			tmp := strings.Replace(str, "，", "", -1)
			if strings.IndexFunc(tmp, unicodeF) == -1 {
				continue
			}
			
			item := tool._dict._dict[str]
			if item == nil {

				tool._dict._dict[str] = nil
				// fmt.Printf(">>>%v\n", str)
			} else {

				tool._focus._dict[str] = item
				// fmt.Printf(">>>%v\n", item)
			}
		}
	}

}


func (tool *CSVTool) nextFile() *FileInfo {

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



func (tool *CSVTool) SetFlushInfo(info *FlushInfo) {

	tool._info = info
}

func (tool *CSVTool) GetFlushInfo() *FlushInfo {

	return tool._info
}

func (tool *CSVTool) flush() {

	tool._focus.save(func(lineStr string, info *FileInfo) string {

		arr := strings.Split(lineStr, ",")
		needReplace := false
		for idx, str := range arr {
			
			if strings.IndexFunc(str, unicodeF) == -1 {
				continue
			}

			tmp := strings.Replace(str, "，", "", -1)
			if strings.IndexFunc(tmp, unicodeF) == -1 {
				continue
			}
			
			item := tool._dict._dict[str]
			if item == nil {
				fmt.Println("<<<no value : #" + str)
				continue
			}
			
			arr[idx] = item.Value
			needReplace = true
		}

		if needReplace {
			lineStr = strings.Join(arr, ",")
		}

		return lineStr
	})
}

func (tool *CSVTool) Flush() {

	tool._dict.Flush()
}
