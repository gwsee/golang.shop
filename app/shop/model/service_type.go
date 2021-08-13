package model

import "gwsee.com.api/app/common"

type ShopServiceType struct {
	common.DbColumn
	TypeId      uint64 `db:"type_id"    form:"typeid" json:"typeid"`
	TypeName    string `db:"type_name"  form:"name" json:"name"`
	TypeDesc    string `db:"type_desc"  form:"desc" json:"desc"`
	TypeMinutes uint64 `db:"type_minutes" form:"minutes" json:"minutes"`
	ShopSn      string `db:"shop_sn" form:"shopsn" json:"shopsn"`
}
