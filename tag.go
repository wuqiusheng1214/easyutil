/*
@Time       : 2022/4/13
@Author     : wuqiusheng
@File       : tag.go
@Description: 获取结构体标签
*/
package easyutil

import (
	"reflect"
)

func GetTagList(v interface{}, tag string) []string {
	ptype := reflect.TypeOf(v).Elem()
	tagList := make([]string, 0, ptype.NumField())
	for i := 0; i < ptype.NumField(); i++ {
		tag := ptype.Field(i).Tag.Get(tag)
		tagList = append(tagList, tag)
	}
	return tagList
}

func GetTagMap(v interface{}, tag string) map[int]string {
	ptype := reflect.TypeOf(v).Elem()
	tagMap := make(map[int]string)
	for i := 0; i < ptype.NumField(); i++ {
		tag := ptype.Field(i).Tag.Get(tag)
		tagMap[i] = tag
	}
	return tagMap
}
