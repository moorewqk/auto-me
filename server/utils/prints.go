package utils

import (
	"errors"
	"fmt"
	//"gitee.com/moorewqk/antcom/server/admin"
)

func Printfln(format string, args ...interface{}) string {
	mess := fmt.Sprintf(format, args...)
	fmt.Println(mess)
	return mess
}

/*
info日志格式化打印
*/
func PfOk(format string, args ...interface{}) string {
	mess := Printfln(format, args...)
	//admin.GV_LOG.Info(mess)
	return mess
}

/*
error日志格式化打印
*/

func PfFail(format string, args ...interface{}) string {
	mess := Printfln(format, args...)
	//admin.GV_LOG.Error(mess)
	return mess
}

func OutEr(format string, args ...interface{}) error {
	mess := Printfln(format, args...)
	err := errors.New(mess)
	return err
}
