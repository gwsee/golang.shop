package service

import (
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/order/model"
	"strconv"
	"strings"
)

//获取订单列表（以订单子表为主表）
func ListOrder(order *model.OrderData, data *common.Data, auth *authU.GlobalConfig) (err error) {
	var w []string
	w = append(w, "order.is_del = 0")
	w = append(w, "order.shop_sn = '"+auth.ShopSn+"'")
	if order.OrderNo != "" {
		w = append(w, "order.order_no like '%"+order.OrderNo+"%'")
	}
	if order.OrderGrade > 0 {
		w = append(w, "order.order_grade = '"+strconv.FormatUint(order.OrderGrade, 10)+"'")

	}
	if order.OrderFrom != "" {
		w = append(w, "order.order_from like '%"+order.OrderFrom+"%'")

	}
	if order.UserName != "" {
		w = append(w, "user.user_name like '%"+order.UserName+"%'")
	}

	if order.GoodsSn != "" {
		w = append(w, "detail.goods_sn like '%"+order.GoodsSn+"%'")
	}
	if order.GoodsTitle != "" {
		w = append(w, "detail.goods_title like '%"+order.GoodsTitle+"%'")

	}
	if order.DetailType != "" {
		w = append(w, "detail.detail_type like '%"+order.DetailType+"%'")

	}
	if order.DetailTransport > 0 {
		w = append(w, "detail.detail_transport = '"+strconv.FormatUint(order.DetailTransport, 10)+"'")

	}
	sqlC := "select count(*) from order_detail detail " +
		"inner join `order` `order` on detail.order_no=`order`.order_no " +
		"inner join user user on `order`.user_id = user.user_id "
	sqlL := "select detail.*," +
		"`order`.order_grade,`order`.order_remarks,`order`.order_pay,`order`.order_from," +
		"`order`.user_id,`order`.order_grade,`order`.order_grade,`order`.order_grade," +
		"`order`.order_grade,`order`.pay_time,`order`.pay_type," +
		"user.user_name " +
		"from order_detail detail " +
		"inner join `order` `order` on detail.order_no=`order`.order_no " +
		"inner join user user on `order`.user_id = user.user_id "
	if w != nil {
		sqlC = sqlC + " where " + strings.Join(w, " and ")
		sqlL = sqlL + " where " + strings.Join(w, " and ")
	}
	var list []*model.OrderData
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
	for _, v := range list {
		v.StateName = model.OrderState[v.State]
		v.TypeName = model.DetailType[v.DetailType]
	}
	data.List = list
	//end
	return

}

//获取订单服务列表（以订单子表为服务的为主表关联订单服务表）
func ListOrderService(service *model.OrderServiceData, data *common.Data, auth *authU.GlobalConfig) (err error) {
	var w []string
	w = append(w, "service.is_del = 0")
	w = append(w, "service.shop_sn = '"+auth.ShopSn+"'")
	if service.OrderNo != "" {
		w = append(w, "service.order_no like '%"+service.OrderNo+"%'")
	}
	if service.UserName != "" {
		w = append(w, "service.user_name like '%"+service.UserName+"%'")
	}

	if service.UserMobile != "" {
		w = append(w, "service.user_mobile like '%"+service.UserMobile+"%'")
	}
	if service.GoodsTitle != "" {
		w = append(w, "detail.goods_title like '%"+service.GoodsTitle+"%'")

	}
	if service.DetailType != "" {
		w = append(w, "detail.detail_type like '%"+service.DetailType+"%'")

	}
	if service.State != "" {
		w = append(w, "detail.detail_transport = '"+service.State+"'")
	}
	sqlC := "select count(*) from order_service service " +
		"inner join order_detail detail on detail.order_no=service.order_no "
	sqlL := "select service.*," +
		"detail.goods_title,detail.goods_price,detail.goods_num,detail.detail_pay," +
		"detail.detail_type,detail.state as detail_tate " +
		"from order_service service " +
		"inner join order_detail detail on detail.order_no=service.order_no "
	if w != nil {
		sqlC = sqlC + " where " + strings.Join(w, " and ")
		sqlL = sqlL + " where " + strings.Join(w, " and ")
	}
	var list []*model.OrderServiceData
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
	for _, v := range list {
		v.StateName = model.OrderState[v.DetailState]
		v.TypeName = model.DetailType[v.DetailType]
	}
	data.List = list
	//end
	return
}
