package service

import (
	"gwsee.com.api/app/activity/model"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"strings"
)

func ListLog(log *model.Log, data *common.Data, auth *authU.GlobalConfig) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	W = append(W, "shop_sn = '"+auth.ShopSn+"'")
	sqlC := "select count(*) from activity_log"
	sqlL := "select * from activity_log"
	if log.ActivityTitle != "" {
		W = append(W, "activity_title like '%"+log.ActivityTitle+"%'")
	}
	if log.UserName != "" {
		W = append(W, "user_name like '%"+log.UserName+"%'")
	}
	if log.OrderSn != "" {
		W = append(W, "order_sn like '%"+log.OrderSn+"%'")
	}
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.Log
	data.Total, err = common.CountTable(sqlC)
	if err != nil {
		return
	}
	limit := common.BuildCount(data)
	sqlL = sqlL + " order by add_time desc " + limit
	err = common.ListTable(&list, sqlL)
	if err != nil {
		return
	}
	data.List = list
	return
}
