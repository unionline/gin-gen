/*
@Time : 2020/4/1 14:52
@Author : FB
@File : valid.go
@Software: GoLand
*/
package utils

import "strings"

func ValidTemplateName(tmplName string) (err error) {

	seq := GetPathSeparator()
	//arr := strings.Split(tmplName, seq)
	//name := arr[len(arr)-1:]
	path_dir := strings.Join(
		[]string{
			"views",
			"template",
		}, seq)

	err = PathExists(path_dir + seq + tmplName)
	if err != nil {
		return
	}
	return
}
