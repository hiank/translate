package trans


// import (
    
//     "strings"
//     // "fmt"
//     "hiank.net/translate/tool"
// )

// // CsvTrans used to 
// type CsvTrans struct {
    
    
//     _data       []string
//     _focusIdx   int
//     _total      int
    
//     // _locked     bool
//     NilArr      map[string]int
// }

// // NewCsvTrans used to make 
// func NewCsvTrans(path string) *CsvTrans {
    
    
//     t := new(CsvTrans)
//     t._focusIdx = -1
//     t.NilArr = make(map[string]int)
    
//     LoadToLine(t, path)
//     return t
// }

// // GetData used to 
// func (t *CsvTrans) GetData() []string {
    
//     return t._data[0:t._total]
// }


// // FormatLine used to format line
// func (t *CsvTrans) FormatLine(lineStr *string, dict tool.Dict) bool {
    
    
//     arr := strings.Split(*lineStr, ",")
//     needReplace := false
//     for idx, str := range arr {
        
//         if strings.IndexFunc(s, matchF) == -1 {
//             continue
//         }
        
//         item := dict.GetItem(str)
//         if item == nil {
            
//             t.NilArr[str] = 0
//             continue
//         }
        
//         arr[idx] = item.Value
//         needReplace = true
//     }

//     if needReplace {
//         *lineStr = strings.Join(arr, ",")
//     }
    
//     return true
// }


// // AddLine used to 
// func (t *CsvTrans) AddLine(line string) {
    
//     // fmt.Printf("addline : #%v#\n", []rune(line))
//     if t._data == nil {
//         t._data = make([]string, 10)
//     }
    
//     if t._total >= len(t._data) {
//         t._data = append(t._data, make([]string, 10)...)
//     }
    
//     t._data[t._total] = line
//     t._total++
// }

// // ContentFilter used to
// func (t *CsvTrans) ContentFilter() tool.Filter {
    
//     return t
// }

// // Match used to
// func (t *CsvTrans) Match(c string) bool {
    
//     if strings.IndexFunc(s, matchF) == -1 {
        
//         return false
//     }
    
//     return true
// }
