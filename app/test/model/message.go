package model

import (
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/test/model/child"
)

type Message struct {
	common.DbColumn
	MessageId    uint64 `db:"message_id" form:"messageid" json:"messageid"`
	ShopSn       string `db:"shop_sn"  form:"shop_sn" json:"shop_sn"`
	Message      string `db:"message"  form:"message" json:"message"`
	SenderId     uint64 `db:"sender_id"  form:"senderid" json:"senderid"`
	ReceiverId   uint64 `db:"receiver_id"  form:"receiverid" json:"receiverid"`
	DelaySeconds uint64 `db:"delay_seconds"  form:"seconds" json:"seconds"` //代表是延迟多少秒后才插入数据的
}

//消息加载
/**
message_id 消息ID  有值的时候根据此值结合type进行数据查询，没值的时候 查询最新消息
message  消息内容  可以通过message来进行模糊查询
type  消息类型  all 全部消息（相当于刷新）;before 查询message_id之前的消息；after查询message_id之后的消息
state 消息状态  1未读  2已读
*/
type MessageLoad struct {
	MessageId   uint64 `form:"message_id" json:"message_id"`
	Message     string `form:"message" json:"message"`
	MessageType string `form:"type" json:"type"`
	State       string `form:"state" json:"state"`
}

type T0 struct {
	B0 string `res:"b0" json:"b0"`
}

type T1 struct {
	T0
	A1 string `res:"a1" json:"a1"`
	A2 string `res:"a2" json:"a2"`
	A3 string `res:"a3" json:"a3"`
	A4 int    `res:"a4" json:"a4"`
	A5 int    `res:"a5" json:"a5"`
	A6 float64    `res:"a6" json:"a6"`
	child.T2
}