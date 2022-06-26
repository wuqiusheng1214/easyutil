/*
@Time       : 2022/1/14
@Author     : wuqiusheng
@File       : stable_hash.go
@Description: 稳定hash
*/
package easyutil

import "reflect"

//字符串稳定hash值
func StableHashCode(text string) int {
	var hash = 0
	for _, v := range text {
		hash = hash*3 + int(v)
	}
	return hash
}

//msg必须是指针
func GetMessageId(msg interface{}) uint16 {
	name := reflect.TypeOf(msg).Elem().Name()
	return uint16(StableHashCode(name) & 0xFFFF)
}

func GetHashCode(name string) uint16 {
	return uint16(StableHashCode(name) & 0xFFFF)
}
