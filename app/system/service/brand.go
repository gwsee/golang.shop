package service

import (
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/system/model"
	"strings"
	"time"
)

func AddBrand(brand *model.Brand, auth *authU.GlobalConfig) (id int64, err error) {
	if brand.BrandId > 0 {
		sqlstr := "UPDATE system_brand SET " +
			"brand_name=?,brand_logo=?,brand_desc=?,brand_sort=?," +
			"brand_website=?,brand_story=?,group_id=?," +
			"edit_time=?,edit_user=?,state=? " +
			"where brand_id=?"
		_, err = common.UpdateTable(sqlstr,
			brand.BrandName, brand.BrandLogo, brand.BrandDesc, brand.BrandSort,
			brand.BrandWebsite, brand.BrandStory, brand.GroupId,
			time.Now().Unix(), auth.User.UserId, brand.State,
			brand.BrandId)

	} else {
		sqlstr := "INSERT INTO system_brand (add_time,add_user,state," +
			"brand_name,brand_logo,brand_desc,brand_sort" +
			",brand_website,brand_story,group_id)" +
			" VALUES (?,?,?," +
			"?,?,?,?," +
			"?,?,?)"
		id, err = common.InsertTable(sqlstr,
			time.Now().Unix(), auth.User.UserId, brand.State,
			brand.BrandName, brand.BrandLogo, brand.BrandDesc, brand.BrandSort,
			brand.BrandWebsite, brand.BrandStory, brand.GroupId)

	}
	return
}

func ListBrand(brand *model.BrandAndGroup, data *common.Data) (err error) {
	var W []string
	W = append(W, "system_brand.is_del = 0")
	if brand.BrandName != "" {
		W = append(W, "system_brand.brand_name like '%"+brand.BrandName+"%'")
	}
	if brand.BrandDesc != "" {
		W = append(W, "system_brand.brand_desc like '%"+brand.BrandDesc+"%'")
	}
	if brand.Name != "" {
		W = append(W, "system_brand_group.name like '%"+brand.Name+"%'")
	}
	if brand.State != "" {
		W = append(W, "system_brand.state = '"+brand.State+"'")
	}

	sqlC := "select count(*) from system_brand" +
		" inner join system_brand_group on system_brand.group_id =system_brand_group.id "
	sqlL := "select system_brand.*,system_brand_group.name from system_brand" +
		" inner join system_brand_group on system_brand.group_id =system_brand_group.id "
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.BrandAndGroup
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

func DelBrand(id string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE system_brand SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where brand_id = ? and is_del = 0"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		id)
	if err != nil {
		return
	}
	return
}

func AddBrandGroup(group *model.BrandGroup, auth *authU.GlobalConfig) (id int64, err error) {
	if group.Id > 0 {
		sqlstr := "UPDATE system_brand_group SET " +
			"`name`=?,`desc`=?,`sort`=?," +
			"`edit_time`=?,`edit_user`=?,`state`=? " +
			"where `id`=?"
		_, err = common.UpdateTable(sqlstr,
			group.Name, group.Desc, group.Sort,
			time.Now().Unix(), auth.User.UserId, group.State,
			group.Id)
	} else {
		sqlstr := "INSERT INTO system_brand_group (add_time,add_user,state," +
			"`name`,`desc`,`sort`) VALUES (?,?,?," +
			"?,?,?)"
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, group.State,
			group.Name, group.Desc, group.Sort)
	}
	return
}

func DelBrandGroup(id string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE system_brand_group SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where id = ? and is_del = 0"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		id)
	if err != nil {
		return
	}
	return
}

func ListBrandGroup(group *model.BrandGroup, data *common.Data) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	if group.Name != "" {
		W = append(W, "name like '%"+group.Name+"%'")
	}
	if group.Desc != "" {
		W = append(W, "desc like '%"+group.Desc+"%'")
	}
	if group.State != "" {
		W = append(W, "state = '"+group.State+"'")
	}
	sqlC := "select count(*) from system_brand_group"
	sqlL := "select * from system_brand_group"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.BrandGroup
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

func ListGroupBrand(group *model.BrandGroup, data *common.Data) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	if group.Name != "" {
		W = append(W, "name like '%"+group.Name+"%'")
	}
	if group.Desc != "" {
		W = append(W, "desc like '%"+group.Desc+"%'")
	}
	if group.State != "" {
		W = append(W, "state = '"+group.State+"'")
	}
	sqlL := "select * from system_brand_group"
	if W != nil {
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.BrandGroup
	err = common.ListTable(&list, sqlL)
	if err != nil {
		return
	}
	var items []*model.Brand
	sqlL = "select * from system_brand"
	err = common.ListTable(&items, sqlL)
	if err != nil {
		return
	}
	var itemList []*model.GroupItems
	for _, v := range list {
		item := model.GroupItems{
			BrandGroup: *v,
			Children:   nil,
		}
		for _, vv := range items {
			if v.Id == vv.GroupId {
				item.Children = append(item.Children, *vv)
			}
		}
		itemList = append(itemList, &item)
	}
	data.List = itemList
	return
}
