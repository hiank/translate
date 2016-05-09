package tool


// type DestInfo struct {
    
//     _dictdst    map[string]string
//     _dicterr    map[string]string
    
//     _filename   string
// }

type Item struct {
    
    Key     int     `xml:"key,attr"`
    Desc    string  `xml:"desc,attr"`
    Value   string  `xml:"value"`
}


// type DictI interface {
    
//     GetDict() *map[string]*Item
//     SetDict(*map[string]*Item)
// }
