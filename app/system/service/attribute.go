package service

import (
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/system/model"
	"strconv"
	"strings"
	"time"
)

func AddAttribute(attribute *model.AttributeAndItem, auth *authU.GlobalConfig) (id int64, err error) {
	common.TransTable(0)
	defer func() {
		if err != nil {
			common.TransTable(2)
		} else {
			common.TransTable(1)
		}
	}()
	if attribute.AttributeId > 0 {
		sqlstr := "UPDATE system_attribute SET " +
			"attribute_name=?,attribute_desc=?,attribute_sort=?," +
			"edit_time=? ,edit_user=? ,state=? " +
			"where attribute_id = ?"
		_, err = common.UpdateTable(sqlstr,
			attribute.AttributeName, attribute.AttributeDesc, attribute.AttributeSort,
			time.Now().Unix(), auth.User.UserId, attribute.State,
			attribute.AttributeId)
	} else {
		sqlstr := "INSERT INTO system_attribute (add_time,add_user,state," +
			"attribute_name,attribute_desc,attribute_sort) VALUES (?,?,?" +
			",?,?,?)"
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, attribute.State,
			attribute.AttributeName, attribute.AttributeDesc, attribute.AttributeSort)
		attribute.AttributeId = uint64(id)
	}
	if err != nil {
		return
	}
	//对子项进行添加
	err = addAttributeItem(attribute, auth)
	if err != nil {
		return
	}
	return
}
func addAttributeItem(attribute *model.AttributeAndItem, auth *authU.GlobalConfig) (err error) {
	var list []*model.AttributeItem
	sqlL := "SELECT * FROM system_attribute_item where attribute_id=" + strconv.FormatUint(attribute.AttributeId, 10)
	err = common.ListTable(&list, sqlL)
	if err != nil {
		return
	}
	attr := make(map[string]uint64)
	if list != nil {
		for _, v := range list {
			attr[v.ItemName] = v.ItemId
		}
	}
	items := attribute.Items
	sqlUp := "UPDATE system_attribute_item SET is_del=?, edit_time=? ,edit_user=? where " +
		"attribute_id=? and item_name not in (?)"
	common.UpdateTable(sqlUp, time.Now().Unix(), time.Now().Unix(), auth.User.UserId,
		attribute.AttributeId, strings.Join(items, ","))
	for k, v := range items {
		id := attr[v]
		sqlstr := ""
		if id > 0 {
			sqlstr = "UPDATE system_attribute_item SET " +
				"is_del=0,item_name=?,item_sort=?," +
				" edit_time=? ,edit_user=? ,state=? " +
				"where item_id = ?"
			_, err = common.UpdateTable(sqlstr,
				v, k,
				time.Now().Unix(), auth.User.UserId, "1",
				id)

		} else {
			sqlstr = "INSERT INTO system_attribute_item (add_time,add_user,state," +
				"attribute_id,item_name,item_sort) VALUES (?,?,?" +
				",?,?,?)"
			_, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, "1",
				attribute.AttributeId, v, k)
		}
		if err != nil {
			return
		}
	}

	return
}
func ListAttribute(attribute *model.Attribute, data *common.Data) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	if attribute.AttributeName != "" {
		W = append(W, "attribute_name like '%"+attribute.AttributeName+"%'")
	}
	if attribute.AttributeDesc != "" {
		W = append(W, "attribute_desc like '%"+attribute.AttributeDesc+"%'")
	}
	if attribute.State != "" {
		W = append(W, "state = '"+attribute.State+"'")
	}
	sqlC := "select count(*) from system_attribute"
	sqlL := "select * from system_attribute"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.AttributeAndItem
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
	//对list处理 获取items
	getAllItems(list)
	data.List = list
	//end
	return
}

func getAllItems(list []*model.AttributeAndItem) {
	var ids []string
	for _, v := range list {
		ids = append(ids, strconv.FormatUint(v.AttributeId, 10))
	}
	var w []string
	w = append(w, "is_del = 0")
	w = append(w, "attribute_id in ("+strings.Join(ids, ",")+")")
	sqlL := "select * from system_attribute_item"
	if w != nil {
		sqlL = sqlL + " where " + strings.Join(w, " and ")
	}
	var itemList []*model.AttributeItem
	err := common.ListTable(&itemList, sqlL)
	if err != nil {
		return
	}
	for _, v := range list {
		var item []string
		for _, it := range itemList {
			if it.AttributeId == v.AttributeId {
				item = append(item, it.ItemName)
			}
		}
		v.Items = item
	}
	return
}

func DelAttribute(id string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE system_attribute SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where attribute_id = ? and is_del = 0"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		id)
	if err != nil {
		return
	}
	return
}
