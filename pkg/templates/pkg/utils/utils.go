/*
@Time : 2020/3/19 16:55
@Author : FB
@File : utils.go
@Software: GoLand
*/
package utils

import (
	"os"
	"runtime"
	"strconv"
	"strings"
)

func PathExists(path string) (err error) {
	_, err = os.Stat(path)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		return
	}
	return
}

func GetSystem() string {
	return runtime.GOOS
}

func IsLinuxPlatform() bool {
	return strings.EqualFold(GetSystem(), "linux")
}

func IsWindowsPlatform() bool {
	return strings.EqualFold(GetSystem(), "windows")
}

func GetPathSeparator() string {
	return string(os.PathSeparator)
}

func Str2Uint(str string) (out uint) {
	f, _ := strconv.ParseFloat(str, 64)
	out = uint(f)
	return
}

func Str2Float(str string) (f float64) {
	f, _ = strconv.ParseFloat(str, 64)
	return
}

func Float2Str(f float64, prec int) (str string) {
	str = strconv.FormatFloat(f, 'f', prec, 64)
	return
}
