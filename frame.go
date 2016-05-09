package translate

import (
	"os"
	"strings"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"hiank.net/translate/tool"
	"bufio"
)



type FileInfo struct {

	_fileName 	string

	_dict 		map[string]*tool.Item

 	_info 		*FlushInfo
}

func NewFileInfo(flushInfo *FlushInfo, fullName string) *FileInfo {

	info := new(FileInfo)

	*info = FileInfo{

		_fileName: 		fullName,
		_dict: 			make(map[string]*tool.Item),

		_info: 			flushInfo,

	}

	return info
}

func (info *FileInfo) name() string {

	return info._fileName
}


func (info *FileInfo) save(f func(content string, i *FileInfo) string) {

	if info._info == nil || info._info.RootDst == "" {
		return
	}

	data, err := ioutil.ReadFile(info._fileName)
	if err != nil {
		fmt.Println("error to read file " + info._fileName)
		return
	}

	str := string(data)
	if strings.Contains(str, "\r\n") {
		fmt.Println("............")
		str = strings.Replace(str, "\r\n", "\n", -1)
	}
	if strings.Contains(str, "\r") {
		
		fmt.Println("***********")
		str = strings.Replace(str, "\r", "\n", -1)
	}
	fmt.Printf("-++--+-+--+-+-+%v\n", strings.Count(str, "\n"))
	arr := strings.Split(str, "\n")
	// fmt.Printf("++++++++%v\n", arr)
	// content := f(string(data), info)

//	for key, value := range info._dict {
//
//		key = "\"" + key + "\""
//		content = strings.Replace(content, key, "L(" + strconv.Itoa(value) + ")/*" + key + "*/", -1)
//
//	}

	dstName := strings.Replace(info._fileName, info._info.RootSrc, info._info.RootDst, 1)
	// fmt.Println("file dst name is " + dstName)

	dirName := string(dstName[0:strings.LastIndex(dstName, "/")])

	if _, e := os.Stat(dirName); e != nil {

		os.MkdirAll(dirName, 0755)

	}
	// ioutil.WriteFile(dstName, []byte(content), 0644)
	file, err := os.OpenFile(dstName, os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("<<<error to write file " + dstName)
		return
	}
	defer file.Close()
	
	w := bufio.NewWriter(file)
	
	for idx, con := range arr {
		show := false
		if idx != 0 {
			if con == "\r" || con == "" {
				fmt.Printf(">>>%v<<<%v\n", info._fileName, con)
				show = true
				continue
			}
			con = f(con, info)
			w.WriteByte('\n')
		}
		if show {
			fmt.Println("+-+-+-+-+-+-+-+-+-+-" + con)
		}
		w.WriteString(con)	
		// fmt.Printf("write :%v\n", con)		
	}
	w.Flush()
	

}


//func (info *FileInfo) flush() {
//
//	if info._info == nil || info._info.RootDst == "" {
//		return
//	}
//
//	data, err := ioutil.ReadFile(info._fileName)
//	if err != nil {
//		fmt.Println("error to read file " + info._fileName)
//		return
//	}
//
//	content := string(data)
//	for key, value := range info._dict {
//
//		key = "\"" + key + "\""
//		content = strings.Replace(content, key, "L(" + strconv.Itoa(value) + ")/*" + key + "*/", -1)
//
//	}
//
//
//	dstName := strings.Replace(info._fileName, info._info.RootSrc, info._info.RootDst, 1)
//	fmt.Println("file dst name is " + dstName)
//
//	dirName := string(dstName[0:strings.LastIndex(dstName, "/")])
//
//	if _, e := os.Stat(dirName); e != nil {
//
//		os.MkdirAll(dirName, 0755)
//
//	}
//	ioutil.WriteFile(dstName, []byte(content), 0644)
//	return
//}


type Frame interface {


	pushLine(line []byte)	// line was end with '\n'
//	pushEnd()				// file read over
	flush()					// operation the file end
//	exception() bool 		// the exception to break

	nextFile() *FileInfo 		// file name

	SetFlushInfo(info *FlushInfo)
	GetFlushInfo() *FlushInfo
}


type FlushInfo struct {

	RootSrc 	string
	RootDst 	string
}


type FileArr struct {

//	_rootDir 	string

	_fileIdx  	int
	_fileCnt  	int      // file count
	_fileArr  	[]string // the files waiting for dispose
}


func (arr *FileArr)InitArr(root string) {

	if !strings.HasSuffix(root, "/") {
		root += "/"
	}

//	arr._rootDir = root
	err := filepath.Walk(root, arr.handleWalk)
	if err != nil {
		fmt.Println("there's an error in dict reset : " + err.Error())
	}
}


func (arr *FileArr) handleWalk(path string, info os.FileInfo, err error) error {

	if info == nil {
		return err
	}

	//	fmt.Println("_______info name : " + info.Name() + "__path : " + path)
	if strings.HasSuffix(path, "/") {
		return nil
	}

	switch {
	case info.IsDir():
	case info.Mode() & os.ModeSymlink > 0:
	case strings.HasPrefix(info.Name(), "."):
	default:
//				fmt.Println("+++file path : " + path)
//		d._fileArr[len(d._fileArr)] = path
		arr.addFile(path)
	}

	return nil
}

func (arr *FileArr) addFile(path string) {

	if len(arr._fileArr) == arr._fileCnt {
		arr._fileArr = append(arr._fileArr, make([]string, 10)...)
	}

	arr._fileArr[arr._fileCnt] = path
	arr._fileCnt++
}
