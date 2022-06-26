/*
@Time       : 2022/1/3
@Author     : wuqiusheng
@File       : broadcast.go
@Description: 广播组播
*/
package easyutil

import "sync"

//组播
type IMulticast interface {
	IsInGroup(group string) bool //所属组 用于过滤消息
}

//广播消息
type BroadcastMessage struct {
	C     chan struct{}                   //chan 通知
	msg   interface{}                     //广播消息
	msgId uint8                           //消息id
	fun   func(multicast IMulticast) bool //过滤函数
}

//获取消息
//multicast 过滤组播消息
func (r *BroadcastMessage) GetMsg(multicast IMulticast) interface{} {
	if r.fun != nil && !r.fun(multicast) {
		return nil
	}
	return r.msg
}

//广播
type Broadcast struct {
	sync.RWMutex
	msgId   uint8
	count   uint8
	msgList []*BroadcastMessage //循环消息 256个，最多同时广播256个消息，不会被覆盖，理论上够用，不做删除
}

//新建广播队列 最小数量为2
func NewBroadcast(count uint8) *Broadcast {
	if count < 2 {
		return nil
	}
	r := &Broadcast{
		msgList: make([]*BroadcastMessage, count),
		count:   count,
	}
	c := make(chan struct{})
	r.msgList[r.msgId] = &BroadcastMessage{C: c}
	return r
}

//监听指定消息
func (r *Broadcast) GetNextMsg(msg *BroadcastMessage) *BroadcastMessage {
	id := (msg.msgId + 1) % r.count
	r.RLock()
	msg = r.msgList[id]
	r.RUnlock()
	return msg
}

//监听指定消息
func (r *Broadcast) GetNewMsg() *BroadcastMessage {
	r.RLock()
	msg := r.msgList[r.msgId]
	r.RUnlock()
	return msg
}

//广播消息
func (r *Broadcast) Broadcast(msg interface{}, fun func(multicast IMulticast) bool) {
	if msg == nil {
		return
	}
	c := make(chan struct{})
	r.Lock()
	gmsg := r.msgList[r.msgId]
	r.msgId = (r.msgId + 1) % r.count
	r.msgList[r.msgId] = &BroadcastMessage{C: c, msgId: r.msgId} //容量256，GetNextMsg阻断，拿不到老数据
	r.Unlock()
	gmsg.msg = msg
	gmsg.fun = fun
	close(gmsg.C)
}

func (r *Broadcast) Multicast(group string, msg interface{}) {
	r.Broadcast(msg, func(multicast IMulticast) bool {
		if multicast == nil {
			return false
		}
		return multicast.IsInGroup(group)
	})
}
