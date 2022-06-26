/*
@Time       : 2022/1/2
@Author     : wuqiusheng
@File       : pack.go
@Description: gob json pb msgpack 打包解包
*/
package easyutil

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/vmihailenco/msgpack"
)

func GobPack(msg interface{}) ([]byte, error) {
	if msg == nil {
		return nil, ErrGobPack
	}
	var bio bytes.Buffer
	enc := gob.NewEncoder(&bio)
	err := enc.Encode(msg)
	if err != nil {
		return nil, ErrGobPack
	}
	return bio.Bytes(), nil
}

func GobUnPack(data []byte, msg interface{}) error {
	if data == nil || msg == nil {
		return ErrGobUnPack
	}
	bio := bytes.NewBuffer(data)
	enc := gob.NewDecoder(bio)
	err := enc.Decode(msg)
	if err != nil {
		return ErrGobUnPack
	}
	return nil
}

func JsonUnPack(data []byte, msg interface{}) error {
	if data == nil || msg == nil {
		return ErrJsonUnPack
	}
	err := json.Unmarshal(data, msg)
	if err != nil {
		return ErrJsonUnPack
	}
	return nil
}

func JsonPack(msg interface{}) ([]byte, error) {
	if msg == nil {
		return nil, ErrJsonPack
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, ErrJsonPack
	}
	return data, nil
}

func MsgPackUnPack(data []byte, msg interface{}) error {
	if data == nil || msg == nil {
		return ErrMsgPackUnPack
	}
	err := msgpack.Unmarshal(data, msg)
	if err != nil {
		return ErrMsgPackUnPack
	}
	return nil
}

func MsgPackPack(msg interface{}) ([]byte, error) {
	if msg == nil {
		return nil, ErrMsgPackPack
	}
	data, err := msgpack.Marshal(msg)
	if err != nil {
		return nil, ErrMsgPackPack
	}
	return data, err
}

func PBUnPack(data []byte, msg interface{}) error {
	if data == nil || msg == nil {
		return ErrPBUnPack
	}
	err := proto.Unmarshal(data, msg.(proto.Message))
	if err != nil {
		return ErrPBUnPack
	}
	return nil
}

func PBPack(msg interface{}) ([]byte, error) {
	if msg == nil {
		return nil, ErrPBPack
	}
	data, err := proto.Marshal(msg.(proto.Message))
	if err != nil {
		return nil, ErrPBPack
	}
	return data, nil
}
