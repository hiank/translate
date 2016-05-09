package main 

import (
	"hiank.net/translate/id"
	"hiank.net/translate"
)

const (
	dst_dir = "/Users/hiank/code/workspace/temp/"

	class_dir = "/Users/hiank/code/workspace/cocos2dx-2.2.6/projects/ShenXian_id/Resources/lang/"
)

func main() {
    
    dictT := id.NewIDCsv()
    dictT.InitDict("/Users/hiank/code/workspace/classCsv")

    // c := translate.NewClass(class_dir)

    // *flushInfo = translate.FlushInfo{
    //     RootSrc:    class_dir,
    //     RootDst:    dst_dir + "Classes/",
    // }
    // c.SetFlushInfo(flushInfo)
    // translate.Trans(c)
    
    
    
    flushInfo := new(translate.FlushInfo)

    // config := translate.NewCSV(csv_dir, dst_dir + csv_dst_file)
    config := translate.NewClass(class_dir)
    config.InitDictWithDict(dictT.GetDict())

    *flushInfo = translate.FlushInfo{
        RootSrc:    class_dir,
        RootDst:    dst_dir + "lang/",
    }
    config.SetFlushInfo(flushInfo)
    translate.Trans(config)


}