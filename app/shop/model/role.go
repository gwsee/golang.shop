package model

import "gwsee.com.api/app/common"

type ShopRole struct {
	common.DbColumn
	RoleId   uint64 `db:"role_id"    form:"roleid" json:"roleid"`
	RoleName string `db:"role_name"  form:"name" json:"name"`
	RoleDesc string `db:"role_desc"  form:"desc" json:"desc"`
	MenuIds  string `db:"menu_ids"   form:"ids" json:"ids"`
	ShopSn   string `db:"shop_sn"   form:"shopsn" json:"shopsn"`
}
