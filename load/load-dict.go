package main

import (
    "hiank.net/translate/id"
    "hiank.net/translate"
    // "hiank.net/translate/tool"
	"time"
	"fmt"
)


const (
	dst_dir = "/Users/hiank/code/workspace/temp/"
	csv_dir = "/Users/hiank/Downloads/version/"
)

// var c chan int

func main() {
    
    lastTime := time.Now().UnixNano()
    dictT := id.NewIdcsv()
    dictT.InitDict("/Users/hiank/code/workspace/csv")
    // dictT.InitDict("/Users/hiank/code/workspace/csvEx")
    // dict := loadDict()
    fmt.Printf("++++%v\n", time.Now().UnixNano() - lastTime)
	// flushInfo := new(translate.FlushInfo)

    // config := translate.NewCSV(csv_dir, dst_dir + csv_dst_file)
    config := translate.NewCSV(csv_dir)
    config.InitDictWithDict(dictT.GetDict())

    flushInfo := translate.FlushInfo{
        RootSrc:    csv_dir,
        RootDst:    dst_dir + "config/",
    }
    config.SetFlushInfo(&flushInfo)
    translate.Trans(config)
    // if dictT.GetDict() == nil {
    //     return
    // }
    
    config.Flush(dst_dir + "temp.csv")
    
    
}

