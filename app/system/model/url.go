package model

import "gwsee.com.api/app/common"

type SystemUrl struct {
	common.DbColumn
	UrlId   uint64 `db:"url_id"    form:"urlid" json:"urlid"`
	UrlName string `db:"url_name"  form:"name"  json:"name"`
	UrlDesc string `db:"url_desc"  form:"desc"  json:"desc"`
	UrlType string `db:"url_type"  form:"type"  json:"type"`
	UrlPath string `db:"url_path"  form:"path"  json:"path"`
}
