package service

import (
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/system/model"
	"strings"
	"time"
)

func AddWebsite(web *model.SystemWebsite, auth *authU.GlobalConfig) (id int64, err error) {
	if web.WebId > 0 {
		sqlStr := "UPDATE system_website SET " +
			"web_name=? , web_host=? , web_favicon=? , web_license=?," +
			"web_logo=? , web_desc=? , web_pic=?," +
			"edit_time=? ,edit_user=? ,state=?," +
			"where web_id = ?"
		_, err = common.UpdateTable(sqlStr,
			web.WebName, web.WebHost, web.WebFavicon, web.WebLicense,
			web.WebLogo, web.WebDesc, web.WebPic,
			time.Now().Unix(), auth.User.UserId, web.State,
			web.WebId)
	} else {
		sqlStr := "INSERT INTO system_website (add_time,add_user,state," +
			"web_name,web_host,web_favicon,web_license," +
			"web_logo,web_desc,web_pic)" +
			" VALUES (?,?,?," +
			"?,?,?,?," +
			"?,?,?)"
		id, err = common.InsertTable(sqlStr, time.Now().Unix(), auth.User.UserId, web.State,
			web.WebName, web.WebHost, web.WebFavicon, web.WebLicense,
			web.WebLogo, web.WebDesc, web.WebPic)
	}
	return
}
func ListWebsite(web *model.SystemWebsite, data *common.Data) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	if web.WebName != "" {
		W = append(W, "web_name like '%"+web.WebName+"%'")
	}
	if web.WebHost != "" {
		W = append(W, "web_host like '%"+web.WebHost+"%'")
	}
	if web.WebLicense != "" {
		W = append(W, "web_license = '"+web.WebFavicon+"'")
	}
	if web.WebDesc != "" {
		W = append(W, "web_desc like '%"+web.WebDesc+"%'")
	}
	if web.State != "" {
		W = append(W, "state = '"+web.State+"'")
	}
	sqlC := "select count(*) from system_website"
	sqlL := "select * from system_website"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	//start
	var list []*model.SystemWebsite
	data.Total, err = common.CountTable(sqlC)
	if err != nil {
		return
	}
	limit := common.BuildCount(data)
	sqlL = sqlL + limit
	err = common.ListTable(&list, sqlL)
	if err != nil {
		return
	}
	data.List = list
	//end
	return
}
func DelWebsite(webid string, auth *authU.GlobalConfig) (err error) {
	sqlStr := "UPDATE system_website SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where web_id = ? and is_del = 0"
	_, err = common.UpdateTable(sqlStr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		webid)
	if err != nil {
		return
	}
	return
}
func FindWebsite(web *model.SystemWebsite, host string) (err error) {
	sqlStr := "select * from  system_website where is_del=0 and web_host = ?"
	err = common.FindTable(web, sqlStr, host)
	if err != nil {
		return
	}
	return
}
