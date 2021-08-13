package service

import (
	"errors"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/goods/model"
	model2 "gwsee.com.api/app/system/model"
	"strconv"
	"strings"
)

func ListCommonsM(commons *model.Commons, data *common.Data, auth *authU.GlobalConfig) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	if auth.ShopSn != "" {
		W = append(W, "shop_sn = '"+auth.ShopSn+"'")
	}
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

func FindCommonsM(info *model.CommonsData, list *[]*model2.Classify, form *[]*model.CommonsAttributeForm, row *string, id string) (err error) {
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
