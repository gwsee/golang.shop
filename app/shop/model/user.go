package model

import (
	"gwsee.com.api/app/common"
	user "gwsee.com.api/app/user/model"
)

type ShopUser struct {
	common.DbColumn
	Id       uint64 `db:"id"    form:"id" json:"id"`
	UserId   uint64 `db:"user_id"    form:"userid" json:"userid"`
	RoleIds  string `db:"role_ids"   form:"roleids" json:"roleids"`
	DepartId uint64 `db:"depart_id"     form:"departid" json:"departid"`
	ShopSn   string `db:"shop_sn"   form:"shopsn" json:"shopsn"`
}

type ShopUserDetail struct {
	ShopUser
	ShopName   string `db:"shop_name" form:"shopname" json:"shopname"`
	DepartName string `db:"depart_name" form:"departname" json:"departname"`
	RoleNames  string `form:"rolenames" json:"rolenames"`
	user.UserBase
	user.UserData
}
