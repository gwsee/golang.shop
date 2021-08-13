package model

import "gwsee.com.api/app/common"

type ShopDepart struct {
	common.DbColumn
	DepartId   uint64 `db:"depart_id"  form:"departid" json:"departid"`
	DepartPid  uint64 `db:"depart_pid"  form:"pid" json:"pid"`
	DepartName string `db:"depart_name"  form:"name" json:"name"`
	DepartSort uint64 `db:"depart_sort"  form:"sort" json:"sort"`
	DepartDesc string `db:"depart_desc"  form:"desc" json:"desc"`
	ShopSn     string `db:"shop_sn" form:"shopsn" json:"shopsn"`
}
type DepartShop struct {
	ShopDepart
	ShopName string `db:"shop_name" form:"shopname" json:"shopname"`
}
type DepartTree struct {
	DepartShop
	Children []DepartTree `form:"children" json:"children"`
}
