package service

import (
	"errors"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/goods/model"
	model2 "gwsee.com.api/app/system/model"
	"gwsee.com.api/app/system/service"
	"gwsee.com.api/utils"
	"strconv"
	"strings"
	"time"
)

func DelCommons(id string, auth *authU.GlobalConfig) (err error) {
	sqlStr := "UPDATE goods_commons SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where commons_id = ?"
	_, err = common.UpdateTable(sqlStr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		id)
	return
}
func AddCommons(commons *model.CommonsData, auth *authU.GlobalConfig) (id int64, err error) {
	common.TransTable(0)
	defer func() {
		if err != nil {
			common.TransTable(2)
		} else {
			common.TransTable(1)
		}
	}()
	// 没问题 1:商品主表的记录
	if commons.CommonsId > 0 {
		sqlStr := "UPDATE goods_commons SET " +
			" classify_id=?,classify_name=?,brand_id=?,commons_title=?,commons_sort=?," +
			" commons_cost=?,commons_price=?,commons_stock=?,commons_desc=?,commons_code=?," +
			" commons_type=?,commons_remark=?,commons_warn=?,commons_trans=?,commons_classify=?," +
			" commons_invented=?,commons_sn=?,commons_service=?,commons_sign=?," +
			" edit_time=?, edit_user=?, state=? " +
			" where commons_id=? and state=? and commons_order=?"
		_, err = common.UpdateTable(sqlStr,
			commons.ClassifyId, commons.ClassifyName, commons.BrandId, commons.CommonsTitle, commons.CommonsSort,
			commons.CommonsCost, commons.CommonsPrice, commons.CommonsStock, commons.CommonsDesc, commons.CommonsCode,
			commons.CommonsType, commons.CommonsRemark, commons.CommonsWarn, commons.CommonsTrans, commons.CommonsClassify,
			commons.CommonsInvented, commons.CommonsSn, commons.CommonsService, commons.CommonsSign,
			time.Now().Unix(), auth.User.UserId, commons.State,
			commons.CommonsId, 2, 0) //下架中商品 且是订单量为0的才能进行编辑
	} else {
		sqlStr := "INSERT INTO goods_commons (" +
			"classify_id,classify_name,brand_id,commons_title,commons_sort," +
			"commons_cost,commons_price,commons_stock,commons_desc,commons_code," +
			"commons_type,commons_remark,commons_warn,commons_trans,commons_classify," +
			"commons_invented,commons_sn,commons_service,commons_sign," +
			"add_time,add_user,state,shop_sn)" +
			"VALUES (" +
			"?,?,?,?,?," +
			"?,?,?,?,?," +
			"?,?,?,?,?," +
			"?,?,?,?," +
			"?,?,?,?)"
		if commons.ShopSn == "" {
			commons.ShopSn = auth.ShopSn
		}
		commons.CommonsSn = utils.BuildUnique("GC", commons.ShopSn)
		id, err = common.InsertTable(sqlStr,
			commons.ClassifyId, commons.ClassifyName, commons.BrandId, commons.CommonsTitle, commons.CommonsSort,
			commons.CommonsCost, commons.CommonsPrice, commons.CommonsStock, commons.CommonsDesc, commons.CommonsCode,
			commons.CommonsType, commons.CommonsRemark, commons.CommonsWarn, commons.CommonsTrans, commons.CommonsClassify,
			commons.CommonsInvented, commons.CommonsSn, commons.CommonsService, commons.CommonsSign,
			time.Now().Unix(), auth.User.UserId, commons.State, commons.ShopSn)
		commons.CommonsId = uint64(id)
	}
	if err != nil {
		return
	}
	// 尝试改成多线程  下面的
	// fmt.Println("第一步  商品的文件表记录")
	commons.File.ClassifyId = commons.ClassifyId
	commons.File.BrandId = commons.BrandId
	err = addFiles(commons.CommonsId, commons.File, auth) //没问题
	if err != nil {
		return
	}
	// fmt.Println("第二步  商品属性表的记录")
	err = addAttributes(commons.CommonsId, commons.Attribute, auth)
	if err != nil {
		return
	}
	// fmt.Println("第三步  商品子项表记录")
	err = addItems(commons, commons.Item, auth)
	if err != nil {
		return
	}
	// fmt.Println("结束")
	return
}

// 给商品添加属性  一对N关系 存在删除  (此处就不记录 属性具体值了 只记录当前商品的真实属性，历史属性需要在item表里去看了）
func addAttributes(id uint64, attribute []*model.CommonsAttribute, auth *authU.GlobalConfig) (err error) {
	var list []*model.CommonsAttribute
	sqlStr1 := "select * from goods_commons_attribute where is_del =0 and commons_id=" + strconv.FormatUint(id, 10)
	err = common.ListTable(&list, sqlStr1)
	if err != nil {
		return
	}
	listKey := make(map[uint64]uint64)
	for _, v := range list {
		listKey[v.AttributeId] = v.Id
	}
	// 先全部删除
	sqlDel := "update goods_commons_attribute set is_del=?, edit_time=? ,edit_user=? where commons_id=? and is_del=0"
	common.UpdateTable(sqlDel, time.Now().Unix(), time.Now().Unix(), auth.User.UserId, id)
	// 处理数据 如果i>0的数据 is_del赋值给0
	var addQ []string
	var addV []interface{}
	for _, v := range attribute {
		var i = listKey[v.AttributeId]
		if i > 0 {
			sqlStr := "UPDATE goods_commons_attribute SET " +
				" attribute_id=?,attribute_name=?,attribute_values=?," +
				" attribute_show=?,commons_id=?," +
				" edit_time=?,edit_user=?,is_del=?,state=? " +
				" where id=?"
			_, err = common.UpdateTable(sqlStr,
				v.AttributeId, v.AttributeName, v.AttributeValues,
				v.AttributeShow, id,
				time.Now().Unix(), auth.User.UserId, 0, 1,
				i) //下架中商品 且是订单量为0的才能进行编辑
		} else {
			if v.ShopSn == "" {
				v.ShopSn = auth.ShopSn
			}
			addQ = append(addQ, " (?,?,?,?,?,?,?,?,?)")
			addV = append(addV, v.AttributeId, v.AttributeName, v.AttributeValues, v.AttributeShow, id)
			addV = append(addV, time.Now().Unix(), auth.User.UserId, 1, auth.ShopSn)
		}
		if err != nil {
			return
		}

	}
	if addQ != nil {
		sqlAdd := "INSERT INTO goods_commons_attribute (" +
			"attribute_id,attribute_name,attribute_values," +
			"attribute_show,commons_id," +
			"add_time,add_user,state,shop_sn)" +
			"VALUES "
		_, err = common.InsertTable(sqlAdd+strings.Join(addQ, ","), addV...)
	}
	return
}

// 给商品添加文件  一对一关系 不存在删除
func addFiles(id uint64, file *model.CommonsFile, auth *authU.GlobalConfig) (err error) {
	var info model.CommonsFile
	sqlStr := "select * from goods_commons_file where is_del =0 and shop_sn = ? and commons_id = ?"
	err = common.FindTable(&info, sqlStr, auth.ShopSn, id)
	if info.FileId > 0 {
		sqlStr := "UPDATE goods_commons_file SET " +
			" commons_id=?,classify_id=?,brand_id=?,file_banner=?," +
			" file_video=?,file_logo=?,file_desc=?," +
			" edit_time=?,edit_user=?,state=? " +
			" where file_id=?"
		_, err = common.UpdateTable(sqlStr,
			id, file.ClassifyId, file.BrandId, file.FileBanner,
			file.FileVideo, file.FileLogo, file.FileDesc,
			time.Now().Unix(), auth.User.UserId, 1,
			info.FileId) //下架中商品 且是订单量为0的才能进行编辑
	} else {
		sqlStr := "INSERT INTO goods_commons_file (" +
			"commons_id,classify_id,brand_id,file_banner," +
			"file_video,file_logo,file_desc," +
			"add_time,add_user,state,shop_sn)" +
			"VALUES (" +
			"?,?,?,?," +
			"?,?,?," +
			"?,?,?,?)"
		if file.ShopSn == "" {
			file.ShopSn = auth.ShopSn
		}
		_, err = common.InsertTable(sqlStr,
			id, file.ClassifyId, file.BrandId, file.FileBanner,
			file.FileVideo, file.FileLogo, file.FileDesc,
			time.Now().Unix(), auth.User.UserId, 1, file.ShopSn)
	}
	return
}

// 给商品添加sku项  一对N关系 存在删除 各个具体子商品
func addItems(commons *model.CommonsData, item []*model.CommonsItem, auth *authU.GlobalConfig) (err error) {
	id := commons.CommonsId
	var str []string
	for _, v := range item {
		str = append(str, v.CommonsAttributes)
	}
	var items []*model.CommonsItem                  // 查询所有的数据
	itemsKey := make(map[string]*model.CommonsItem) // 这里以属性id为主来处理数据
	// 查询所有这个商品的数据
	sqlList := "select * from goods_commons_item where commons_id=" + strconv.FormatUint(id, 10) + " and is_del=0"
	err = common.ListTable(&items, sqlList)
	if err != nil {
		return
	}
	for _, v := range items {
		itemsKey[v.CommonsAttributes] = v
	}
	// 先删除未删除的数据
	sqlDel := "UPDATE goods_commons_item SET is_del=?,edit_time=?,edit_user=? where is_del=0 and commons_id=?"
	common.UpdateTable(sqlDel, time.Now().Unix(), time.Now().Unix(), auth.User.UserId, id)
	// 属性相同的情况下，
	//              1 如果未产生销量 订单的情况下，就直接修改他的数据，
	//              2 如果已经产生了订单数 销售数的情况下 可以改库存（往大于当前销量的改）(如果是改小的话 就直接做成第三种情况）
	//              3 如果已经产生了订单数 销售数的情况下 要改价格之类的 则是新增一个新的数据 （原相同属性数据则做假删除is_del）
	lenItem := len(items)
	var addQ []string
	var addV []interface{}
	for _, v := range item {
		v.CommonsId = id
		info := itemsKey[v.CommonsAttributes]
		editFlag := false // 默认不是编辑 是新增
		if info != nil && info.ItemId > 0 {
			if info.ItemSale == 0 && info.ItemOrder == 0 {
				// 1
				editFlag = true
			} else {
				// 价格一致允许修改 且库存数大于 原库存数的 2 3
				if info.ItemPrice == v.ItemPrice &&
					info.ItemMarket == v.ItemMarket &&
					info.ItemCost == v.ItemCost &&
					v.ItemStock >= info.ItemStock {
					editFlag = true
				}
			}
		}
		// 在itemid大于0的时候就给拿回来
		if editFlag {
			sqlStr := "UPDATE goods_commons_item SET " +
				" item_stock=?,item_cost=?,item_market=?,item_price=?," +
				" item_sn=?,item_code=?,commons_id=?,commons_attributes=?," +
				" edit_time=?, edit_user=?, state=?,is_del=? " +
				" where item_id=?"
			_, err = common.UpdateTable(sqlStr,
				v.ItemStock, v.ItemCost, v.ItemMarket, v.ItemPrice,
				info.ItemSn, info.ItemCode, v.CommonsId, v.CommonsAttributes,
				time.Now().Unix(), auth.User.UserId, 1, 0,
				info.ItemId) //下架中商品 且是订单量为0的才能进行编辑
		} else {
			lenItem = lenItem + 1
			if v.ShopSn == "" {
				v.ShopSn = auth.ShopSn
			}

			addQ = append(addQ, " (?,?,?,?,?,?,?,?,?,?,?,?)")
			addV = append(addV, v.ItemStock, v.ItemCost, v.ItemMarket, v.ItemPrice)
			addV = append(addV, commons.CommonsSn+"_"+strconv.Itoa(lenItem), commons.CommonsCode+"_"+strconv.Itoa(lenItem), v.CommonsId, v.CommonsAttributes)
			addV = append(addV, time.Now().Unix(), auth.User.UserId, 1, auth.ShopSn)

		}
		if err != nil {
			break
		}
	}
	if addQ != nil {
		sqlAdd := "INSERT INTO goods_commons_item (" +
			"item_stock,item_cost,item_market,item_price," +
			"item_sn,item_code,commons_id,commons_attributes," +
			"add_time,add_user,state,shop_sn)" +
			"VALUES "
		_, err = common.InsertTable(sqlAdd+strings.Join(addQ, ","), addV...)
	}
	return
}
func ListCommons(commons *model.Commons, data *common.Data, auth *authU.GlobalConfig) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	W = append(W, "shop_sn = '"+auth.ShopSn+"'")
	if commons.CommonsTitle != "" {
		W = append(W, "commons_title like '%"+commons.CommonsTitle+"%'")
	}
	if commons.CommonsRemark != "" {
		W = append(W, "commons_remark like '%"+commons.CommonsRemark+"%'")
	}
	if commons.State != "" {
		W = append(W, "state = '"+commons.State+"'")
	}
	sqlC := "select count(*) from goods_commons"
	sqlL := "select * from goods_commons"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.Commons
	data.Total, err = common.CountTable(sqlC)
	if err != nil {
		return
	}
	limit := common.BuildCount(data)
	sqlL = sqlL + " order by commons_sort desc " + limit
	err = common.ListTable(&list, sqlL)
	if err != nil {
		return
	}
	data.List = list
	return
}

func FindCommons(info *model.CommonsData, list *[]*model2.Classify, form *[]*model.CommonsAttributeForm, row *string, id string) (err error) {
	sqlStr := "select * from goods_commons where commons_id = ?"
	err = common.FindTable(info, sqlStr, id)
	if err != nil {
		return
	}
	if info == nil {
		err = errors.New("未找到此商品")
		return
	}
	// 多项表格的处理 table （即属性值与价格关系）
	var item []*model.CommonsItem
	err = findCommonsItems(&item, id)
	if err != nil {
		return
	}
	info.Item = item
	//  文件处理  -- 完成
	var file model.CommonsFile
	err = findCommonsFiles(&file, id)
	if err != nil {
		return
	}
	info.File = &file
	// 属性处理与属性值的处理
	var attribute []*model.CommonsAttribute
	err = findCommonsAttributes(&attribute, form, id, info.ClassifyId)
	if err != nil {
		return
	}
	info.Attribute = attribute
	FindAllClassifys(info.ClassifyId, list)
	// 处理销售属性的 id和名称 在table的第一行需要显示
	attrKey := make(map[uint64]string)
	for _, v := range attribute {
		attrKey[v.AttributeId] = v.AttributeName
	}
	var rowArr []string
	rowOneStr := item[0].CommonsAttributes
	rowOneArr := strings.Split(rowOneStr, "||")
	for _, v := range rowOneArr {
		vArr := strings.Split(v, ":")
		rowKeyStr := vArr[0]
		intNum, _ := strconv.Atoi(rowKeyStr)
		int64Num := uint64(intNum)
		rowName := attrKey[int64Num]
		rowArr = append(rowArr, rowKeyStr+":"+rowName)
	}

	*row = strings.Join(rowArr, "||")
	return
}
func findCommonsFiles(info *model.CommonsFile, id string) (err error) {
	sqlStr := "select * from goods_commons_file where commons_id = ? and is_del=0"
	err = common.FindTable(info, sqlStr, id)
	return
}

// 对商品的属性与其对应的值进行组装
func findCommonsAttributes(list *[]*model.CommonsAttribute, form *[]*model.CommonsAttributeForm, id string, classifyId uint64) (err error) {
	// 第一步根据类目ID查询他的 属性和属性值有哪些
	var info model2.ClassifyAndAttribute
	// 必须查询状态为1的数据 如果不是1的话 就可能出错
	err = service.FindClassifyAttribute(&info, strconv.FormatUint(classifyId, 10), "0")
	if err != nil {
		return
	}
	// 第二步根据商品id 查询出来他的属性值有哪些
	sqlStr := "select * from goods_commons_attribute where commons_id = " + id + " and is_del=0"
	err = common.ListTable(list, sqlStr)
	if err != nil {
		return
	}
	//      把它们根据属性id  把它的数值给划分下
	tempList := make(map[uint64]*model.CommonsAttribute)
	for _, v := range *list {
		tempList[v.AttributeId] = v
	}
	// 第三步 把查询出来的属性值给赋值到 CommonsAttributeForm 里面去
	for _, v := range info.Items {
		//取出属性组值 组成map对象 方便赋值
		itemList := v.Items
		itemKey := make(map[string]model2.ClassifyAttributeItem)
		for _, vv := range itemList {
			itemKey[vv.ItemName] = *vv
		}
		res := tempList[v.AttributeId]
		if res == nil {
			continue
		}
		var name = ""
		var nameList []string
		nameList = strings.Split(res.AttributeValues, "|")
		name = nameList[0]
		//if strings.Index(res.AttributeValues,"|")>-1{
		//	nameList = strings.Split(res.AttributeValues,"|")
		//	name=nameList[0]
		//}else{
		//	name = res.AttributeValues
		//}
		// 存在多选的情况下  他才是一个数组 如果不是多选的情况下 就是单个的
		var val interface{}
		if v.IsMultiple == 1 {
			var valList []model2.ClassifyAttributeItem
			for _, vvv := range nameList {
				if itemKey[vvv].ItemId != 0 {
					valList = append(valList, itemKey[vvv])
				} else {
					valList = append(valList, model2.ClassifyAttributeItem{
						ItemId:   0,
						ItemName: vvv,
					})
				}
			}
			val = valList

		} else {
			if itemKey[name].ItemId != 0 {
				val = itemKey[name]
			} else {
				val = model2.ClassifyAttributeItem{
					ItemId:   0,
					ItemName: name,
				}
			}
		}
		item := model.CommonsAttributeForm{
			*v,
			val,
		}
		*form = append(*form, &item)
	}
	// 然后查询这些属性的 各种信息
	return
}
func findCommonsItems(list *[]*model.CommonsItem, id string) (err error) { // goods_commons_item
	sqlStr := "select * from goods_commons_item where commons_id =" + id + " and is_del=0"
	err = common.ListTable(list, sqlStr)
	return
}

func SetCommons(id, state string, auth *authU.GlobalConfig) (err error) {
	sqlStr := "UPDATE goods_commons SET " +
		"state=? " +
		", edit_time=? ,edit_user=? " +
		"where commons_id = ?"
	_, err = common.UpdateTable(sqlStr,
		state,
		time.Now().Unix(), auth.User.UserId,
		id)
	return
}
