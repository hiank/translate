package id


import (
    "hiank.net/translate/tool"
    "fmt"
	"strings"
)

// Idcsv save the dictionary info
type Idcsv struct {
    
    _dict       map[string]*tool.Item
    // _rootDir    string
    
    _inited     bool
    // _dictErr    map[]
}

// NewIdcsv used to create new Idcsv memery
func NewIdcsv() *Idcsv {
    
    obj := new(Idcsv)
    obj._dict = nil
    // obj._rootDir = dirRoot
    obj._inited = false
    
    return obj
}

// InitDict used to init Idcsv._dict
func (obj *Idcsv) InitDict(path string) {
    
    dir := new(tool.DirInfo)
    if dir.InitDir(path) != nil {
        return
    }

    obj._inited = true    
    if obj._dict == nil {
        
        obj._dict = make(map[string]*tool.Item)
    }
    
    ci := make(chan *tool.TFile)
    cnt := 0
    for {
        
        name := dir.NextFile()
        if name == nil {
            break
        }

        cnt++
        go func () {
            
            file := tool.NewTFile(*name)
            // fmt.Printf(">>>>>>%v\n", *name)
            ci <- file
        }()
        // file := tool.NewTFile(*name)
        // if file != nil {
        //     obj.load(file)
        // }
    }

    
    
L:  for {
        select {
            case t := <- ci:
                if t != nil {
                    obj.load(t)
                } else {
                    fmt.Println("--=-=-=-=-=-=")
                }

                cnt--
                if cnt < 1 {
                    break L
                }
        }
    }
    
    fmt.Printf("-----%v\n", len(obj._dict))
}

func (obj *Idcsv) load(t *tool.TFile) {
    
    headStr := t.NextLine(0)   //the first line
    headArr := strings.Split(*headStr, ",")
    
    num := len(headArr)
    if num < 2 {
        fmt.Println("error format for dict")
        return
    }
    
L:  for {
        lineStr := t.Next()
        if lineStr == nil {
            break L
        }
        
        lineArr := strings.Split(*lineStr, ",")
        if len(lineArr) != num {
            fmt.Printf("error format string #%v#, bad num #%v#%v#\n", lineStr, len(lineArr), num)
            continue
        }
        
        for idx, str := range headArr {
            
            if str != "" || lineArr[idx] == "" {
                continue
            }
            
            obj.addItem(lineArr[idx-1], strings.Trim(lineArr[idx], " "))

        }
    }
}


func (obj *Idcsv) addItem(key string, value string) {
    
    if key == "" {
        fmt.Println("add no key item")
        return
    }
    item := obj._dict[key]
    if item != nil {
        
        if item.Value != value {
            fmt.Printf("already has %v : #%v# --- #%v#\n", key, item.Value, value)
        }    
        return
    }
    
    if value == "" {
        fmt.Printf("add no value item #%v#\n", key)
        return
    }

    item = new(tool.Item)
    *item = tool.Item {
        
        Key:    0,
        Desc:   key,
        
        Value:  value,      
    }
    
    
    obj._dict[key] = item
    // fmt.Printf("add item %v\n", item)
    return
}


