/*
@Time : 2019/11/28 0028 下午 4:41
@Author : huanfuan
@File : pathUtils
@Software: GoLand
*/

package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func RootPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Panicln("发生错误", err.Error())
	}
	i := strings.LastIndex(s, "\\")
	path := s[0 : i+1]
	return path
}
