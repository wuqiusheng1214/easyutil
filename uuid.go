/*
@Time       : 2022/1/9
@Author     : wuqiusheng
@File       : uuid.go
@Description: 唯一ID
*/
package easyutil

import (
	"sync/atomic"
	"time"
	"unsafe"
)

type UUID struct {
	Timestamp int32
	ServerId  uint16
	Index     uint16
}

var uuidIndex uint32 = 0

//单服同一秒请求超过65535或重复
func GetUUID(sid int32) int64 {
	index := atomic.AddUint32(&uuidIndex, 1)
	index %= 65536
	uuid := UUID{int32(time.Now().Unix()), uint16(sid), uint16(index)}
	return *((*int64)(unsafe.Pointer(&uuid)))
}

func GetUUIDStr(sid int32) string {
	return Itoa(GetUUID(sid))
}

func GetServerIdByUUIDStr(sessionStr string) int32 {
	session := Atoi64(sessionStr)
	p := (*UUID)(unsafe.Pointer(&session))
	return int32(p.ServerId)
}

func GetTimestampByUUIDStr(sessionStr string) int64 {
	session := Atoi64(sessionStr)
	p := (*UUID)(unsafe.Pointer(&session))
	return int64(p.Timestamp)
}

//loginServer使用
type GID struct {
	Timestamp uint32 //时间戳
	ServerId  uint8 //255 loginServerId 1-255
	CellId 	  uint8 //255  分片ID 0-255 100w 2亿5千万用户
	Index     uint16 //自增ID
}

var gidIndex uint32 = 0

var startTimeStamp uint32 = 1651334400 //2022.5.1
//单服同一秒请求超过65535或重复
func GenGid(sid uint8,cellId uint8) int64 {
	index := atomic.AddUint32(&gidIndex, 1)
	index %= 65536
	gid := GID{uint32(time.Now().Unix())-startTimeStamp, sid, cellId,uint16(index)}
	return *((*int64)(unsafe.Pointer(&gid)))
}

func GetGidInfo(gid int64) *GID {
	return (*GID)(unsafe.Pointer(&gid))
}
