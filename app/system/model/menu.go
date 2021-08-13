package model

import "gwsee.com.api/app/common"

type SystemMenu struct {
	common.DbColumn
	MenuId     uint64 `db:"menu_id" form:"menuid" json:"menuid"`
	MenuPid    uint64 `db:"menu_pid"  form:"pid" json:"pid"`
	MenuName   string `db:"menu_name"  form:"name" json:"name"`
	MenuIcon   string `db:"menu_icon"  form:"icon" json:"icon"`
	MenuTarget string `db:"menu_target"   form:"target" json:"target"`
	MenuCache  uint64 `db:"menu_cache"   form:"cache" json:"cache"`
	MenuLayout string `db:"menu_layout"  form:"layout" json:"layout"`
	MenuSort   uint64 `db:"menu_sort"   form:"sort" json:"sort"`
	MenuDesc   string `db:"menu_desc"  form:"desc" json:"desc"`
	UrlId      uint64 `db:"url_id" form:"urlid" json:"urlid"`
}

type MenuUrl struct {
	SystemMenu
	UrlName string `db:"url_name" form:"urlname" json:"urlname"`
	UrlPath string `db:"url_path"  form:"path"  json:"path"`
}
type MenuTree struct {
	MenuUrl
	Children []MenuTree `form:"children" json:"children" ` //不大写就外部不收数据
}

type MenuData struct {
	MenuUrl
	UrlComponent string     `json:"component"`
	UrlRedirect  string     `json:"redirect"`
	Children     []MenuData `json:"routes"`
}
