package tool

import (
    "strings"
    "path/filepath"
    "os"
    "fmt"
)

type DirInfo struct {
    
    _focus      int
    _fileCnt    int
    _fileArr    []string
}


func (info *DirInfo) InitDir(root string) error {
    
    *info = DirInfo {
        _focus:     0,
        _fileCnt:   0,
        _fileArr:   nil,
    }
    if !strings.HasSuffix(root, "/") {
        root += "/"
    }
    
    err := filepath.Walk(root, info.handleWalk)
    if err != nil {
        
        fmt.Println("there's an error in dict reset : " + err.Error())
    }
    
    return err
}

func (info *DirInfo) NextFile() *string {
    
    var name *string = nil
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
    case file.Mode() & os.ModeSymlink > 0:
    case strings.HasPrefix(file.Name(), "."):
    case file.IsDir():
    default:
        info.addFile(path)
        
    }
    
    return nil
}

func (info *DirInfo) addFile(path string) {
    
    if len(info._fileArr) == info._fileCnt {
        
        info._fileArr = append(info._fileArr, make([]string, 10)...)
    }
    
    info._fileArr[info._fileCnt] = path
    info._fileCnt++
}