/*
@Time       : 2022/1/2
@Author     : wuqiusheng
@File       : error.go
@Description: 基础错误码
*/
package easyutil

var (
	ErrMsgPackPack   = &Error{1, "msgpack打包错误"}
	ErrMsgPackUnPack = &Error{2, "msgpack解析错误"}
	ErrPBPack        = &Error{3, "pb打包错误"}
	ErrPBUnPack      = &Error{4, "pb解析错误"}
	ErrJsonPack      = &Error{5, "json打包错误"}
	ErrJsonUnPack    = &Error{6, "json解析错误"}
	ErrGobPack       = &Error{7, "gob打包错误"}
	ErrGobUnPack     = &Error{8, "gob解析错误"}

	ErrHttpRequest = &Error{9, "http请求错误"}
)

type Error struct {
	Id  uint16
	Str string
}

func (r *Error) Error() string {
	return r.Str
}
