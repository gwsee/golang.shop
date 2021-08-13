package model

import "gwsee.com.api/app/common"

type GradeInterests struct {
	common.DbColumn
	InterestsId   uint64 `db:"interests_id" form:"interestsid" json:"interestsid"`
	ClassifyId    uint64 `db:"classify_id" form:"classifyid" json:"classifyid"`
	GradeId       uint64 `db:"grade_id" form:"gradeid" json:"gradeid"`
	ClassifyName  string `db:"classify_name"  form:"name" json:"name"`
	InterestsUnit string `db:"interests_unit"  form:"unit" json:"unit"`
	ShopSn        string `db:"shop_sn"   form:"shopsn" json:"shopsn"`
}
