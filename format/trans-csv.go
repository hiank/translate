package trans

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"hiank.net/translate/core"
)

// "strings"

// CSVTrans used to translate ccb file
type CSVTrans struct {

	// _nilArr     []string
	_nilMap map[string]int
	// _nilCnt     int
	// _dict       core.Dict
}

// NewCSVTrans create and initialize struct CSVTrans
func NewCSVTrans() *CSVTrans {

	t := new(CSVTrans)
	*t = CSVTrans{

		_nilMap: make(map[string]int),
		// _nilCnt : 0,
		// _dict   : nil,
	}

	return t
}

// Format realize interface core.Trans
func (t *CSVTrans) Format(dstPath string, srcPath string, dict core.Dict) map[string]int {
	d := dict.GetData()
	nilMap := make(map[string]int)

	// file, err := os.Open(srcPath)
	// if err != nil {
	// 	fmt.Println("open file err : " + err.Error())
	// 	return nil
	// }
	dfile, sfile, err := openFiles(dstPath, srcPath)
	if err != nil {
		return nil
	}
	defer func() {
		dfile.Close()
		sfile.Close()
	}()

	r := bufio.NewReader(sfile)
	w := bufio.NewWriter(dfile)

	f := func(lineByte []byte) {

		lineStr := string(lineByte)
		t.formatLine(&lineStr, d, nilMap)
		w.WriteString(lineStr)
		// if needN {
		// 	w.WriteByte('\n')
		// }
	}

L:
	for {

		leftByte, _ := r.Peek(r.Buffered())
		lineByte, err := r.ReadSlice('\n')

		switch err {
		case nil:
			f(lineByte)
		case io.EOF:
			if len(leftByte) != 0 {
				f(lineByte)
			}
			break L
		default:
			fmt.Println("")
			break L
		}
	}

	w.Flush()

	return nilMap
}

func (t *CSVTrans) formatLine(lineStr *string, dict map[string]core.Value, nilMap map[string]int) {

	arr := strings.Split(*lineStr, ",")
	needReplace := false
	for idx, str := range arr {

		if strings.IndexFunc(str, matchF) == -1 {
			continue
		}

		if value, ok := dict[str]; ok {

			arr[idx] = value.ToString()
			needReplace = true
		} else {

			if _, ok := nilMap[str]; !ok {

				nilMap[str] = 0
			}
		}

	}

	if needReplace {
		*lineStr = strings.Join(arr, ",")
	}

}

// AddNil realize interface core.AddNil
func (t *CSVTrans) AddNil(m map[string]int) {

	for k := range m {

		if _, ok := t._nilMap[k]; !ok {
			t._nilMap[k] = 0
		}
	}
}

// GetNilArr realize interface core.GetNilArr
func (t *CSVTrans) GetNilArr(ignoreArr []string) []string {

	for _, s := range ignoreArr {

		if _, ok := t._nilMap[s]; ok {
			delete(t._nilMap, s)
		}
	}

	arr := make([]string, len(t._nilMap))
	i := 0
	for key := range t._nilMap {
		arr[i] = key
		i++
	}
	return arr
}
