package dict

import (
	"fmt"
	"strings"

	"hiank.net/translate/core"
	"hiank.net/translate/tool"
)

// Rout used to
type Rout struct {
	Dict core.Dict
	// T    *tool.TFile
}

// Filter used to
func (r *Rout) Filter() core.Filter {

	return nil
}

// Run used to
func (r *Rout) Run(path string) core.RoutineChan {

	t := tool.NewTFile(path)
	// r.T = t

	return t
}

// End used to
func (r *Rout) End(ch core.RoutineChan) {

	var t *tool.TFile
	var ok bool
	if t, ok = ch.(*tool.TFile); !ok {
		fmt.Println("dict-csv.go error")
		return
	}

	headStr := t.NextLine(0) //the first line
	headArr := strings.Split(*headStr, ",")

	num := len(headArr)
	if num < 2 {
		fmt.Println("error format for dict")
		return
	}

L:
	for {
		lineStr := t.Next()
		if lineStr == nil {
			break L
		}

		lineArr := make([]string, num)
		formatKey([]rune(*lineStr), lineArr)
		if len(lineArr) != num {
			fmt.Printf("error format string #%v#, bad num #%v#%v#\n", *lineStr, len(lineArr), num)
			continue
		}

		for idx, str := range headArr {

			if str != "" || lineArr[idx] == "" {
				continue
			}

			addItem(r.Dict, lineArr[idx-1], strings.Trim(lineArr[idx], " "))
		}
	}
}

func formatKey(data []rune, keyArr []string) {

	locked := false
	// var key []rune
	left, idx := 0, 0
	for i, r := range data {

		switch r {
		case ',':
			if !locked {
				// addKey(keyArr, data[left:i], idx)
				value := strings.Trim(string(data[left:i]), "\"")
				keyArr[idx] = strings.Replace(value, ",", "~", -1)
				fmt.Printf("#%v#\n", keyArr[idx])
				idx++
				left = i + 1
			}
		case '"':
			locked = !locked

		}
	}
	fmt.Printf("left : %v..len : %v\n", left, len(data))
	value := strings.Trim(string(data[left:len(data)]), "\"")
	keyArr[idx] = strings.Replace(value, ",", "~", -1)

	fmt.Printf("-----%v\n", keyArr[idx])
	// addKey(keyArr, data[left:len(data)-1], idx)

}

// LoadCSV used to load dictionary from csv format file
func LoadCSV(dict core.Dict, path string) {

	r := new(Rout)
	r.Dict = dict

	core.RoutineLoadDir(r, path)

}
