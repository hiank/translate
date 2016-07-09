package trans

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"hiank.net/translate/core"
	"hiank.net/translate/tool"
	// "strconv"
)

// SaveNilToCSV used to save key not content value to csv format
func SaveNilToCSV(nilArr []string, path string) {

	parseDir(path)

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("error to write " + path)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	i := 0
	// nilArr := s.GetNilArr()
	for _, k := range nilArr {

		w.WriteString(strconv.Itoa(i))
		w.WriteByte(',')
		w.WriteString(k)
		w.WriteByte('\n')

		i++
	}
	w.Flush()
}

func parseDir(path string) {

	dirName := string(path[0:strings.LastIndex(path, "/")])

	if _, e := os.Stat(dirName); e != nil {

		os.MkdirAll(dirName, 0755)
	}
}

func matchF(c rune) bool {

	if !tool.UnicodeF(c) {
		return false
	}

	rlt := true
	switch c {
	case '，':
		rlt = false
	}

	return rlt
}

// Rout used to operate director
type Rout struct {
	Cfg *tool.Config
	T   core.Trans
	D   core.Dict

	// 此处总觉得写的不严谨，感觉会出问题，以后再用更严谨的方式重写
	// nilMap  map[string]int
}

// Filter used to filter the file
func (r *Rout) Filter() core.Filter {

	return nil
}

// Run operate file
func (r *Rout) Run(path string) core.RoutineChan {

	tmp := []rune(path)
	dstPath := r.Cfg.DstDir + string(tmp[len(r.Cfg.SrcDir):len(path)])

	// ch := new(TransChan)
	// ch.NilMap = r.T.Format(dstPath, path, r.D)

	// return ch
	return r.T.Format(dstPath, path, r.D)
}

// End work after operate file end
func (r *Rout) End(ch core.RoutineChan) {

	// r.T.AddNil((map[string]int)ch)
	switch c := ch.(type) {
	case map[string]int:
		r.T.AddNil(c)

	default:
		fmt.Printf("type error %v\n", c)
	}
}

// Format used to
func Format(t core.Trans, d core.Dict, cfg *tool.Config) {

	rout := new(Rout)
	rout.Cfg = cfg
	rout.T = t
	rout.D = d

	core.RoutineLoadDir(rout, cfg.SrcDir)

}

func openFiles(dstPath string, srcPath string) (dfile *os.File, sfile *os.File, err error) {

	sfile, err = os.OpenFile(srcPath, os.O_RDONLY, 0444)
	if err != nil {

		fmt.Println("open file err : " + err.Error())
		return
	}
	defer func() {
		if err != nil {
			sfile.Close()
		}
	}()
	// defer sfile.Close()

	dirName := string(dstPath[0:strings.LastIndex(dstPath, "/")])
	if _, e := os.Stat(dirName); e != nil {

		os.MkdirAll(dirName, 0755)
	}
	dfile, err = os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {

		fmt.Println("create or open file err : " + err.Error())
		return
	}
	defer func() {
		if err != nil {
			dfile.Close()
		}
	}()

	return
}
