package gstr

import (
	"regexp"
	"strconv"
)

var (
	// octReg 检测八进制字符串的正则表达式
	octReg = regexp.MustCompile(`\\[0-7]{3}`)
)

// OctStr converts string container octal string to its original string,
// for example, to Chinese string.
// Eg: `\346\200\241` -> 怡
func OctStr(str string) string {
	return octReg.ReplaceAllStringFunc(
		str,
		func(s string) string {
			i, _ := strconv.ParseInt(s[1:], 8, 0)
			return string([]byte{byte(i)})
		},
	)
}
