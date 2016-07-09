package main

import (
	"fmt"
	"time"

	"hiank.net/translate/dict"
	"hiank.net/translate/format"
	"hiank.net/translate/tool"
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

	fmt.Printf("++++%v\n", time.Now().UnixNano()-lastTime)

	t := trans.NewCSVTrans()
	trans.Format(t, ccbD, cfg)

	trans.SaveNilToCSV(t.GetNilArr(cfg.DIgnore.Array), cfg.DstDir+"../ccb.left")
}

// const (
// 	dstDir = "/Users/hiank/code/workspace/god_trans/"
// 	csvDir = "/Volumes/data_mac/local_svn/root/god_tw_local/Resources/config/"
// )

// // var c chan int

// // Rout used to
// type Rout struct {
// 	Dict   tool.Dict
// 	NilArr map[string]int
// 	T      map[string]int
// }

// // MatchName used to
// func (r *Rout) MatchName(name string) bool {

// 	rlt := true
// 	switch {
// 	case strings.HasSuffix(name, "HeroFirstName.csv"):
// 		fallthrough
// 	case strings.HasSuffix(name, "HeroMan.csv"):
// 		fallthrough
// 	case strings.HasSuffix(name, "HeroWoman.csv"):
// 		fallthrough
// 	case strings.HasSuffix(name, "XKeyWord.csv"):
// 		fallthrough
// 	case strings.HasSuffix(name, "downConfig.csv"):
// 		rlt = false
// 	}

// 	return rlt
// }

// // Filter used to
// func (r *Rout) Filter() tool.FileFilter {

// 	return r
// }

// // Run used to
// func (r *Rout) Run(path string) {

// 	t := trans.NewCsvTrans(path)

// 	trans.Format(t, r.Dict)

// 	tmp := []rune(path)

// 	dstPath := dstDir + "ccb/" + string(tmp[len(csvDir):len(path)])
// 	trans.SaveFile(t, dstPath)
// 	r.T = t.NilArr
// }

// // End used to
// func (r *Rout) End() {

// 	arr := r.T
// 	for k, _ := range arr {

// 		_, ok := r.NilArr[k]
// 		if !ok {
// 			r.NilArr[k] = 0
// 		}
// 	}
// }

// func main() {

// 	lastTime := time.Now().UnixNano()
// 	csvD := dict.NewDict()
// 	dict.LoadEx(csvD, dstDir+"csv_dir")
// 	// dict.LoadCSV(csvD, "/Users/hiank/code/workspace/god_trans/dict/csv")
// 	// dict.LoadEx(csvD, "/Users/hiank/code/workspace/god_trans/dict/ex")

// 	// dict := loadDict()
// 	fmt.Printf("++++%v\n", time.Now().UnixNano()-lastTime)

// 	r := new(Rout)
// 	r.Dict = csvD
// 	r.NilArr = make(map[string]int)
// 	tool.RoutineLoadDir(r, csvDir)

// 	trans.SaveNilToCSV(r.NilArr, dstDir+"csv.left")
// }
