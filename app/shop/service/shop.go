package service

import (
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/shop/model"
	"gwsee.com.api/utils"
	"strings"
	"time"
)

//新增  门店编辑的时候 门店编号 门店id 与门店所有着不能被编辑
func Add(shop *model.Shop, auth *authU.GlobalConfig) (id int64, err error) {
	if shop.ShopId > 0 {
		sqlstr := "UPDATE shop" +
			" SET shop_name=?, shop_logo=?, shop_pid=?, shop_addr=?," +
			" shop_breif=?, shop_fullname=?, shop_tel=?, " +
			" edit_time=?, edit_user=?, state=? " +
			" where shop_id=?"
		_, err = common.UpdateTable(sqlstr,
			shop.ShopName, shop.ShopLogo, shop.ShopPid, shop.ShopAddr,
			shop.ShopBreif, shop.ShopBreif, shop.ShopTel,
			time.Now().Unix(), auth.User.UserId, shop.State,
			shop.ShopId)
	} else {
		sqlstr := "INSERT INTO shop (add_time,add_user,state," +
			"shop_name,shop_logo,shop_sn,shop_pid,shop_addr," +
			"shop_breif,shop_fullname,shop_tel,shop_owner) VALUES (?,?,?," +
			"?,?,?,?,?," +
			"?,?,?,?)"
		if shop.ShopOwner == 0 {
			shop.ShopOwner = auth.User.UserId
		}
		if shop.ShopSn == "" {
			shop.ShopSn = utils.Krand(6, 3)
		}
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, 1,
			shop.ShopName, shop.ShopLogo, shop.ShopSn, shop.ShopPid, shop.ShopAddr,
			shop.ShopBreif, shop.ShopBreif, shop.ShopTel, shop.ShopOwner)
	}

	return
}

//查询
func Find(shop *model.Shop, shopsn string) (err error) {
	sqlstr := "select * from shop where shop_sn = ?"
	err = common.FindTable(shop, sqlstr, shopsn)
	return
}

//修改
func Edit(shop *model.Shop, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop SET shop_name=?, shop_logo=?, shop_pid=?, shop_addr=?," +
		" shop_breif=?, shop_fullname=?, shop_tel=?, shop_owner=?," +
		" edit_time=?, edit_user=?, state=? " +
		" where shop_id=?"
	_, err = common.UpdateTable(sqlstr,
		shop.ShopName, shop.ShopLogo, shop.ShopPid, shop.ShopAddr,
		shop.ShopBreif, shop.ShopBreif, shop.ShopTel, shop.ShopOwner,
		time.Now().Unix(), auth.User.UserId, shop.State,
		shop.ShopId)
	return
}

//删除
func Del(shopid string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where shop_id = ?"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		shopid)
	return
}

//状态修改
func Set(shopid, state string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop SET " +
		"state=? " +
		", edit_time=? ,edit_user=? " +
		"where shop_id = ?"
	_, err = common.UpdateTable(sqlstr,
		state,
		time.Now().Unix(), auth.User.UserId,
		shopid)
	return
}
func List(shop *model.Shop, data *common.Data, auth *authU.GlobalConfig) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	W = append(W, "shop_sn = '"+auth.ShopSn+"'")
	if shop.ShopName != "" {
		W = append(W, "shop_name like '%"+shop.ShopName+"%'")
	}
	if shop.ShopFullname != "" {
		W = append(W, "shop_fullname like '%"+shop.ShopFullname+"%'")
	}
	if shop.ShopBreif != "" {
		W = append(W, "shop_breif like '%"+shop.ShopBreif+"%'")
	}
	if shop.ShopAddr != "" {
		W = append(W, "shop_addr like '%"+shop.ShopAddr+"%'")
	}
	if shop.State != "" {
		W = append(W, "state = '"+shop.State+"'")
	}
	sqlC := "select count(*) from shop"
	sqlL := "select * from shop"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.Shop
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

//列表 -- 获取这个人在这个网站上面有哪些 门店信息
func ListShopUser(list interface{}, where string) (err error) {
	if where != "" {
		where = " where " + where
	}
	sqlstr := "select shop.* from shop INNER JOIN shop_user on shop.shop_sn = shop_user.shop_sn " + where
	err = common.ListTable(list, sqlstr)
	return
}
