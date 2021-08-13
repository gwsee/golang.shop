package model

import "gwsee.com.api/app/common"

type Classify struct {
	common.DbColumn
	ClassifyId    uint64 `db:"classify_id" form:"classifyid" json:"classifyid"`
	ClassifyName  string `db:"classify_name"  form:"name" json:"name"`
	ClassifyAlias string `db:"classify_alias"  form:"alias" json:"alias"`
	ClassifyPic   string `db:"classify_pic"  form:"pic" json:"pic"`
	ClassifySort  uint64 `db:"classify_sort"  form:"sort" json:"sort"`
	ClassifyIds   string `db:"classify_ids"  form:"ids" json:"ids"`
	ShopSn        string `db:"shop_sn"   form:"shopsn" json:"shopsn"`
}
