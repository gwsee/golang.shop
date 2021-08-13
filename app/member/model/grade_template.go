package model

import "gwsee.com.api/app/common"

type GradeTemplate struct {
	common.DbColumn
	Id     uint64 `db:"id"    form:"id" json:"id"`
	Name   string `db:"name"  form:"name" json:"name"`
	Step   uint64 `db:"step"  form:"step" json:"step"`
	Logo   string `db:"logo"   form:"logo" json:"logo"`
	Pic    string `db:"pic"   form:"pic" json:"pic"`
	ShopSn string `db:"shop_sn"   form:"shopsn" json:"shopsn"`
}
