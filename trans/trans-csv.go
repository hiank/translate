package trans

import (

    "io"
    "os"
	"fmt"
    "bufio"
	"hiank.net/translate/dict"
    "strings"
    
	"hiank.net/translate/tool"
)

// CSVData used to format ccb
type CSVData struct {

    _data       []string
    _focusIdx   int
    _total      int
}

// Load used to
func (c *CSVData) Load(f *os.File) {

    r := bufio.NewReader(f)
    loop:	for {
        // fmt.Printf("+-+-+-:%v\n", r.Buffered())
        leftByte, _ := r.Peek(r.Buffered())
        lineByte, err := r.ReadSlice('\n')
        // r.Peek()
        // lineByte, _, err := r.ReadLine()
        switch err {
        case nil:
            c.addLine(string(lineByte))
            // fmt.Printf("__#%v#\n", string(lineByte))
        case io.EOF:
            if len(leftByte) != 0 {
                c.addLine(string(leftByte))
                // fmt.Printf("EX__#%v#\n", string(lineByte))
            }
            break loop
        default:
            fmt.Println("there's an error when read line : " + err.Error())
            break loop
        }

    }
}

func (c *CSVData) addLine(l string) {

        // fmt.Printf("addline : #%v#\n", []rune(line))
    if c._data == nil {
        c._data = make([]string, 10)
    }
    
    if c._total >= len(c._data) {
        c._data = append(c._data, make([]string, 10)...)
    }
    
    c._data[c._total] = l
    c._total++
}



// FormatCSV used to translate
func FormatCSV(t *CSVData, d *dict.Dict) {

    format := func (lineStr *string) bool {

        arr := strings.Split(*lineStr, ",")
        needReplace := false
        for idx, str := range arr {
            
            if strings.IndexFunc(str, matchF) == -1 {
                continue
            }
            
            item := d.GetItem(str)
            if item == nil {
                
                // t.NilArr[str] = 0
                continue
            }
            
            arr[idx] = item.Value
            needReplace = true
        }

        if needReplace {
            *lineStr = strings.Join(arr, ",")
        }
        
        return needReplace
    }

    for i, l := range t._data {

        if !format(&l) {
            continue
        }

        t._data[i] = l
    }

}
 


func matchF(c rune) bool {
    
    if !tool.UnicodeF(c) {
        return false
    }
    
    rlt := true
    switch c {
    case '，':
        rlt = false
    }
    
    return rlt
}