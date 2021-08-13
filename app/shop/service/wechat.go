package service

import (
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/shop/model"
	"strings"
	"time"
)

func AddWechat(wechat *model.ShopWechat, auth *authU.GlobalConfig) (id int64, err error) {
	if wechat.WechatId > 0 {
		sqlstr := "UPDATE shop_wechat SET wechat_appkey=?, wechat_appsecrect=?, wechat_account=?," +
			" edit_time=?, edit_user=?, state=? " +
			" where wechat_id=?"
		_, err = common.UpdateTable(sqlstr,
			wechat.WechatAppKey, wechat.WechatAppSecrect, wechat.WechatAccount,
			time.Now().Unix(), auth.User.UserId, wechat.State,
			wechat.WechatId)
	} else {
		sqlstr := "INSERT INTO shop_wechat (add_time,add_user,state,shop_sn," +
			"wechat_appkey,wechat_appsecrect,wechat_account) VALUES (?,?,?,?," +
			"?,?,?)"
		if wechat.ShopSn == "" {
			wechat.ShopSn = auth.ShopSn
		}
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, wechat.State, wechat.ShopSn,
			wechat.WechatAppKey, wechat.WechatAppSecrect, wechat.WechatAccount)
	}
	return
}
func FindWechat(we *model.ShopWechat, key string) (err error) {
	sqlstr := "select * from shop_wechat where wechat_appkey = ?"
	err = common.FindTable(we, sqlstr, key)
	return
}
func DelWechat(wechatId string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop_wechat SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where wechat_id = ?"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		wechatId)
	return
}

//列表
func ListWechat(wechat *model.ShopWechat, data *common.Data, auth *authU.GlobalConfig) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	W = append(W, "shop_sn = '"+auth.ShopSn+"'")
	if wechat.WechatAppKey != "" {
		W = append(W, "wechat_appkey like '%"+wechat.WechatAppKey+"%'")
	}
	if wechat.WechatAccount != "" {
		W = append(W, "wechat_account like '%"+wechat.WechatAccount+"%'")
	}
	if wechat.State != "" {
		W = append(W, "state = '"+wechat.State+"'")
	}
	sqlC := "select count(*) from shop_wechat"
	sqlL := "select * from shop_wechat"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.ShopWechat
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
	for _, v := range list {
		v.WechatAppSecrect = "*******"
	}
	data.List = list
	//end
	return
}
