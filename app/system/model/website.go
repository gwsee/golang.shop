package model

import "gwsee.com.api/app/common"

type SystemWebsite struct {
	common.DbColumn
	WebId      uint64 `db:"web_id"    form:"id" json:"id"`
	WebName    string `db:"web_name"  form:"name"  json:"name"`
	WebHost    string `db:"web_host"  form:"host"  json:"host"`
	WebFavicon string `db:"web_favicon"  form:"favicon"  json:"favicon"`
	WebLicense string `db:"web_license"  form:"license"  json:"license"`
	WebLogo    string `db:"web_logo"  form:"logo"  json:"logo"`
	WebDesc    string `db:"web_desc"  form:"desc"  json:"desc"`
	WebPic     string `db:"web_pic"  form:"pic"  json:"pic"`
}
