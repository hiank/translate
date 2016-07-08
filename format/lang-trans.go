package trans


// import (
//     "hiank.net/translate/tool"
//     "strings"
// )


// type LANGTrans struct {
    
    
// }


// // FormatLang used to 
// func FormatLang(t tool.Trans, d tool.Dict) {
    
//     t.Format(func (l *string) bool {
        
                
//         arr := strings.Split(*l, ",")
//         needReplace := false
//         for idx, str := range arr {
            
//             if strings.IndexFunc(str, matchF) == -1 {
//                 continue
//             }
            
//             item := d.GetItem(str)
//             if item == nil {
                
//                 t.NilArr[str] = 0
//                 continue
//             }
            
//             arr[idx] = item.Value
//             needReplace = true
//         }

//         if needReplace {
//             *l = strings.Join(arr, ",")
//         }
        
//         return true
//     })
// }