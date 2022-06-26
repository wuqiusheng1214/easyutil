/*
@Time       : 2022/1/10
@Author     : wuqiusheng
@File       : rw_split.go
@Description: 读写分离,单协程写,多协程读
*/
package easyutil

//读写分离
type RWSplit struct {
	readIndex uint8
	RWList    [2]interface{}
}

func NewRWSplit(v interface{}) *RWSplit {
	return &RWSplit{
		RWList: [2]interface{}{v},
	}
}

//多协程读
func (r *RWSplit) Get() interface{} {
	return r.RWList[r.readIndex]
}

//单协程写
func (r *RWSplit) Set(v interface{}) {
	index := (r.readIndex + 1) % 2
	r.RWList[index] = v
	oldindex := r.readIndex
	r.readIndex = index
	r.RWList[oldindex] = nil
}
