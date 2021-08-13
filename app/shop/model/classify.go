package model

import "gwsee.com.api/app/common"

type ShopClassify struct {
	common.DbColumn
	ID           uint64 `db:"id"    form:"id" json:"id"`
	ClassifyId   string `db:"classify_id"  form:"classifyid" json:"classifyid"`
	BrandId      string `db:"brand_id"  form:"brandid" json:"brandid"`
	ClassifyName string `db:"classify_name"  form:"classifyname" json:"classifyname"`
	BrandName    string `db:"brand_name"  form:"brandname" json:"brandname"`
	ShopSn       string `db:"shop_sn"   form:"shopsn" json:"shopsn"`
}
