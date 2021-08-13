package model

import "gwsee.com.api/app/common"

type Grade struct {
	common.DbColumn
	GradeId    uint64 `db:"grade_id"    form:"gradeid" json:"gradeid"`
	GradeName  string `db:"grade_name"  form:"name" json:"name"`
	GradeDesc  string `db:"grade_desc"  form:"desc" json:"desc"`
	GradeTime  uint64 `db:"grade_time"   form:"time" json:"time"`
	GradePrice uint64 `db:"grade_price"   form:"price" json:"price"`
	ShopSn     string `db:"shop_sn"   form:"shopsn" json:"shopsn"`
	Id         uint64 `db:"id"  form:"id" json:"id"`
}
type GradeList struct {
	Grade
	Step uint64 `db:"step"  form:"step" json:"step"`
	Logo string `db:"logo"   form:"logo" json:"logo"`
	Pic  string `db:"pic"   form:"pic" json:"pic"`
}
