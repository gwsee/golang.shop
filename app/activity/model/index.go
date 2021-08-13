package model

import "gwsee.com.api/app/common"

type Activity struct {
	common.DbColumn
	ActivityId        uint64  `db:"activity_id" form:"id" json:"id"`
	ActivityTitle     string  `db:"activity_title"  form:"title" json:"title"`
	ActivityType      string  `db:"activity_type"  form:"type" json:"type"`
	ActivityDesc      string  `db:"activity_desc"  form:"desc" json:"desc"`
	ActivityBegin     int     `db:"activity_begin"  form:"begin" json:"begin"`
	ActivityEnd       int     `db:"activity_end"  form:"end" json:"end"`
	ActivityDetail    string  `db:"activity_detail"  form:"detail" json:"detail"`
	ActivityRecommend int     `db:"activity_recommend"  form:"recommend" json:"recommend"`
	ActivityRules     string  `db:"activity_rules"  form:"rules" json:"rules"`
	ActivitySort      int     `db:"activity_sort"  form:"sort" json:"sort"`
	ActivityStock     int     `db:"activity_stock"  form:"stock" json:"stock"`
	ActivityInvented  int     `db:"activity_invented"  form:"invented" json:"invented"`
	ActivityRemark    string  `db:"activity_remark"  form:"remark" json:"remark"`
	ActivitySale      int     `db:"activity_sale"  form:"sale" json:"sale"`
	ActivityClicked   int     `db:"activity_clicked"  form:"clicked" json:"clicked"`
	ActivityOrder     int     `db:"activity_order"  form:"order" json:"order"`
	ActivityPeople    int     `db:"activity_people"  form:"people" json:"people"`
	ActivityLimit     int     `db:"activity_limit"  form:"limit" json:"limit"`
	ActivityPrice     float64 `db:"activity_price"  form:"price" json:"price"`
	ActivityNum       int     `db:"activity_num"  form:"num" json:"num"`
	ShopSn            string  `db:"shop_sn"   form:"shopsn" json:"shopsn"`
}
type ActivityFile struct {
	common.DbColumn
	FileId     uint64 `db:"file_id" form:"fileid" json:"fileid"`
	FileBanner string `db:"file_banner"  form:"banner" json:"banner"`
	FileVideo  string `db:"file_video"   form:"video" json:"video"`
	FileLogo   string `db:"file_logo"   form:"logo" json:"logo"`
	FileDesc   string `db:"file_desc"   form:"desc" json:"desc"`
	ShopSn     string `db:"shop_sn"   form:"shopsn" json:"shopsn"`
	ActivityId uint64 `db:"activity_id" form:"activityid" json:"activityid"`
}
type ActivityItem struct {
	common.DbColumn
	ItemId     uint64 `db:"item_id"    form:"item_id" json:"item_id"`
	ItemFrom   string `db:"item_from"  form:"from" json:"from"`
	ItemKey    string `db:"item_key"   form:"key" json:"key"`
	ItemNum    string `db:"item_num"   form:"num" json:"num"`
	ShopSn     string `db:"shop_sn"     form:"shopsn" json:"shopsn"`
	ActivityId uint64 `db:"activity_id" form:"activityid" json:"activityid"`
}
type ActivityData struct {
	Activity
	Item []*ActivityItem `form:"item" json:"item"`
	File *ActivityFile   `form:"file" json:"file"`
}

type ActivityLog struct {
	common.DbColumn
	LogId      uint64 `db:"log_id"    form:"log_id" json:"log_id"`
	ShopSn     string `db:"shop_sn"     form:"shopsn" json:"shopsn"`
	ActivityId uint64 `db:"activity_id" form:"activityid" json:"activityid"`
}
type ActivityClick struct {
	common.DbColumn
	ClickId    uint64 `db:"click_id" form:"id" json:"id"`
	ActivityId uint64 `db:"activity_id" form:"activityid" json:"activityid"`
	UserId     uint64 `db:"user_id" form:"userid" json:"userid"`
}
