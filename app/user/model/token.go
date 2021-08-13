package model

import (
	"gwsee.com.api/app/common"
)

// 用于获取个人基本信息的结构体
type UserToken struct {
	common.DbColumn
	TokenId uint64 `db:"token_id" form:"tokenid" json:"tokenid"`
	Token   string `db:"token"    form:"token" json:"token"`
	User    string `db:"user"     form:"user" json:"user"`
	Num     uint64 `db:"num"      form:"num" json:"num"`
	Shop    string `db:"shop"     form:"shop" json:"shop"`
	ShopSn  string `db:"shop_sn"     form:"shopsn" json:"shopsn"`
}
