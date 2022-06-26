/*
@Time       : 2022/1/4
@Author     : wuqiusheng
@File       : encrypt.go
@Description: 加解密算法
*/
package easyutil

import "encoding/base64"

var encryptKey = [16]byte{253, 1, 56, 52, 62, 176, 42, 138, 32, 123, 43, 12, 167, 89, 231, 222}

//加密,相对base64快1倍
func Encrypt(buf []byte) []byte {
	buflen := len(buf)
	for i := 0; i < buflen; i++ {
		n := byte(i%7 + 1)                       //移位长度(1-7)
		b := (buf[i] << n) | (buf[i] >> (8 - n)) // 向左循环移位
		buf[i] = b ^ encryptKey[i%16]
	}
	return buf
}

//解密
func Decrypt(buf []byte) []byte {
	buflen := len(buf)
	for i := 0; i < buflen; i++ {
		b := buf[i] ^ encryptKey[i%16]
		n := byte(i%7 + 1)                 //移位长度(1-7)
		buf[i] = (b >> n) | (b << (8 - n)) // 向右循环移位
	}
	return buf
}

const base64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"

var coder = base64.NewEncoding(base64Table)

//Base64编码
func Base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

//Base64解码
func Base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}
