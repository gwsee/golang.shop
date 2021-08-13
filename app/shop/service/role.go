package service

import (
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/shop/model"
	"strings"
	"time"
)

//新增
func AddRole(role *model.ShopRole, auth *authU.GlobalConfig) (id int64, err error) {
	if role.RoleId > 0 {
		sqlstr := "UPDATE shop_role SET role_name=?, role_desc=?, menu_ids=?," +
			" edit_time=?, edit_user=?, state=? " +
			" where role_id=?"
		_, err = common.UpdateTable(sqlstr,
			role.RoleName, role.RoleDesc, role.MenuIds,
			time.Now().Unix(), auth.User.UserId, role.State,
			role.RoleId)
	} else {
		sqlstr := "INSERT INTO shop_role (add_time,add_user,state,shop_sn," +
			"role_name,role_desc,menu_ids) VALUES (?,?,?,?," +
			"?,?,?)"
		if role.ShopSn == "" {
			role.ShopSn = auth.ShopSn
		}
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, role.State, role.ShopSn,
			role.RoleName, role.RoleDesc, role.MenuIds)
	}
	return
}

//查询
func FindRole(role *model.ShopRole, roleid string) (err error) {
	sqlstr := "select * from shop_role where role_id = ?"
	err = common.FindTable(role, sqlstr, roleid)
	return
}

//修改
func EditRole(role *model.ShopRole, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop_role SET role_name=?, role_desc=?, menu_ids=?," +
		" edit_time=?, edit_user=?, state=? " +
		" where role_id=?"
	_, err = common.UpdateTable(sqlstr,
		role.RoleName, role.RoleDesc, role.MenuIds,
		time.Now().Unix(), auth.User.UserId, role.State,
		role.RoleId)
	return
}

//删除
func DelRole(roleid string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop_role SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where role_id = ?"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		roleid)
	return
}

//状态修改
func SetRole(roleid, state string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop_role SET " +
		"state=? " +
		", edit_time=? ,edit_user=? " +
		"where role_id = ?"
	_, err = common.UpdateTable(sqlstr,
		state,
		time.Now().Unix(), auth.User.UserId,
		roleid)
	return
}

//列表
func ListRole(role *model.ShopRole, data *common.Data, auth *authU.GlobalConfig) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	W = append(W, "shop_sn = '"+auth.ShopSn+"'")
	if role.RoleName != "" {
		W = append(W, "role_name like '%"+role.RoleName+"%'")
	}
	if role.RoleDesc != "" {
		W = append(W, "role_desc like '%"+role.RoleDesc+"%'")
	}
	if role.State != "" {
		W = append(W, "state = '"+role.State+"'")
	}
	sqlC := "select count(*) from shop_role"
	sqlL := "select * from shop_role"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.ShopRole
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
