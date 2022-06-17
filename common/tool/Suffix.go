package tool

import "strings"

//获取后缀
func GetSuffix(fileName string) string {
	return fileName[strings.LastIndex(fileName, ".")+1:]
}
