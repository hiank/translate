package main 

import (
	"hiank.net/translate/dict"
	"hiank.net/translate/format"
    // "hiank.net/translate/tool"
    // "strings"
	"hiank.net/translate/tool"
)

// // Rout used to 
// type Rout struct {
    
//     cfg     tool.Config
//     dict    tool.Dict
// }
// // Run used to
// func (r *Rout) Run(path string) {
    
//     // t := trans.NewTrans()
//     // trans.LoadToXML(t, path)

//     // trans.Format(t)
//     // dstPath = dstDir + strings.TrimLeft(path, classDir)
//     // trans.SaveFile(t, dstPath)

// }
// // End used to
// func (r *Rout) End() {
    
    
// }

func main() {

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
    
    lang := dict.NewDict()
    dict.LoadXML(lang, cfg.SrcDir)
    
    arr := make(map[string]int)
    m := lang.GetData()
    for k, _ := range m {
        
        item := d.GetItem(k)
        if item == nil {
            arr[k] = 0
            continue
        }
        
        item.Key = m[k].Key
        m[k] = item
    }
    trans.SaveXML(lang, cfg.DstDir + "zh_tw.xml")
    
    trans.SaveNilToCSV(arr, cfg.DstDir + "../lang.left")
    // r := new(Rout)
    // r.dict = d
    // r.cfg = cfg
    
    // tool.RoutineLoadDir(r, cfg.SrcDir)
    
    // trans.SaveXMLKeyToCSV(d, cfg.DstDir + "../lang.left")
}