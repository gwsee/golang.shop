package model

import "gwsee.com.api/app/common"

type SystemSettingType struct {
	common.DbColumn
	TypeId   uint64 `db:"type_id"    form:"typeid" json:"typeid"`
	TypeName string `db:"type_name"  form:"name" json:"name"`
	Type     string `db:"type"  form:"type" json:"type"`
	TypeDesc string `db:"type_desc"  form:"desc" json:"desc"`
}
