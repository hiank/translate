package dict

import (
	"fmt"
	"strconv"
	"strings"

	"hiank.net/translate/core"
	"hiank.net/translate/tool"
)

// EXData used to
type EXData struct {
	_mapL map[int]string
	_mapR map[int]string
}

// RoutEx used to
type RoutEx struct {
	Dict core.Dict
	// T    *EXData
}

// Match is the function realize for FileFilter interface
func (r *RoutEx) Match(name string) bool {

	rlt := false
	switch {
	case strings.HasSuffix(name, ".left"):
		rlt = true

	}

	return rlt
}

// Filter is the function realize for Routine interface
func (r *RoutEx) Filter() core.Filter {

	return r
}

func (r *RoutEx) loadMap(t *tool.TFile) map[int]string {

	m := make(map[int]string)
L:
	for {
		lineStr := t.Next()
		if lineStr == nil {
			break L
		}

		// lineArr := strings.Split(*lineStr, ",")
		locked := false
		lineArr := strings.FieldsFunc(*lineStr, func(ru rune) bool {
			rlt := false
			switch ru {
			case ',':
				if !locked {
					rlt = true
				}
			case '"':
				locked = !locked
			}
			return rlt
		})

		if len(lineArr) != 2 {
			fmt.Printf("dict-ex.go error #%v#\n", len(lineArr))
			continue
		}

		keyStr := strings.TrimFunc(lineArr[0], tool.NumberNF)
		keyInt, err := strconv.Atoi(keyStr)
		if err != nil {
			fmt.Printf("dict-ex error key : #%v#\n", keyStr)
			continue
		}

		if m[keyInt] != "" {
			fmt.Printf("had : %v", m[keyInt])
			continue
		}

		m[keyInt] = strings.Trim(lineArr[1], "\"")
	}
	return m
}

// Run used to
func (r *RoutEx) Run(path string) core.RoutineChan {

	tmp := []rune(path)
	fileR := string(tmp[0:len(path)-len(".left")]) + ".right"
	t := tool.NewTFile(fileR)
	fmt.Printf(".....%v\n", fileR)
	if t == nil {
		return r
	}
	d := new(EXData)
	d._mapR = r.loadMap(t)

	t = tool.NewTFile(path)
	d._mapL = r.loadMap(t)

	// r.T = d
	return d
}

// End used to
func (r *RoutEx) End(ch core.RoutineChan) {

	var data *EXData
	var ok bool
	if data, ok = ch.(*EXData); !ok {

		fmt.Println("dict-ex.go chan error")
		return
	}
	// data := r.T
	if data == nil || data._mapL == nil || data._mapR == nil {
		return
	}

	mapL, mapR := data._mapL, data._mapR

	for key, value := range mapL {

		if mapR[key] == "" {
			fmt.Printf("cann't load key : %v\n", value)
			continue
		}

		addItem(r.Dict, value, mapR[key])
	}
}

// LoadEx used to load two file type dictionary
func LoadEx(dict core.Dict, path string) {

	r := new(RoutEx)
	r.Dict = dict

	core.RoutineLoadDir(r, path)
}
