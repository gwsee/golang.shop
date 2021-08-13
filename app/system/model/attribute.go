package model

import "gwsee.com.api/app/common"

type Attribute struct {
	common.DbColumn
	AttributeId   uint64 `db:"attribute_id" form:"attributeid" json:"attributeid"`
	AttributeName string `db:"attribute_name" form:"name" json:"name"`
	AttributeDesc string `db:"attribute_desc" form:"desc" json:"desc"`
	AttributeSort int    `db:"attribute_sort" form:"sort" json:"sort"`
}

type AttributeItem struct {
	common.DbColumn
	AttributeId uint64 `db:"attribute_id" form:"attributeid" json:"attributeid"`
	ItemId      uint64 `db:"item_id" form:"itemid" json:"itemid"`
	ItemName    string `db:"item_name" form:"name" json:"name"`
	ItemSort    int    `db:"item_sort" form:"sort" json:"sort"`
}

type AttributeAndItem struct {
	Attribute
	Items []string `form:"items" json:"items"`
}
