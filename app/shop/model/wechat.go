package model

import "gwsee.com.api/app/common"

type ShopWechat struct {
	common.DbColumn
	WechatId         uint64 `db:"wechat_id"    form:"wechatid" json:"wechatid"`
	WechatAppKey     string `db:"wechat_appkey"   form:"wechatappkey" json:"wechatappkey"`
	WechatAppSecrect string `db:"wechat_appsecrect"     form:"wechatappsecrect" json:"wechatappsecrect"`
	WechatAccount    string `db:"wechat_account"     form:"wechataccount" json:"wechataccount"`
	ShopSn           string `db:"shop_sn"   form:"shopsn" json:"shopsn"`
}
