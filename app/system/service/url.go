package service

import (
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/system/model"
	"strings"
	"time"
)

func EditUrl(url *model.SystemUrl) (err error) {

	if err != nil {
		return
	}
	return
}
func AddUrl(url *model.SystemUrl, auth *authU.GlobalConfig) (id int64, err error) {
	if url.UrlId > 0 {
		sqlstr := "UPDATE system_url SET " +
			"url_name=? , url_desc=? , url_type=? , url_path=? " +
			", edit_time=? ,edit_user=? ,state=? " +
			"where url_id = ?"
		_, err = common.UpdateTable(sqlstr,
			url.UrlName, url.UrlDesc, url.UrlType, url.UrlPath,
			time.Now().Unix(), auth.User.UserId, url.State,
			url.UrlId)
	} else {
		sqlstr := "INSERT INTO system_url (add_time,add_user,state," +
			"url_name,url_desc,url_type,url_path) VALUES (?,?,?" +
			",?,?,?,?)"
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, url.State,
			url.UrlName, url.UrlDesc, url.UrlType, url.UrlPath)
	}
	return
}
func ListUrl(url *model.SystemUrl, data *common.Data) (err error) {
	// sqlW := ""
	var W []string
	W = append(W, "is_del = 0")
	if url.UrlName != "" {
		W = append(W, "url_name like '%"+url.UrlName+"%'")
	}
	if url.UrlDesc != "" {
		W = append(W, "url_desc like '%"+url.UrlDesc+"%'")
	}
	if url.UrlPath != "" {
		W = append(W, "url_path like '%"+url.UrlPath+"%'")
	}
	if url.UrlType != "" {
		W = append(W, "url_type = '"+url.UrlType+"'")
	}
	if url.State != "" {
		W = append(W, "state = '"+url.State+"'")
	}
	sqlC := "select count(*) from system_url"
	sqlL := "select * from system_url"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	//start
	var list []*model.SystemUrl
	data.Total, err = common.CountTable(sqlC)
	if err != nil {
		return
	}
	limit := common.BuildCount(data)
	sqlL = sqlL + " order by add_time desc " + limit
	err = common.ListTable(&list, sqlL)
	if err != nil {
		return
	}
	data.List = list
	//end
	return
}
func DelUrl(urlid string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE system_url SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where url_id = ? and is_del = 0"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		urlid)
	if err != nil {
		return
	}
	return
}
func SetUrl(urlid, state string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE system_url SET " +
		"state=? " +
		", edit_time=? ,edit_user=? " +
		"where url_id = ?"
	_, err = common.UpdateTable(sqlstr,
		state,
		time.Now().Unix(), auth.User.UserId,
		urlid)
	if err != nil {
		return
	}
	return
}
func FindUrl(url *model.SystemUrl, urlid string) (err error) {
	sqlstr := "select * from  system_url where url_id = ?"
	err = common.FindTable(url, sqlstr, urlid)
	if err != nil {
		return
	}
	return
}
