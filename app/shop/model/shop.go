package model

import "gwsee.com.api/app/common"

type Shop struct {
	common.DbColumn
	ShopId       uint64 `db:"shop_id"   form:"shopid" json:"shopid"`
	ShopName     string `db:"shop_name"  form:"shopname" json:"shopname"`
	ShopLogo     string `db:"shop_logo"   form:"logo" json:"logo"`
	ShopSn       string `db:"shop_sn"   form:"shopsn" json:"shopsn"`
	ShopPid      uint64 `db:"shop_pid"   form:"pid" json:"pid"`
	ShopAddr     string `db:"shop_addr"  form:"addr" json:"addr"`
	ShopBreif    string `db:"shop_breif" form:"breif" json:"breif"`
	ShopFullname string `db:"shop_fullname"  form:"fullname" json:"fullname"`
	ShopTel      string `db:"shop_tel"  form:"tel" json:"tel"`
	ShopOwner    uint64 `db:"shop_owner"  form:"owner" json:"owner"`
}
