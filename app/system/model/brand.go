package model

import "gwsee.com.api/app/common"

type BrandGroup struct {
	common.DbColumn
	Id   uint64 `db:"id" form:"id" json:"id"`
	Name string `db:"name" form:"name" json:"name"`
	Desc string `db:"desc" form:"desc" json:"desc"`
	Sort int    `db:"sort" form:"sort" json:"sort"`
}

//用于分组显示各组下面的品牌
type GroupItems struct {
	BrandGroup
	Children []Brand `form:"children" json:"children"`
}
type Brand struct {
	common.DbColumn
	BrandId      uint64 `db:"brand_id" form:"brandid" json:"brandid"`
	BrandName    string `db:"brand_name" form:"name" json:"name"`
	BrandLogo    string `db:"brand_logo" form:"logo" json:"logo"`
	BrandSort    int    `db:"brand_sort" form:"sort" json:"sort"`
	BrandDesc    string `db:"brand_desc" form:"desc" json:"desc"`
	BrandWebsite string `db:"brand_website" form:"website" json:"website"`
	BrandStory   string `db:"brand_story" form:"story" json:"story"`
	GroupId      uint64 `db:"group_id" form:"id" json:"id"`
}

type BrandAndGroup struct {
	Brand
	Name string `db:"name" form:"groupname" json:"groupname"`
}
