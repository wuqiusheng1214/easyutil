/*
@Time       : 2022/1/3
@Author     : wuqiusheng
@File       : md5.go
@Description: md5函数包装
*/
package easyutil

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Str(s string) string {
	return MD5Bytes([]byte(s))
}

func MD5Bytes(s []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(s)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func MD5File(path string) string {
	data, err := ReadFile(path)
	if err != nil {
		Println("calc md5 failed path:%v", path)
		return ""
	}
	return MD5Bytes(data)
}
