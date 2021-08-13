package model

import (
	"gwsee.com.api/app/common"
	model2 "gwsee.com.api/app/system/model"
)

type Commons struct {
	common.DbColumn
	ClassifyId      uint64  `db:"classify_id"  form:"classifyid" json:"classifyid"` //类目ID
	ClassifyName    string  `db:"classify_name" form:"name" json:"name"`            //类目的名字组合
	BrandId         uint64  `db:"brand_id"  form:"brandid" json:"brandid"`          //品牌ID
	CommonsId       uint64  `db:"commons_id" form:"commonsid" json:"commonsid"`     // 商品ID
	CommonsService  uint64  `db:"commons_service" form:"service" json:"service"`    //服务类型
	CommonsSign     uint64  `db:"commons_sign" form:"sign" json:"sign"`             //报名类型
	CommonsTitle    string  `db:"commons_title" form:"title" json:"title"`
	CommonsSort     int     `db:"commons_sort" form:"sort" json:"sort"`
	CommonsPrice    float64 `db:"commons_price" form:"price" json:"price"` //市场价格
	CommonsCost     float64 `db:"commons_cost" form:"cost" json:"cost"`    //成本价格
	CommonsStock    int     `db:"commons_stock" form:"stock" json:"stock"`
	CommonsClassify int     `db:"commons_classify" form:"classify" json:"classify"` //商品分类
	CommonsDesc     string  `db:"commons_desc" form:"desc" json:"desc"`
	CommonsCode     string  `db:"commons_code" form:"code" json:"code"` //条形码
	CommonsSn       string  `db:"commons_sn" form:"sn" json:"sn"`       //编码
	CommonsType     string  `db:"commons_type" form:"type" json:"type"`
	CommonsRemark   string  `db:"commons_remark" form:"remark" json:"remark"`
	CommonsWarn     int     `db:"commons_warn" form:"warn" json:"warn"`
	CommonsTrans    string  `db:"commons_trans" form:"trans" json:"trans"`
	CommonsInvented int     `db:"commons_invented" form:"invented" json:"invented"` //虚拟人数
	CommonsSale     int     `db:"commons_sale" form:"sale" json:"sale"`             //已售订单交易成功
	CommonsClicked  int     `db:"commons_clicked" form:"clicked" json:"clicked"`    //点击量
	CommonsOrder    int     `db:"commons_order" form:"order" json:"order"`          //订单数 下了单的
	CommonsPeople   int     `db:"commons_people" form:"people" json:"people"`       //引流人数
	ShopSn          string  `db:"shop_sn" form:"shopsn" json:"shopsn"`
}

type CommonsAttribute struct {
	common.DbColumn
	Id              uint64 `db:"id" form:"id" json:"id"`
	AttributeId     uint64 `db:"attribute_id" form:"attributeid" json:"attributeid"`
	AttributeName   string `db:"attribute_name" form:"name" json:"name"`
	AttributeValues string `db:"attribute_values" form:"values" json:"values"`
	AttributeShow   string `db:"attribute_show" form:"show" json:"show"`
	CommonsId       uint64 `db:"commons_id" form:"commonsid" json:"commonsid"`
	ShopSn          string `db:"shop_sn" form:"shopsn" json:"shopsn"`
}
type CommonsFile struct {
	common.DbColumn
	CommonsId  uint64 `db:"commons_id"  form:"commonsid" json:"commonsid"`
	ClassifyId uint64 `db:"classify_id"  form:"classifyid" json:"classifyid"`
	BrandId    uint64 `db:"brand_id"  form:"brandid" json:"brandid"`
	FileId     uint64 `db:"file_id" form:"fileid" json:"fileid"`
	FileBanner string `db:"file_banner"  form:"banner" json:"banner"`
	FileVideo  string `db:"file_video"   form:"video" json:"video"`
	FileLogo   string `db:"file_logo"   form:"logo" json:"logo"`
	FileDesc   string `db:"file_desc"   form:"desc" json:"desc"`
	ShopSn     string `db:"shop_sn"   form:"shopsn" json:"shopsn"`
}

type CommonsItem struct {
	common.DbColumn
	ItemId            uint64  `db:"item_id" form:"itemid" json:"itemid"`
	ItemStock         int     `db:"item_stock" form:"stock" json:"stock"`
	ItemCost          float64 `db:"item_cost" form:"cost" json:"cost"`       //进货价
	ItemMarket        float64 `db:"item_market" form:"market" json:"market"` //市场价
	ItemPrice         float64 `db:"item_price" form:"price" json:"price"`    //销售价格
	ItemSn            string  `db:"item_sn" form:"sn" json:"sn"`
	ItemCode          string  `db:"item_code" form:"code" json:"code"`
	CommonsId         uint64  `db:"commons_id" form:"commonsid" json:"commonsid"`
	CommonsAttributes string  `db:"commons_attributes" form:"attributes" json:"attributes"`
	ItemSale          int     `db:"item_sale" form:"sale" json:"sale"`
	ItemOrder         int     `db:"item_order" form:"order" json:"order"`
	ShopSn            string  `db:"shop_sn" form:"shopsn" json:"shopsn"`
}

type CommonsData struct {
	Commons
	Attribute []*CommonsAttribute `form:"attribute" json:"attribute"`
	File      *CommonsFile        `form:"file" json:"file"`
	Item      []*CommonsItem      `form:"item" json:"item"`
}

// 商品的属性组(各个属性对应的数据） -- 输出 （form)
type CommonsAttributeForm struct {
	model2.ClassifyAttributeAndItem
	Value interface{} `form:"value" json:"value"`
}
