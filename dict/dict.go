package dict

import (
    "hiank.net/translate/tool"
    // "strings"
    "fmt"
	"hiank.net/translate/core"
)


// Dict used to 
type Dict struct {
    
    _data   map[string]core.Value
}

// NewDict used to create a ClassDict and initialize the memory
func NewDict() *Dict {
    
    dict := new(Dict)
    dict._data = make(map[string]core.Value)
    
    return dict
}

// GetData used to get _data
func (dict *Dict) GetData() map[string]core.Value {
    
    return dict._data
}


// GetValue used to get value from _data
func (dict *Dict) GetValue(key string) core.Value {
    
    
    return dict._data[key]
}

// AddDict used to realize interface core.Dict
func (dict *Dict) AddDict(d core.Dict) {


}

// AddValue used to add value to _data
func (dict *Dict) AddValue(key string, value core.Value) {
    
    dict._data[key] = value
}



func addItem(dict core.Dict, key string, value string) {
    
    if key == "" {
        fmt.Println("add no key item")
        return
    }
    // item := dict.GetItem(key)
    v := dict.GetValue(key)
    if v != nil && v.ToString() == value {
        return
    }
    
    if value == "" {
        fmt.Printf("add no value item #%v#\n", key)
        return
    }

    item := new(tool.Item)
    *item = tool.Item {
        
        Key:    0,
        Desc:   key,
        
        Value:  value,      
    }
    
    // dict.AddItem(key, item)
    dict.AddValue(key, item)
    // fmt.Printf("++++%v\n", key)
    return
}


