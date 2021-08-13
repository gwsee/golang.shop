package model

import "gwsee.com.api/app/common"

type Log struct {
	common.DbColumn
	LogId         uint64 `db:"log_id"   form:"log_id" json:"id"`
	UserId        uint64 `db:"user_id"  form:"user_id" json:"userid"`
	UserName      string `db:"user_name"  form:"user_name" json:"name"`
	ActivityId    uint64 `db:"activity_id" form:"activity_id" json:"activityid"`
	ActivityTitle string `db:"activity_title" form:"activity_title" json:"title"`
	ActivityPrice string `db:"activity_price" form:"activity_price" json:"price"`
	OrderSn       string `db:"order_sn"   form:"order_sn" json:"sn"`
	ShopSn        string `db:"shop_sn"   form:"shop_sn" json:"shopsn"`
}
