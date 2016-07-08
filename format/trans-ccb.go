package trans

import (
	"hiank.net/translate/core"
	// "strings"
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"hiank.net/translate/tool"
)

// CCBTrans used to translate ccb file
type CCBTrans struct {

	// _nilArr     []string
	_nilMap map[string]int
	// _nilCnt     int
	// _dict       core.Dict
}

// NewCCBTrans create and initialize struct CCBTrans
func NewCCBTrans() *CCBTrans {

	t := new(CCBTrans)
	*t = CCBTrans{

		_nilMap: make(map[string]int),
		// _nilCnt : 0,
		// _dict   : nil,
	}

	return t
}

// Format realize interface core.Trans
func (t *CCBTrans) Format(dstPath string, srcPath string, dict core.Dict) map[string]int {

	if dstPath == srcPath {
		fmt.Println("dstPath and srcPath must not same")
		return nil
	}

	sfile, err := os.OpenFile(srcPath, os.O_RDONLY, 0444)
	if err != nil {

		fmt.Println("open file err : " + err.Error())
		return nil
	}
	defer sfile.Close()

	dirName := string(dstPath[0:strings.LastIndex(dstPath, "/")])
	if _, e := os.Stat(dirName); e != nil {

		os.MkdirAll(dirName, 0755)
	}
	dfile, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {

		fmt.Println("create or open file err : " + err.Error())
		return nil
	}
	defer dfile.Close()

	decoder := xml.NewDecoder(sfile)
	encoder := xml.NewEncoder(dfile)

	nilMap := format(encoder, decoder, dict)
	// for k := range nilMap {

	//     if _, ok := t._nilMap[k]; !ok {
	//         fmt.Printf("key : %v_%v\n", k, t._nilMap)
	//         t._nilMap[k] = 0
	//     }
	// }

	encoder.Flush()
	return nilMap
}

func format(encoder *xml.Encoder, decoder *xml.Decoder, d core.Dict) (nilMap map[string]int) {

	// nilArr  := make([]string, 10)
	// nilCnt  := 0
	nilMap = make(map[string]int)
	dict := d.GetData()

	focus := false
	for tk, e := decoder.Token(); e == nil; tk, e = decoder.Token() {

	S:
		switch tmp := tk.(type) {

		case xml.StartElement:
			focus = tmp.Name.Local == "string"
			fmt.Printf("token show : %v\n", tmp.Name.Space)

		// case xml.EndElement:
		case xml.CharData:
			if !focus {
				break S
			}

			content := string([]byte(tmp))
			if strings.IndexFunc(content, matchF) == -1 {
				break S
			}

			if value, ok := dict[content]; ok {

				tk = xml.CharData([]byte(value.ToString()))
			} else {

				if _, ok := nilMap[content]; !ok {

					nilMap[content] = 0
				}
			}

		default:
		}
		encoder.EncodeToken(tk)
	}

	return
}

// AddNil realize interface core.AddNil
func (t *CCBTrans) AddNil(m map[string]int) {

	for k := range m {

		if _, ok := t._nilMap[k]; !ok {
			t._nilMap[k] = 0
		}
	}
}

// GetNilArr realize interface core.GetNilArr
func (t *CCBTrans) GetNilArr(ignoreArr []string) []string {

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

// FormatCCB used to
func FormatCCB(t core.Trans, d core.Dict, cfg *tool.Config) {

	// rout := new(Rout)
	// rout.Cfg = cfg
	// rout.T = t
	// rout.D = d

	// tool.RoutineLoadDir(rout, cfg.SrcDir)

	t.Format("/Users/hiank/code/workspace/translate/angel_en/ccb/AlchemyCell.ccb", "/Users/hiank/code/workspace/cocos2dx-2.2.6/projects/shenxian_auto_en/builderProject/ShenxianBuilder/Resources/ccb/AlchemyCell.ccb", d)
}
