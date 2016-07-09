package main

import (
	"hiank.net/translate/core"
	"hiank.net/translate/format"
	"hiank.net/translate/tool"
)

// var c chan int

// Rout used to
type Rout struct {
	dict core.Dict
	cfg  *tool.Config

	NilArr map[string]int
	T      map[string]int
}

// Filter used to
func (r *Rout) Filter() core.Filter {

	return nil
}

// Run used to
func (r *Rout) Run(path string) {

	t := trans.NewTrans()
	trans.LoadCCB()
	trans.Format(t, r.dict)

	tmp := []rune(path)

	dstPath := r.cfg.DstDir + string(tmp[len(r.cfg.SrcDir):len(path)])
	trans.SaveFile(t, dstPath)
	r.T = t.NilArr
}

// End used to
func (r *Rout) End() {

	arr := r.T
	for k := range arr {

		_, ok := r.NilArr[k]
		if !ok {
			r.NilArr[k] = 0
		}
	}

}

// LoadCCB used to translate ccb file
func LoadCCB(cfg tool.Config, d tool.Dict) {

	r := new(Rout)
	r.dict = d
	r.cfg = cfg
	r.NilArr = make(map[string]int)
	tool.RoutineLoadDir(r, cfg.SrcDir)

	trans.SaveNilToCSV(r.NilArr, cfg.DstDir+"../ccb.left")
}
