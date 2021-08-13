package service

import (
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/shop/model"
	"strings"
	"time"
)

func AddClassify(classify *model.ShopClassify, auth *authU.GlobalConfig) (id int64, err error) {
	if classify.ID > 0 {
		sqlstr := "UPDATE shop_classify SET classify_id=?, brand_id=?,brand_name=?,classify_name=?," +
			" edit_time=?, edit_user=?, state=? " +
			" where id=?"
		_, err = common.UpdateTable(sqlstr,
			classify.ClassifyId, classify.BrandId, classify.BrandName, classify.ClassifyName,
			time.Now().Unix(), auth.User.UserId, classify.State,
			classify.ID)
	} else {
		sqlstr := "INSERT INTO shop_classify (add_time,add_user,state,shop_sn," +
			"classify_id,brand_id,classify_name,brand_name) VALUES (?,?,?,?," +
			"?,?,?,?)"
		if classify.ShopSn == "" {
			classify.ShopSn = auth.ShopSn
		}
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, classify.State, classify.ShopSn,
			classify.ClassifyId, classify.BrandId, classify.ClassifyName, classify.BrandName)
	}
	return
}
func ListClassify(classify *model.ShopClassify, data *common.Data) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	if classify.State != "" {
		W = append(W, "state = '"+classify.State+"'")
	}
	sqlC := "select count(*) from shop_classify"
	sqlL := "select * from shop_classify"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.ShopClassify
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
	return
}
func DelClassify(id string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop_classify SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where id = ?"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		id)
	return
}
