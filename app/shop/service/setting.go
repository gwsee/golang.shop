package service

import (
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/shop/model"
	model2 "gwsee.com.api/app/system/model"
	"strings"
	"time"
)

//新增
func AddSettingDefault(setting *model.ShopSettingDefault, auth *authU.GlobalConfig) (id int64, err error) {
	var info model.ShopSettingDefault
	err = FindSettingDefault(&info, setting.DefaultKey, auth.ShopSn)
	if err != nil {
		return 0, err
	}

	if info.DefaultId > 0 {
		sqlstr := "UPDATE shop_setting_default SET default_value=?, " +
			" edit_time=?, edit_user=?, state=? " +
			" where default_id=?"
		_, err = common.UpdateTable(sqlstr,
			setting.DefaultValue,
			time.Now().Unix(), auth.User.UserId, setting.State,
			info.DefaultId)
	} else {
		sqlstr := "INSERT INTO shop_setting_default (add_time,add_user,state,shop_sn," +
			"default_key,default_value) VALUES (?,?,?,?," +
			"?,?)"
		if setting.ShopSn == "" {
			setting.ShopSn = auth.ShopSn
		}
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, setting.State, setting.ShopSn,
			setting.DefaultKey, setting.DefaultValue)
	}
	return
}

//查询
func FindSettingDefault(setting *model.ShopSettingDefault, key, shopsn string) (err error) {
	sqlstr := "select * from shop_setting_default where default_key = ? and shop_sn = ? and is_del = 0"
	err = common.FindTable(setting, sqlstr, key, shopsn)
	return
}

//状态修改
func SetSettingDefault(keys, state string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop_setting_default SET " +
		"state=? " +
		", edit_time=? ,edit_user=? " +
		"where is_del=0 and shopsn= ? and default_key = ?"
	_, err = common.UpdateTable(sqlstr,
		state,
		time.Now().Unix(), auth.User.UserId,
		auth.ShopSn, keys)
	return
}

//列表
func ListSettingDefault(setting *model2.SystemSettingType, data *common.Data, auth *authU.GlobalConfig) (err error) {
	//先查数据 后查值
	var W []string
	W = append(W, "is_del = 0")
	if setting.TypeName != "" {
		W = append(W, "type_name like '%"+setting.TypeName+"%'")
	}
	if setting.Type != "" {
		W = append(W, "type like '%"+setting.Type+"%'")
	}
	if setting.State != "" {
		W = append(W, "state = '"+setting.State+"'")
	}
	sqlC := "select count(*) from system_setting_type"
	sqlL := "select * from system_setting_type"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.ShopSettingDefaultData
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
	var keys []string
	for _, v := range list {
		keys = append(keys, v.TypeName)
	}
	var w1 []string
	sqlD := "select default_key,default_value from shop_setting_default"
	w1 = append(w1, "is_del = 0")
	w1 = append(w1, "shop_sn = '"+auth.ShopSn+"'")
	w1 = append(w1, "default_key in ('"+strings.Join(keys, "','")+"')")
	sqlD = sqlD + " where " + strings.Join(w1, " and ")
	var temp []*model.ShopSettingDefault
	err = common.ListTable(&temp, sqlD)
	if err != nil {
		return
	}
	values := make(map[string]string)
	for _, v := range temp {
		values[v.DefaultKey] = v.DefaultValue
	}
	for _, v := range list {
		if values[v.TypeName] != "" {
			v.DefaultValue = values[v.TypeName]
		}
	}
	data.List = list
	//end
	return
}
