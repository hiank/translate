package tool

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"bufio"
)



type TFile struct {
    
    _arr        []string    // all content to line array
    _focusLine  int         // current operation line
    
    _path       string      // file name
}


func NewTFile(path string) *TFile {
    
    f := new(TFile)
    
    fmt.Printf(">>>#%v#", f)
    f._focusLine = 0
    f._path = path
    
/// the code below was used to read line array from file
    data, err := ioutil.ReadFile(f._path)
	if err != nil {
		fmt.Printf("error to read file #%v#\n", f._path)
		return nil
	}

	str := string(data)
	if strings.Contains(str, "\r\n") {
		fmt.Println("content windows new line symbol")
		str = strings.Replace(str, "\r\n", "\n", -1)
	}
	if strings.Contains(str, "\r") {
		
		fmt.Println("content mac os new line symbol")
		str = strings.Replace(str, "\r", "\n", -1)
	}

	f._arr = strings.Split(str, "\n")
	return f
}

func (f *TFile) Next() *string {
	
	return f.NextLine(f._focusLine)
}

func (f *TFile) NextLine(lastIdx int) *string {
    
	f._focusLine = lastIdx
    if f._arr == nil || (f._focusLine+1 == len(f._arr)) {
        return nil
    }
    
    next := &f._arr[f._focusLine]
    f._focusLine++
    
    return next
}

func (f *TFile) Save(opt func (line *string)) {
    
    f.SaveFile(opt, f._path)
}

func (f *TFile) SaveFile(opt func (line *string), dstName string) {
   
	// dstName := strings.Replace(info._fileName, info._info.RootSrc, info._info.RootDst, 1)
	// // fmt.Println("file dst name is " + dstName)

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
	
	for idx, con := range f._arr {
		// show := false
		if idx != 0 {
			if con == "\r" || con == "" {
				continue
			}
			opt(&con)
			w.WriteByte('\n')
		}
		// if show {
		// 	fmt.Println("+-+-+-+-+-+-+-+-+-+-" + con)
		// }
		w.WriteString(con)

	}
	w.Flush()
}

