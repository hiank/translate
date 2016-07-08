package trans

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"hiank.net/translate/core"
	"hiank.net/translate/tool"
	// "strconv"
)

func parseDir(path string) {

	dirName := string(path[0:strings.LastIndex(path, "/")])

	if _, e := os.Stat(dirName); e != nil {

		os.MkdirAll(dirName, 0755)
	}
}

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
	// i := 0
	// nilArr := s.GetNilArr()
	for _, k := range nilArr {

		// w.WriteString(strconv.Itoa(i))
		// w.WriteByte(',')
		w.WriteString(k)
		w.WriteByte('\n')

		// i++
	}
	w.Flush()
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
func (r *Rout) Filter() tool.Filter {

	return nil
}

// Run operate file
func (r *Rout) Run(path string) tool.RoutineChan {

	tmp := []rune(path)
	dstPath := r.Cfg.DstDir + string(tmp[len(r.Cfg.SrcDir):len(path)])

	// ch := new(TransChan)
	// ch.NilMap = r.T.Format(dstPath, path, r.D)

	// return ch
	return r.T.Format(dstPath, path, r.D)
}

// End work after operate file end
func (r *Rout) End(ch tool.RoutineChan) {

	// r.T.AddNil((map[string]int)ch)
	switch c := ch.(type) {
	case map[string]int:
		r.T.AddNil(c)

	default:
		fmt.Printf("type error %v\n", c)
	}
}
