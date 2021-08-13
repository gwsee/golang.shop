package service

import (
	"gwsee.com.api/app/activity/model"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"strings"
	"time"
)

func AddActivity(activity *model.ActivityData, auth *authU.GlobalConfig) (id int64, err error) {
	common.TransTable(0)
	defer func() {
		if err != nil {
			common.TransTable(2)
		} else {
			common.TransTable(1)
		}
	}()
	if activity.ActivityId > 0 {
		sqlStr := "UPDATE activity SET" +
			"activity_title=?,activity_type=?,activity_desc=?,activity_begin=?," +
			"activity_end=?,activity_detail=?,activity_recommend=?,activity_rules=?," +
			"activity_sort=?,activity_stock=?,activity_invented=?,activity_remark=?," +
			"activity_limit=?,activity_price=?,activity_num=?," +
			"edit_time=?, edit_user=?, state=? " +
			"where activity_id=? and state=? and commons_order=?"
		_, err = common.UpdateTable(sqlStr,
			activity.ActivityTitle, activity.ActivityType, activity.ActivityDesc, activity.ActivityBegin,
			activity.ActivityEnd, activity.ActivityDetail, activity.ActivityRecommend, activity.ActivityRules,
			activity.ActivitySort, activity.ActivityStock, activity.ActivityInvented, activity.ActivityRemark,
			activity.ActivityLimit, activity.ActivityPrice, activity.ActivityNum,
			time.Now().Unix(), auth.User.UserId, 0,
			activity.ActivityId, 0, 0)
	} else {
		sqlStr := "INSERT INTO activity (" +
			"activity_title,activity_type,activity_desc,activity_begin," +
			"activity_end,activity_detail,activity_recommend,activity_rules," +
			"activity_sort,activity_stock,activity_invented,activity_remark," +
			"activity_limit,activity_price,activity_num" +
			"add_time,add_user,state,shop_sn) VALUES(" +
			"?,?,?,?," +
			"?,?,?,?," +
			"?,?,?,?," +
			"?,?,?," +
			"?,?,?,?)"
		if activity.ShopSn == "" {
			activity.ShopSn = auth.ShopSn
		}
		id, err = common.InsertTable(sqlStr,
			activity.ActivityTitle, activity.ActivityType, activity.ActivityDesc, activity.ActivityBegin,
			activity.ActivityEnd, activity.ActivityDetail, activity.ActivityRecommend, activity.ActivityRules,
			activity.ActivitySort, activity.ActivityStock, activity.ActivityInvented, activity.ActivityRemark,
			activity.ActivityLimit, activity.ActivityPrice, activity.ActivityNum,
			time.Now().Unix(), auth.User.UserId, activity.State, activity.ShopSn)
		activity.ActivityId = uint64(id)
	}
	if err != nil {
		return
	}
	activity.File.ActivityId = activity.ActivityId
	err = addFiles(activity.ActivityId, activity.File, auth)
	if err != nil {
		return
	}
	err = addItems(activity.ActivityId, activity.Item, auth)
	return
}
func addItems(id uint64, item []*model.ActivityItem, auth *authU.GlobalConfig) (err error) {
	//要先全部删除
	delTime := time.Now().Unix()
	sqlDel := "UPDATE activity_item SET is_del=?,edit_time=?,edit_user=? where is_del=0 and activity_id=?"
	common.UpdateTable(sqlDel, delTime, delTime, auth.User.UserId, id)

	//然后查询
	for _, v := range item {
		var info model.ActivityItem
		sql := "select * from activity_item where is_del=?,activity_id=?,item_from=?,item_key=?"
		err = common.FindTable(&info, sql, delTime, id, v.ItemFrom, v.ItemKey)
		if err != nil {
			return
		}
		if v.ShopSn == "" {
			v.ShopSn = auth.ShopSn
		}
		if info.ItemId > 0 {
			sql = "UPDATE activity_item SET " +
				"item_num=?,edit_time=?,edit_user=?,state=?,is_del=? " +
				"where item_id=?"
			_, err = common.UpdateTable(sql,
				v.ItemNum, delTime, auth.User.UserId, 1, 0,
				info.ItemId)
		} else {
			sql = "INSERT INTO activity_item (" +
				"activity_id,item_from,item_key,item_num," +
				"add_time,add_user,state,shop_sn)"
			_, err = common.InsertTable(sql,
				id, v.ItemFrom, v.ItemKey, v.ItemNum,
				delTime, auth.User.UserId, 1, v.ShopSn)
		}

		if err != nil {
			return
		}
	}
	return
}
func addFiles(id uint64, file *model.ActivityFile, auth *authU.GlobalConfig) (err error) {
	var info model.ActivityFile
	sqlStr := "select * from activity_file where is_del =0 and shop_sn = ? and commons_id = ?"
	err = common.FindTable(&info, sqlStr, auth.ShopSn, id)
	if info.FileId > 0 {
		sqlStr := "UPDATE activity_file SET " +
			" activity_id=?,file_banner=?," +
			" file_video=?,file_logo=?,file_desc=?," +
			" edit_time=?,edit_user=?,state=? " +
			" where file_id=?"
		_, err = common.UpdateTable(sqlStr,
			id, file.FileBanner,
			file.FileVideo, file.FileLogo, file.FileDesc,
			time.Now().Unix(), auth.User.UserId, 1,
			info.FileId) //下架中商品 且是订单量为0的才能进行编辑
	} else {
		sqlStr := "INSERT INTO activity_file (" +
			"activity_id,file_banner," +
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
			id, file.FileBanner,
			file.FileVideo, file.FileLogo, file.FileDesc,
			time.Now().Unix(), auth.User.UserId, 1, file.ShopSn)
	}
	return
}

func ListActivity(activity *model.Activity, data *common.Data, auth *authU.GlobalConfig) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	W = append(W, "shop_sn = '"+auth.ShopSn+"'")
	sqlC := "select count(*) from activity"
	sqlL := "select * from activity"
	if activity.ActivityTitle != "" {
		W = append(W, "activity_title like '%"+activity.ActivityTitle+"%'")
	}
	if activity.ActivityType != "" {
		W = append(W, "activity_type = '"+activity.ActivityType+"'")
	}
	if activity.ActivityDesc != "" {
		W = append(W, "activity_desc like '%"+activity.ActivityDesc+"%'")
	}
	if activity.ActivityRemark != "" {
		W = append(W, "activity_remark like '%"+activity.ActivityRemark+"%'")
	}
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.Activity
	data.Total, err = common.CountTable(sqlC)
	if err != nil {
		return
	}
	limit := common.BuildCount(data)
	sqlL = sqlL + " order by activity_sort desc " + limit
	err = common.ListTable(&list, sqlL)
	if err != nil {
		return
	}
	data.List = list
	return
	return
}

func FindActivity(id string, info *model.ActivityData) (err error) {
	var activity model.Activity
	var file model.ActivityFile
	var item []*model.ActivityItem
	sqlStr := "select * from activity where activity_id = ?"
	err = common.FindTable(&activity, sqlStr, id)
	if err != nil {
		return
	}
	info.Activity = activity
	err = findActivityFiles(&file, id)
	if err != nil {
		return
	}
	info.File = &file
	err = findActivityItems(&item, id)
	if err != nil {
		return
	}
	info.Item = item
	return
}
func findActivityItems(list *[]*model.ActivityItem, id string) (err error) {
	sqlStr := "select * from activity_item where activity_id =" + id + " and is_del=0"
	err = common.ListTable(list, sqlStr)
	return
}
func findActivityFiles(info *model.ActivityFile, id string) (err error) {
	sqlStr := "select * from activity_file where activity_id = ? and is_del=0"
	err = common.FindTable(info, sqlStr, id)
	return
}
func SetActivity(id, state string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE activity SET " +
		"state=? " +
		", edit_time=? ,edit_user=? " +
		"where activity_id = ?"
	_, err = common.UpdateTable(sqlstr,
		state,
		time.Now().Unix(), auth.User.UserId,
		id)
	return
}

func DelActivity(id string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE activity SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where activity_id = ?"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		id)
	return
}
