package model

import "gwsee.com.api/app/common"

type Classify struct {
	common.DbColumn
	ClassifyId   uint64 `db:"classify_id" form:"classifyid" json:"classifyid"`
	ClassifyName string `db:"classify_name" form:"name" json:"name"`
	ClassifyDesc string `db:"classify_desc" form:"desc" json:"desc"`
	ClassifySort int    `db:"classify_sort" form:"sort" json:"sort"`
	ClassifyEnd  int    `db:"classify_end" form:"end" json:"end"`
	ClassifyPid  uint64 `db:"classify_pid" form:"pid" json:"pid"`
}

type ClassifyTree struct {
	Classify
	Children []ClassifyTree `form:"children" json:"children" ` //不大写就外部不收数据
}

type ClassifyAttribute struct {
	common.DbColumn
	ClassifyId    uint64 `db:"classify_id" form:"classifyid" json:"classifyid"`
	AttributeId   uint64 `db:"attribute_id" form:"attributeid" json:"attributeid"`
	AttributeName string `db:"attribute_name" form:"name" json:"name"`
	AttributeSort int    `db:"attribute_sort" form:"sort" json:"sort"`
	IsAlias       int    `db:"is_alias" form:"alias" json:"alias"`
	IsColor       int    `db:"is_color" form:"color" json:"color"`
	IsEnumeration int    `db:"is_enumeration" form:"enumeration" json:"enumeration"`
	IsInput       int    `db:"is_input" form:"input" json:"input"`
	IsCrux        int    `db:"is_crux" form:"crux" json:"crux"`
	IsSale        int    `db:"is_sale" form:"sale" json:"sale"`
	IsSearch      int    `db:"is_search" form:"search" json:"search"`
	IsMust        int    `db:"is_must" form:"must" json:"must"`
	IsMultiple    int    `db:"is_multiple" form:"multiple" json:"multiple"`
}
type ClassifyAttributeItem struct {
	common.DbColumn
	ClassifyId  uint64 `db:"classify_id"  json:"classifyid"`
	AttributeId uint64 `db:"attribute_id"  json:"attributeid"`
	ItemId      uint64 `db:"item_id"  json:"itemid"`
	ItemName    string `db:"item_name" form:"name" json:"name"`
	ItemSort    int    `db:"item_sort"  json:"sort"`
}

//用于获取单个类目 与其对应的类目属性

type ClassifyAndAttribute struct {
	Classify
	Items []*ClassifyAttributeAndItem `form:"items" json:"items"`
}
type ClassifyAttributeAndItem struct {
	ClassifyAttribute
	Items []*ClassifyAttributeItem `form:"items" json:"items"`
}
type ClassifyBrand struct {
	common.DbColumn
	Id         uint64 `db:"id" form:"id" json:"id"`
	ClassifyId uint64 `db:"classify_id" form:"classifyid" json:"classifyid"`
	BrandId    uint64 `db:"brand_id" form:"brandid" json:"brandid"`
	BrandName  string `db:"brand_name" form:"name" json:"name"`
	BrandLogo  string `db:"brand_logo" form:"logo" json:"logo"`
}
type ClassifyBrandIds struct {
	ClassifyId uint64   `db:"classify_id" form:"classifyid" json:"classifyid"`
	Ids        []string `form"ids" json:"ids"`
}

//模板brandid  应该是类目品牌的  ClassifyBrand的ID 不是 另外的那个brandid
type ClassifyTemplate struct {
	common.DbColumn
	TemplateId   uint64 `db:"template_id" form:"templateid" json:"templateid"`
	TemplateName string `db:"template_name" form:"name" json:"name"`
	BrandId      uint64 `db:"brand_id" form:"brandid" json:"brandid"` //类目品牌（classify_brand)表的ID 有类目的属性了
}
type ClassifyTemplateAttribute struct {
	common.DbColumn
	Id             uint64 `db:"id" form:"id" json:"id"`
	TemplateId     uint64 `db:"template_id" form:"templateid" json:"templateid"`
	AttributeId    uint64 `db:"attribute_id" form:"attributeid" json:"attributeid"`
	AttributeName  string `db:"attribute_name" form:"name" json:"name"`
	IsEdit         int    `db:"is_edit" form:"edit" json:"edit"`
	AttributeValue string `db:"attribute_value" form:"value" json:"value"`
}
type ClassifyTemplateAndAttribute struct {
	ClassifyTemplate
	Attribute []ClassifyTemplateAttribute `form:"attribute" json:"attribute"`
}

//获取一个模板属性的基本信息
type ClassifyTemplateAttributeBase struct {
	ClassifyTemplateAttribute
	ClassifyAttributeAndItem
}
type ClassifyTemplateInfo struct {
	ClassifyTemplate
	Attribute []*ClassifyTemplateAttributeBase `form:"attribute" json:"attribute"`
}
