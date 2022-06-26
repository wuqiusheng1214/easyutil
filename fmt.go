/*
@Time       : 2021/12/31
@Author     : wuqiusheng
@File       : fmt.go
@Description: 重写fmt
*/
package easyutil

import "fmt"

func Print(a ...interface{}) (int, error) {
	return fmt.Print(a...)
}
func Println(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}
func Printf(format string, a ...interface{}) (int, error) {
	return fmt.Printf(format, a...)
}
func Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}
func Sprint(a ...interface{}) string {
	return fmt.Sprint(a...)
}
