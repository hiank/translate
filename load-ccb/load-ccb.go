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

    ccbD := dict.NewDict()
    for _, v := range cfg.DictDir.CSVDir {
        
        dict.LoadCSV(ccbD, v)
    }
    
    for _, v := range cfg.DictDir.XMLDir {
        
        dict.LoadXML(ccbD, v)
    }
    
    for _, v := range cfg.DictDir.EXDir {
        
        dict.LoadEx(ccbD, v)
    }

    fmt.Printf("++++%v\n", time.Now().UnixNano() - lastTime)
    

    // t := new(trans.CCBTrans)
    t := trans.NewCCBTrans()
    trans.FormatCCB(t, ccbD, cfg)

    trans.SaveNilToCSV(t.GetNilArr(cfg.DIgnore.Array), cfg.DstDir + "../ccb.left")
}

