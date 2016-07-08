package translate

import (
	"strings"
	"fmt"
)

/// beside string in class
func classBEx (str string) bool {
	rlt := false

	switch {
	case strings.Contains(str, "天命aaaaaa"): fallthrough
	case strings.Contains(str, "typeid"): fallthrough
	case strings.Contains(str, "npc数据没有"): fallthrough
	case strings.Contains(str, "登录线程"): fallthrough
	case strings.Contains(str, "－－－－"): fallthrough
	case strings.Contains(str, "======="): fallthrough
	case strings.Contains(str, "PanelType_"):
		return true
	}

	switch 0 {
	case strings.Compare(str, "】"): fallthrough
	case strings.Compare(str, "："): fallthrough
	case strings.Compare(str, "，"): fallthrough
	case strings.Compare(str, "（%s）"): fallthrough
	case strings.Compare(str, "一步一景"): fallthrough
	case strings.Compare(str, "拓拔野与姑射"): fallthrough
	case strings.Compare(str, "这是toast"): fallthrough
	case strings.Compare(str, "宋体"):
		return true
	}
	return rlt
}

func classB (lineStr string) bool {

	rlt := false

	switch {
	case strings.HasPrefix(lineStr, "//") || strings.HasPrefix(lineStr, "#"): fallthrough
	case strings.HasPrefix(lineStr, "cout") || strings.HasPrefix(lineStr, "CCLOG(") || strings.HasPrefix(lineStr, "CCLog(") || strings.HasPrefix(lineStr, "NSLog("): fallthrough
	case !strings.Contains(lineStr, "\""): fallthrough
	case strings.IndexFunc(lineStr, unicodeF) == -1:
		rlt = true
	case strings.HasPrefix(lineStr, "\""):
		fmt.Println("follow TEXT : " + lineStr)
	}

	return rlt
}

/// unicodeFunc
func unicodeF (data rune) bool {

	return data > 127
}