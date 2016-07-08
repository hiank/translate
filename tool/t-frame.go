package tool

// import (
// "encoding/xml"
// )
// type DestInfo struct {

//     _dictdst    map[string]string
//     _dicterr    map[string]string

//     _filename   string
// }

// Item used to operate dictionary in xml mode
type Item struct {
	Key   int    `xml:"key,attr"`
	Desc  string `xml:"desc,attr"`
	Value string `xml:"value"`
}

// ToString used to realize interface core.Value
func (item *Item) ToString() string {

	return item.Value
}

// Data used to
type Data struct {

	// Name    xml.Name    `xml:"Data"`
	Data []Item `xml:"Item"`
}

// Filter used to check up if the file name  match condition
type Filter interface {
	Match(s string) bool
}

// File used to operate the file
type File interface {
}

// type DictI interface {

//     GetDict() *map[string]*Item
//     SetDict(*map[string]*Item)
// }

// // Dict used to manager dictionary
// type Dict interface {

//     GetData() map[string]*Item

//     GetItem(key string) *Item
//     AddItem(key string, value *Item)
// }

// // Trans used
// type Trans interface {

//     Format(f func (l *string) bool)
//     GetData() []string
//     AddLine(line string)
//     AddNilItem(key string)
// }

// Routine used to operate
type Routine interface {
	Filter() Filter

	Run(path string) RoutineChan
	End(RoutineChan)
}

// RoutineChan used to
type RoutineChan interface {
}
