package service

import (
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/shop/model"
	"strings"
	"time"
)

//新增
func AddServiceType(obj *model.ShopServiceType, auth *authU.GlobalConfig) (id int64, err error) {
	if obj.TypeId > 0 {
		sqlstr := "UPDATE shop_service_type SET type_name=?, type_desc=?, type_minutes=?," +
			" edit_time=?, edit_user=?, state=? " +
			" where type_id=?"
		_, err = common.UpdateTable(sqlstr,
			obj.TypeName, obj.TypeDesc, obj.TypeMinutes,
			time.Now().Unix(), auth.User.UserId, obj.State,
			obj.TypeId)
	} else {
		sqlstr := "INSERT INTO shop_service_type (add_time,add_user,state,shop_sn," +
			"type_name,type_desc,type_minutes) VALUES (?,?,?,?," +
			"?,?,?)"
		if obj.ShopSn == "" {
			obj.ShopSn = auth.ShopSn
		}
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, obj.State, obj.ShopSn,
			obj.TypeName, obj.TypeDesc, obj.TypeMinutes)
	}
	return
}

//查询
func FindServiceType(obj *model.ShopServiceType, typeid string) (err error) {
	sqlstr := "select * from shop_service_type where type_id = ?"
	err = common.FindTable(obj, sqlstr, typeid)
	return
}

//删除
func DelServiceType(typeid string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop_service_type SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where type_id = ?"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		typeid)
	return
}

//状态修改
func SetServiceType(typeid, state string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop_service_type SET " +
		"state=? " +
		", edit_time=? ,edit_user=? " +
		"where type_id = ?"
	_, err = common.UpdateTable(sqlstr,
		state,
		time.Now().Unix(), auth.User.UserId,
		typeid)
	return
}

//列表
func ListServiceType(obj *model.ShopServiceType, data *common.Data, auth *authU.GlobalConfig) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	W = append(W, "shop_sn = '"+auth.ShopSn+"'")
	if obj.TypeName != "" {
		W = append(W, "type_name like '%"+obj.TypeName+"%'")
	}
	if obj.TypeDesc != "" {
		W = append(W, "type_desc like '%"+obj.TypeDesc+"%'")
	}
	if obj.State != "" {
		W = append(W, "state = '"+obj.State+"'")
	}
	sqlC := "select count(*) from shop_service_type"
	sqlL := "select * from shop_service_type"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.ShopServiceType
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
