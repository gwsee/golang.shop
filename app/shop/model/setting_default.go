package model

import (
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/system/model"
)

type ShopSettingDefault struct {
	common.DbColumn
	DefaultId    uint64 `db:"default_id"    form:"defaultid" json:"defaultid"`
	DefaultKey   string `db:"default_key"  form:"key" json:"key"`
	DefaultValue string `db:"default_value"  form:"value" json:"value"`
	ShopSn       string `db:"shop_sn"   form:"shopsn" json:"shopsn"`
}
type ShopSettingDefaultData struct {
	model.SystemSettingType
	DefaultValue string `form:"value" json:"value"`
}
