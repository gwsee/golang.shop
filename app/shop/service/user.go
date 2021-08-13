package service

import (
	"errors"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/shop/model"
	ModelSystem "gwsee.com.api/app/system/model"
	ServiceSystem "gwsee.com.api/app/system/service"
	userMo "gwsee.com.api/app/user/model"
	userSe "gwsee.com.api/app/user/service"
	"strconv"
	"strings"
	"time"
)

//新增
func AddUser(user *model.ShopUser, auth *authU.GlobalConfig) (id int64, err error) {
	//查询这个角色在这个门店是否存在 不存在就新增 存在就编辑
	var obj model.ShopUser
	sqlstr := "select * from shop_user where is_del=0 and user_id= ? and shop_sn=?"
	// 添加新账号是在
	err = common.FindTable(&obj, sqlstr, user.UserId, user.ShopSn)
	strInt64 := strconv.FormatUint(obj.UserId, 10)
	id, _ = strconv.ParseInt(strInt64, 10, 64)
	if id > 0 {
		user.Id = obj.Id
		err = EditUser(user, auth)
	} else {
		sqlstr = "INSERT INTO shop_user (add_time,add_user,state," +
			"shop_sn,user_id,role_ids,depart_id) VALUES (?,?,?," +
			"?,?,?,?)"
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, user.State,
			user.ShopSn, user.UserId, user.RoleIds, user.DepartId)
	}
	return
}

//新增
func AddAccount(data *model.ShopUserDetail, auth *authU.GlobalConfig) (id int64, err error) {
	common.TransTable(0)
	defer func() {
		if err != nil {
			common.TransTable(2)
		} else {
			common.TransTable(1)
		}
	}()
	user := userMo.User{
		UserBase: data.UserBase,
		UserData: data.UserData,
	}
	id, err = userSe.AddUser(&user, auth, true)
	if err != nil {
		return
	}
	data.ShopUser.UserId = uint64(id)
	data.ShopUser.ShopSn = auth.ShopSn
	id, err = AddUser(&data.ShopUser, auth)
	return
}

//用户管理之门店信息 注册
func Register(shop *model.Shop, user *userMo.User, isAccounted, code string, auth *authU.GlobalConfig) (err error) {
	common.TransTable(0)
	defer func() {
		if err != nil {
			common.TransTable(2)
		} else {
			common.TransTable(1)
		}
	}()
	if isAccounted != "是" { //使用旧账号
		if user.UserName == "" || user.UserAccount == "" || user.UserPassword == "" || user.UserMobile == "" {
			err = errors.New("新账户数据不全")
			return
		}
	}
	var mess ModelSystem.SystemMessage
	mess.MessageMobile = user.UserMobile
	mess.MessageCode, _ = strconv.ParseUint(code, 0, 64)
	mess.MessageType = 1
	err = ServiceSystem.ValidMessage(&mess)
	if err != nil {
		return
	}
	//添加用户
	var id int64
	if isAccounted == "是" { //使用旧账号
		err = userSe.GetInfo(user, user.UserAccount)
		if err != nil {
			return
		}
		if user.UserId < 1 {
			err = errors.New("账户不存在")
			return
		}
		id = int64(user.UserId)

	} else { //新注册账户
		id, err = userSe.AddUser(user, auth, true)
		if err != nil {
			return
		}
	}

	//判断是否存在 和店铺的关系  如果已经存在 就不添加了 （先做1个人只能有一个店铺）
	sqlstr := "select * from shop_user where is_del=0 and user_id= ? "
	var obj model.ShopUser
	// 添加新账号是在
	err = common.FindTable(&obj, sqlstr, id)
	if err != nil {
		return
	}
	if obj.Id > 0 {
		err = errors.New("暂且一个账户只能绑定一个门店")
		return
	}
	//添加门店
	shop.ShopOwner = uint64(id)
	shop.AddUser = uint64(id)
	id, err = Add(shop, auth)
	if err != nil {
		return
	}
	//绑定门店和用户的关系
	var shopU model.ShopUser
	shopU.ShopSn = shop.ShopSn
	shopU.UserId = shop.ShopOwner
	_, err = AddUser(&shopU, auth)
	return
}

//查询
func FindUser(user *model.ShopUser, id string) (err error) {
	sqlstr := "select * from shop_user where id = ?"
	err = common.FindTable(user, sqlstr, id)
	return
}

//修改
func EditUser(user *model.ShopUser, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop_user SET shop_sn=?, user_id=?, role_ids=?, depart_id=?," +
		" edit_time=?, edit_user=?, state=? " +
		" where id=?"
	_, err = common.UpdateTable(sqlstr,
		user.ShopSn, user.UserId, user.RoleIds, user.DepartId,
		time.Now().Unix(), auth.User.UserId, user.State,
		user.Id)
	return
}

//删除
func DelUser(id string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop_user SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where id = ?"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		id)
	return
}

//状态修改
func SetUser(id, state string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop_user SET " +
		"state=? " +
		", edit_time=? ,edit_user=? " +
		"where id = ?"
	_, err = common.UpdateTable(sqlstr,
		state,
		time.Now().Unix(), auth.User.UserId,
		id)
	return
}
func CountUser() {

}

//列表
func ListUser(obj *model.ShopUserDetail, data *common.Data, auth *authU.GlobalConfig) (err error) {
	var W []string
	W = append(W, "shop_user.is_del = 0")
	//W = append(W, "user.is_del = 0")
	W = append(W, "shop_user.shop_sn = '"+auth.ShopSn+"'")
	if obj.DepartName != "" {
		W = append(W, "shop_depart.depart_name like '%"+obj.DepartName+"%'")
	}
	if obj.ShopName != "" {
		W = append(W, "shop.shop_name like '%"+obj.ShopName+"%'")
	}
	//角色 根据名称查询出来然后findinset
	if obj.UserName != "" {
		W = append(W, "user.user_name like '%"+obj.UserName+"%'")
	}
	if obj.UserMobile != "" {
		W = append(W, "user.user_mobile like '%"+obj.UserMobile+"%'")
	}
	if obj.UserEmail != "" {
		W = append(W, "user.user_email like '%"+obj.UserEmail+"%'")
	}
	if obj.UserWechat != "" {
		W = append(W, "user.user_wechat = '"+obj.UserWechat+"'")
	}
	if obj.State != "" {
		W = append(W, "shop_user.state = '"+obj.State+"'")
	}
	sqlC := "select count(*) from shop_user shop_user " +
		"inner join user user on user.user_id=shop_user.user_id " +
		"inner join shop shop on shop_user.shop_sn = shop.shop_sn " +
		"left join shop_depart shop_depart on shop_user.depart_id=shop_depart.depart_id "

	sqlL := "select shop_user.*,shop.shop_name,IFNULL(shop_depart.depart_name, '')as depart_name," +
		"user.user_name,user.user_account,user.user_avatar,user.user_mobile," +
		"user.user_address,user.user_email,user.user_sex,user.user_wechat " +
		"from shop_user shop_user " +
		"inner join user user on user.user_id=shop_user.user_id " +
		"inner join shop shop on shop_user.shop_sn = shop.shop_sn 	" +
		"left join shop_depart shop_depart on shop_user.depart_id=shop_depart.depart_id "
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	//start
	var list []*model.ShopUserDetail
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
