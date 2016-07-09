package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// DirInfo used to manage director
type DirInfo struct {
	_focus   int
	_fileCnt int
	_fileArr []string

	_filter Filter
}

// InitDir used to initialize the DirInfo
func (info *DirInfo) InitDir(root string, filter Filter) error {

	*info = DirInfo{
		_focus:   0,
		_fileCnt: 0,
		_fileArr: nil,
	}
	if !strings.HasSuffix(root, "/") {
		root += "/"
	}

	info._filter = filter
	err := filepath.Walk(root, info.handleWalk)
	if err != nil {

		fmt.Println("there's an error in dict reset : " + err.Error())
	}

	return err
}

// NextFile used to pop next file name
func (info *DirInfo) NextFile() *string {

	var name *string
	if info._focus < info._fileCnt {

		name = &info._fileArr[info._focus]
		info._focus++
	}

	return name
}

func (info *DirInfo) handleWalk(path string, file os.FileInfo, err error) error {

	if file == nil {
		return err
	}

	if strings.HasSuffix(path, "/") {
		return nil
	}

	switch {
	case file.Mode()&os.ModeSymlink > 0:
	case strings.HasPrefix(file.Name(), "."):
	case file.IsDir():
	default:
		info.addFile(path)

	}

	return nil
}

func (info *DirInfo) addFile(path string) {

	if info._filter != nil && !info._filter.Match(path) {
		return
	}

	if len(info._fileArr) == info._fileCnt {

		info._fileArr = append(info._fileArr, make([]string, 10)...)
	}

	info._fileArr[info._fileCnt] = path
	info._fileCnt++
}
