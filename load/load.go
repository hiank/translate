package main

import (
    "hiank.net/translate/dict"
    "hiank.net/translate/tool"
    "hiank.net/translate/format"
	"time"
	"fmt"
)


func main() {
        
    lastTime := time.Now().UnixNano()

    cfg := new(tool.Config)
    tool.LoadConfig(cfg, "./config.json")
    
    d := dict.NewDict()
    for _, v := range cfg.DictDir.CSVDir {
        
        dict.LoadCSV(d, v)
    }
    
    for _, v := range cfg.DictDir.XMLDir {
        
        dict.LoadXML(d, v)
    }
    
    for _, v := range cfg.DictDir.EXDir {
        
        dict.LoadEx(d, v)
    }

    switch cfg.Type {
        case "ccb":     LoadCCB(cfg, d)
        case "config":  LoadConfig(cfg, d)
        case "lang":    LoadLang(cfg, d)
        case "Class":   LoadClass(cfg, d)
    }

    // trans.SaveNilToCSV(r.NilArr, cfg.DstDir + "../ccb.left")
    fmt.Printf("++++%v\n", time.Now().UnixNano() - lastTime)
}

