package service

import (
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/system/model"
	"gwsee.com.api/utils"
	"math"
	"strconv"
	"strings"
	"time"
)

func AddClassify(classify *model.Classify, auth *authU.GlobalConfig) (id int64, err error) {
	if classify.ClassifyId > 0 {
		sqlstr := "UPDATE system_classify SET " +
			"classify_name=?,classify_desc=?,classify_sort=?,classify_end=?,classify_pid=?," +
			"edit_time=?,edit_user=?,state=? " +
			"where classify_id=?"
		_, err = common.UpdateTable(sqlstr,
			classify.ClassifyName, classify.ClassifyDesc, classify.ClassifySort, classify.ClassifyEnd, classify.ClassifyPid,
			time.Now().Unix(), auth.User.UserId, classify.State,
			classify.ClassifyId)
	} else {
		sqlstr := "INSERT INTO system_classify (add_time,add_user,state," +
			"classify_name,classify_desc,classify_sort,classify_end,classify_pid) VALUES (?,?,?," +
			"?,?,?,?,?)"
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, classify.State,
			classify.ClassifyName, classify.ClassifyDesc, classify.ClassifySort, classify.ClassifyEnd, classify.ClassifyPid)
	}
	return
}

// 商品的类目选择
func ListClassify(classify *model.Classify, data *common.Data) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	W = append(W, "classify_pid = '"+strconv.FormatUint(classify.ClassifyPid, 10)+"'")
	if classify.ClassifyName != "" {
		W = append(W, "classify_name like '%"+classify.ClassifyName+"%'")
	}
	if classify.ClassifyDesc != "" {
		W = append(W, "classify_desc like '%"+classify.ClassifyDesc+"%'")
	}
	if classify.State != "" {
		W = append(W, "state = '"+classify.State+"'")
	}
	sqlC := "select count(*) from system_classify"
	sqlL := "select * from system_classify"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.Classify
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

//树形的menu
func TreeClassify(classify *model.Classify, data *common.Data) (err error) {
	sqlL := "select * from system_classify where is_del=0" //这里查询的部门可能多个
	var list []model.Classify
	err = common.ListTable(&list, sqlL)
	if err != nil {
		return
	}
	var ids []uint64
	var topIds []uint64
	//获取需要显示的id切片
	for _, v := range list {
		if classify.State != "" {
			if v.State == classify.State {
				ids = append(ids, v.ClassifyId)
			}
		} else if classify.ClassifyName != "" {
			if strings.Contains(v.ClassifyName, classify.ClassifyName) {
				ids = append(ids, v.ClassifyId)
			}
		} else if classify.ClassifyDesc != "" {
			if strings.Contains(v.ClassifyDesc, classify.ClassifyDesc) {
				ids = append(ids, v.ClassifyId)
			}
		} else {
			ids = append(ids, v.ClassifyId)
		}
	}
	//获取需要显示的顶级ID切片
	for _, v := range ids {
		res := getTopItem(v, list)
		topIds = append(topIds, res.ClassifyId)
	}
	topIds = utils.UniqueSlice(topIds)

	data.Total = len(topIds)
	if data.PageSize < 1 {
		data.PageSize = 15
	}
	data.PageTotal = int(math.Ceil(float64(data.Total) / float64(data.PageSize)))
	if data.PageNo < 1 || data.PageNo > data.PageTotal {
		data.PageNo = 1
	}
	top := 0
	if data.PageNo*data.PageSize > data.Total {
		top = data.Total
	} else {
		top = data.PageNo * data.PageSize
	}
	//得出需要显示的顶层ID数组切片
	showTop := topIds[(data.PageNo-1)*data.PageSize : top]
	//需要显示的id 有哪些，通过各个符合条件的id（获取其所有上层id）
	var seeIds []uint64
	for _, v := range ids {
		var ids = getParentItem(v, list)
		flag := false
		for _, i := range showTop {
			flag = utils.NumInSlice(i, ids)
			if flag {
				//切片添加切片
				seeIds = append(seeIds, ids...)
				break
			}
		}
	}
	seeIds = utils.UniqueSlice(seeIds)
	//用得到的所有id 去获取他树形结构的数据
	tree := buildTreeClassify(0, seeIds, list)
	data.List = tree
	return
}
func buildTreeClassify(pid uint64, ids []uint64, list []model.Classify) (back []model.ClassifyTree) {
	for _, v := range list {
		if utils.NumInSlice(v.ClassifyId, ids) && v.ClassifyPid == pid {
			child := buildTreeClassify(v.ClassifyId, ids, list)
			tmp := model.ClassifyTree{
				Classify: v,
				Children: child,
			}
			back = append(back, tmp)
		}
	}
	return
}
func getParentItem(id uint64, list []model.Classify) (ids []uint64) {
	for _, v := range list {
		if v.ClassifyId == id {
			ids = append(ids, v.ClassifyId)
			if v.ClassifyPid != 0 {
				ids = append(ids, getParentItem(v.ClassifyPid, list)...)
			}
			break
		}
	}
	return
}
func getTopItem(id uint64, list []model.Classify) (back model.Classify) {
	for _, v := range list {
		if v.ClassifyId == id {
			if v.ClassifyPid == 0 {
				back = v
			} else {
				back = getTopItem(v.ClassifyPid, list)
			}
			break
		}
	}
	return back
}

func DelClassify(id string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE system_classify SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where classify_id = ? and is_del = 0"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		id)
	if err != nil {
		return
	}
	return
}

//根类目属性
func FindClassifyAttribute(classify *model.ClassifyAndAttribute, id, state string) (err error) {
	//获取类目
	sqlstr1 := "select * from  system_classify where classify_id = ?"
	err = common.FindTable(classify, sqlstr1, id)
	if err != nil {
		return
	}
	//获取类目有的属性
	var item []*model.ClassifyAttributeAndItem
	sqstr2 := "select * from system_classify_attribute where  is_del=0 and classify_id = " + id
	if state == "1" {
		sqstr2 = sqstr2 + " and state=1"
	}
	err = common.ListTable(&item, sqstr2)
	if err != nil {
		return
	}
	//获取查询出来属性的 属性id数组
	var child []string
	if item != nil {
		for _, v := range item {
			child = append(child, strconv.FormatUint(v.AttributeId, 10))
		}
	}
	//如果有属性 则把属性对应的属性值查出来 对应到他的子项上面
	if child != nil {
		var itemChild []*model.ClassifyAttributeItem
		sqstr3 := "select * from system_classify_attribute_item where is_del=0 and classify_id = " + id + " and attribute_id in (" + strings.Join(child, ",") + ")"
		if state == "1" {
			sqstr3 = sqstr3 + " and state=1"
		}
		err = common.ListTable(&itemChild, sqstr3)
		if err != nil {
			return
		}
		if itemChild != nil {
			for _, v := range item {
				var tmp []*model.ClassifyAttributeItem
				for _, vv := range itemChild {
					if v.AttributeId == vv.AttributeId {
						tmp = append(tmp, vv)
					}
				}
				v.Items = tmp
			}
		}
	}
	classify.Items = item
	return
}
func UpdateClassifyAttribute(attribute *model.ClassifyAndAttribute, auth *authU.GlobalConfig) (err error) {
	common.TransTable(0)
	defer func() {
		if err != nil {
			common.TransTable(2)
		} else {
			common.TransTable(1)
		}
	}()
	items := attribute.Items
	id := attribute.ClassifyId
	//获取属性名数组
	var name []string
	for _, v := range items {
		name = append(name, v.AttributeName)
	}
	//1: "system_classify_attribute" 根据类目ID和传过来的属性名称 找到 所有数据库中存在的属性
	var itemList1 []*model.ClassifyAttribute
	sqlstr1 := "select * from system_classify_attribute where classify_id=" + strconv.FormatUint(id, 10) + "" +
		" and attribute_name in ('" + strings.Join(name, "','") + "')"
	err = common.ListTable(&itemList1, sqlstr1)
	if err != nil {
		return
	}
	//1.1:获取所有属性名与属性attribute_id的切片
	itemKey := make(map[string]uint64)
	var ids []string
	for _, v := range itemList1 {
		itemKey[v.AttributeName] = v.AttributeId
		ids = append(ids, strconv.FormatUint(v.AttributeId, 10))
	}
	//2: "system_classify_attribute" 执行一次删除操作 将这个类目ID的所有属性全部都干成已删除状态  -- 不删除传过来的属性名称
	sqlstr2 := "UPDATE system_classify_attribute SET is_del=?, edit_time=? ,edit_user=? where " +
		"classify_id=? and attribute_id not in (?)"
	common.UpdateTable(sqlstr2, time.Now().Unix(), time.Now().Unix(), auth.User.UserId,
		attribute.ClassifyId, strings.Join(ids, ","))
	//3: "system_classify_attribute" 执行循环操作将属性进行添加或者修改（1.1中获取的attribute_id）操作
	for k, v := range items {
		attributeId := itemKey[v.AttributeName]
		if attributeId > 0 {
			sqlstr := "UPDATE system_classify_attribute SET " +
				"classify_id=?,attribute_name=?,attribute_sort=?,is_alias=?," +
				"is_color=?,is_enumeration=?,is_input=?,is_crux=?," +
				"is_sale=?,is_search=?,is_must=?,is_multiple=?," +
				"edit_time=? ,edit_user=? ,state=?,is_del=0 " +
				"where attribute_id = ?"
			_, err = common.UpdateTable(sqlstr,
				id, v.AttributeName, k, v.IsAlias,
				v.IsColor, v.IsEnumeration, v.IsInput, v.IsCrux,
				v.IsSale, v.IsSearch, v.IsMust, v.IsMultiple,
				time.Now().Unix(), auth.User.UserId, v.State,
				attributeId)
			v.AttributeId = attributeId
		} else {
			var i int64
			sqlstr := "INSERT INTO system_classify_attribute " +
				"(add_time,add_user,state," +
				"classify_id,attribute_name,attribute_sort,is_alias," +
				"is_color,is_enumeration,is_input,is_crux," +
				"is_sale,is_search,is_must,is_multiple) VALUES " +
				"(?,?,?," +
				"?,?,?,?," +
				"?,?,?,?," +
				"?,?,?,?)"
			i, err = common.InsertTable(sqlstr,
				time.Now().Unix(), auth.User.UserId, v.State,
				id, v.AttributeName, k, v.IsAlias,
				v.IsColor, v.IsEnumeration, v.IsInput, v.IsCrux,
				v.IsSale, v.IsSearch, v.IsMust, v.IsMultiple)
			v.AttributeId = uint64(i)
		}
		if err != nil {
			return
		}
		var child = v.Items
		var itemName []string

		//3.1: "system_classify_attribute_item" 在循环过程中 查看属性值数组，然后根据属性值数组 获取他可能存在的item_id
		for _, v1 := range child {
			itemName = append(itemName, v1.ItemName)
		}
		var itemList2 []*model.ClassifyAttributeItem
		sqlstr1 := "select * from system_classify_attribute_item where" +
			" attribute_id=" + strconv.FormatUint(v.AttributeId, 10) + "" +
			" and classify_id=" + strconv.FormatUint(id, 10) + "" +
			" and item_name in ('" + strings.Join(itemName, "','") + "')"
		err = common.ListTable(&itemList2, sqlstr1)
		if err != nil {
			return
		}
		//3.1.1: "system_classify_attribute_item" 获取所有值名与item_id的切片
		itemKey2 := make(map[string]uint64)
		var itemIds []string
		for _, v := range itemList2 {
			itemKey2[v.ItemName] = v.ItemId
			itemIds = append(itemIds, strconv.FormatUint(v.ItemId, 10))
		}
		//3.2: "system_classify_attribute_item" 获取之后 执行一次删除操作 删除根据 =attribute_id 且排除 not in ( item_id )
		sqlstr2 := "UPDATE system_classify_attribute_item SET is_del=?, edit_time=? ,edit_user=? where " +
			"classify_id=? and attribute_id=? and item_id not in (?)"
		common.UpdateTable(sqlstr2, time.Now().Unix(), time.Now().Unix(), auth.User.UserId,
			id, v.AttributeId, strings.Join(itemIds, ","))
		//3.3: "system_classify_attribute_item" 执行新增 或者 删除操作
		for k1, v1 := range child {
			itemId := itemKey2[v1.ItemName]
			if itemId > 0 {
				sqlstr := "UPDATE system_classify_attribute_item SET " +
					"classify_id=?,attribute_id=?,item_name=?,item_sort=?," +
					"edit_time=? ,edit_user=? ,state=?,is_del=0 " +
					"where item_id = ?"
				_, err = common.UpdateTable(sqlstr,
					id, v.AttributeId, v1.ItemName, k1,
					time.Now().Unix(), auth.User.UserId, "1",
					itemId)
			} else {
				sqlstr := "INSERT INTO system_classify_attribute_item (add_time,add_user,state," +
					"classify_id,attribute_id,item_name,item_sort) VALUES (?,?,?" +
					",?,?,?,?)"
				_, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, "1",
					id, v.AttributeId, v1.ItemName, k1)
			}
			if err != nil {
				return
			}
		}
	}
	if err != nil {
		return
	}
	return
}
func handleClassifyAttribute(attribute model.ClassifyAttribute, ids []string) (id int64, err error) {
	return
}
func handleClassifyAttributeItem() (err error) {
	return
}

//根类目品牌
func AddClassifyBrand(ids *model.ClassifyBrandIds, auth *authU.GlobalConfig) (id int64, err error) {
	common.TransTable(0)
	defer func() {
		if err != nil {
			common.TransTable(2)
		} else {
			common.TransTable(1)
		}
	}()
	//获取已经存在的数据
	sqlL := "select * from system_classify_brand where classify_id=" + strconv.FormatUint(ids.ClassifyId, 10)
	var list []*model.ClassifyBrand
	err = common.ListTable(&list, sqlL)
	if err != nil {
		return
	}
	keys := make(map[uint64]uint64)
	for _, v := range list {
		keys[v.BrandId] = v.Id
	}
	item := ids.Ids
	var itemStr []string
	for _, v := range item {
		itemStr = append(itemStr, v)
	}
	//将所有数据删除
	sqlDel := "UPDATE system_classify_brand SET is_del=?, edit_time=? ,edit_user=? where " +
		"classify_id=? and brand_id not in (?)"
	common.UpdateTable(sqlDel, time.Now().Unix(), time.Now().Unix(), auth.User.UserId,
		ids.ClassifyId, strings.Join(itemStr, ","))
	//查询最新的品牌信息
	var brands []*model.Brand
	sqlL = "select * from system_brand where brand_id in (" + strings.Join(itemStr, ",") + ")"
	err = common.ListTable(&brands, sqlL)
	if err != nil {
		return
	}
	brandKes := make(map[uint64]*model.Brand)
	for _, v := range brands {
		brandKes[v.BrandId] = v
	}
	//进行添加
	sqlAdd := "INSERT INTO system_classify_brand(add_time,add_user,state," +
		"classify_id,brand_id,brand_name,brand_logo) VALUES "
	var addQ []string
	var addV []interface{}
	for _, v := range item {
		i, _ := strconv.ParseUint(v, 10, 64)
		if keys[i] > 0 {
			sqlStr := "UPDATE system_classify_brand SET " +
				"is_del=0,brand_name=?,brand_logo=?," +
				"edit_time=?,edit_user=?,state=1 " +
				"where id=?"
			_, err = common.UpdateTable(sqlStr, brandKes[i].BrandName, brandKes[i].BrandLogo,
				time.Now().Unix(), auth.User.UserId,
				keys[i])
		} else {
			addQ = append(addQ, " (?,?,?,?,?,?,?)")
			addV = append(addV, time.Now().Unix())
			addV = append(addV, auth.User.UserId)
			addV = append(addV, 1)
			addV = append(addV, ids.ClassifyId)
			addV = append(addV, i)
			addV = append(addV, brandKes[i].BrandName)
			addV = append(addV, brandKes[i].BrandLogo)
		}
	}
	if addQ != nil {
		_, err = common.InsertTable(sqlAdd+strings.Join(addQ, ","), addV...)
	}
	return
}
func ListClassifyBrand(brand *model.ClassifyBrand, data *common.Data) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	if brand.BrandName != "" {
		W = append(W, "brand_name like '%"+brand.BrandName+"%'")
	}
	if brand.ClassifyId > 0 {
		W = append(W, "classify_id = '"+strconv.FormatUint(brand.ClassifyId, 10)+"'")
	}
	if brand.State != "" {
		W = append(W, "state = '"+brand.State+"'")
	}
	sqlC := "select count(*) from system_classify_brand"
	sqlL := "select * from system_classify_brand"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.ClassifyBrand
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
func DelClassifyBrand(id string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE system_classify_brand SET " +
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

//根类目品牌（子品）模板
func AddClassifyTemplate(template *model.ClassifyTemplateAndAttribute, auth *authU.GlobalConfig) (id int64, err error) {
	//1:对模板表进行新增或者编辑（有template_id) 注意：存的brand_id是system_classify_brand的id
	if template.TemplateId > 0 {
		sqlstr := "UPDATE system_classify_template SET " +
			"template_name=?,brand_id=?," +
			"edit_time=? ,edit_user=? ,state=?,is_del=0 " +
			"where template_id = ?"
		_, err = common.UpdateTable(sqlstr,
			template.TemplateName, template.BrandId,
			time.Now().Unix(), auth.User.UserId, template.State,
			template.TemplateId)
	} else {
		var i int64
		sqlstr := "INSERT INTO system_classify_template (add_time,add_user,state," +
			"template_name,brand_id) VALUES (?,?,?" +
			",?,?)"
		i, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, template.State,
			template.TemplateName, template.BrandId)
		template.TemplateId = uint64(i)
	}
	attrItems := template.Attribute
	var attrIds []string
	for _, v := range attrItems {
		attrIds = append(attrIds, strconv.FormatUint(v.AttributeId, 10))
	}
	//2:根据 template_id 查询已经存在的id 做成 切片 方便进行判断是不是已经存在
	sqlstr := "select * from system_classify_template_attribute where is_del=0 " +
		" and template_id=" + strconv.FormatUint(template.TemplateId, 10) +
		" and attribute_id in (" + strings.Join(attrIds, ",") + ")"
	var attrList []*model.ClassifyTemplateAttribute
	err = common.ListTable(&attrList, sqlstr)
	var attrKey map[uint64]uint64
	for _, v := range attrList {
		attrKey[v.AttributeId] = v.Id
	}
	//3:执行一次删除操作，排除已经存在的id
	sqlDel := "UPDATE system_classify_template_attribute SET is_del=?, edit_time=? ,edit_user=? where " +
		"template_id=? and attribute_id not in (?)"
	common.UpdateTable(sqlDel, time.Now().Unix(), time.Now().Unix(), auth.User.UserId,
		template.TemplateId, strings.Join(attrIds, ","))
	//4:循环 对属性进行新增或者编辑
	for _, v := range attrItems {
		i := attrKey[v.AttributeId]
		if i > 0 {
			sqlstr := "UPDATE system_classify_template_attribute SET " +
				"template_id=?,attribute_id=?,attribute_name=?,is_edit=?,attribute_value=?," +
				"edit_time=? ,edit_user=? ,state=?,is_del=0 " +
				"where id = ?"
			_, err = common.UpdateTable(sqlstr,
				v.TemplateId, v.AttributeId, v.AttributeName, v.IsEdit, v.AttributeValue,
				time.Now().Unix(), auth.User.UserId, template.State,
				template.TemplateId)
		} else {
			sqlstr := "INSERT INTO system_classify_template_attribute (add_time,add_user,state," +
				"template_id,attribute_id,attribute_name,is_edit,attribute_value) VALUES (?,?,?" +
				",?,?,?,?,?)"
			_, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, template.State,
				template.TemplateName, template.BrandId)
		}
	}

	return
}

//查询
func FindClassifyTemplate(info *model.ClassifyTemplateInfo, id string) (err error) {
	//1:根据templateid 获取到他的基本信息  system_classify_template
	var template model.ClassifyTemplate
	sqlstr1 := "select * from  system_classify_template where template_id = ?"
	err = common.FindTable(template, sqlstr1, id)
	//2:根据templateid 获取他的属性
	var attrValues []*model.ClassifyTemplateAttribute
	sqlstr2 := "select * from system_classify_template_attribute where  is_del=0 and template_id = " + id
	err = common.ListTable(&attrValues, sqlstr2)
	if err != nil {
		return
	}
	var keys []string //需要查询的属性 key和value的 id数组
	for _, v := range attrValues {
		keys = append(keys, strconv.FormatUint(v.AttributeId, 10))
	}
	var attrKeys []*model.ClassifyAttributeAndItem
	sqlstr3 := "select * from system_classify_attribute where  " +
		"is_del=0 and attribute_id in (" + strings.Join(keys, ",") + ")"
	err = common.ListTable(&attrKeys, sqlstr3)
	if err != nil {
		return
	}
	var attrItems []*model.ClassifyAttributeItem
	sqlstr4 := "select * from system_classify_attribute_item where  " +
		"is_del=0 and attribute_id in (" + strings.Join(keys, ",") + ")"
	err = common.ListTable(&attrItems, sqlstr4)
	if err != nil {
		return
	}
	for _, v := range attrKeys {
		var item []*model.ClassifyAttributeItem
		for _, vv := range attrItems {
			if v.AttributeId == vv.AttributeId {
				item = append(item, vv)
			}
		}
		v.Items = item
	}

	var attr []*model.ClassifyTemplateAttributeBase
	for _, v := range attrValues {
		var i model.ClassifyAttributeAndItem
		for _, vv := range attrKeys {
			if v.AttributeId == vv.AttributeId {
				i = *vv
			}
		}
		item := model.ClassifyTemplateAttributeBase{
			ClassifyTemplateAttribute: *v,
			ClassifyAttributeAndItem:  i,
		}
		attr = append(attr, &item)
	}
	//最后进行数据组装
	info = &model.ClassifyTemplateInfo{
		ClassifyTemplate: template,
		Attribute:        attr,
	}
	return
}
func ListClassifyTemplate(template *model.ClassifyTemplate, data *common.Data) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	if template.TemplateName != "" {
		W = append(W, "template_name like '%"+template.TemplateName+"%'")
	}
	if template.BrandId > 0 {
		W = append(W, "brand_id = '"+strconv.FormatUint(template.BrandId, 10)+"'")
	}
	if template.State != "" {
		W = append(W, "state = '"+template.State+"'")
	}
	sqlC := "select count(*) from system_classify_template"
	sqlL := "select * from system_classify_template"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.ClassifyTemplateAndAttribute
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
	//获取模板对应的属性值
	data.List = list
	return
}
func DelClassifyTemplate(id string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE system_classify_template SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where template_id = ? and is_del = 0"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		id)
	if err != nil {
		return
	}
	return
}
