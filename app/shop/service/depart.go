package service

import (
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/shop/model"
	"gwsee.com.api/utils"
	"math"
	"strings"
	"time"
)

//新增
func AddDepart(depart *model.ShopDepart, auth *authU.GlobalConfig) (id int64, err error) {
	if depart.DepartId > 0 {
		sqlstr := "UPDATE shop_depart SET depart_pid=?, depart_name=?,depart_sort=?, depart_desc=?,shop_sn=?, " +
			" edit_time=?, edit_user=?, state=? " +
			" where depart_id=?"
		_, err = common.UpdateTable(sqlstr,
			depart.DepartPid, depart.DepartName, depart.DepartSort, depart.DepartDesc, depart.ShopSn,
			time.Now().Unix(), auth.User.UserId, depart.State,
			depart.DepartId)
	} else {
		sqlstr := "INSERT INTO shop_depart (add_time,add_user,state," +
			"depart_pid,depart_name,depart_sort,depart_desc,shop_sn) VALUES (?,?,?," +
			"?,?,?,?,?)"
		if depart.ShopSn == "" {
			depart.ShopSn = auth.ShopSn
		}
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, depart.State,
			depart.DepartPid, depart.DepartName, depart.DepartSort, depart.DepartDesc, depart.ShopSn)
	}

	return
}

//查询
func FindDepart(depart *model.ShopDepart, shopid string) (err error) {
	sqlstr := "select * from shop_depart where shop_id = ?"
	err = common.FindTable(depart, sqlstr, shopid)
	return
}

//修改
func EditDepart(depart *model.ShopDepart, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop_depart SET depart_pid=?, depart_name=?, depart_desc=?, " +
		" edit_time=?, edit_user=?, state=? " +
		" where depart_id=?"
	_, err = common.UpdateTable(sqlstr,
		depart.DepartPid, depart.DepartName, depart.DepartDesc,
		time.Now().Unix(), auth.User.UserId, depart.State,
		depart.DepartId)
	return
}

//删除
func DelDepart(departid string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop_depart SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where depart_id = ?"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		departid)
	return
}

//状态修改
func SetDepart(departid, state string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE shop_depart SET " +
		"state=? " +
		", edit_time=? ,edit_user=? " +
		"where depart_id = ?"
	_, err = common.UpdateTable(sqlstr,
		state,
		time.Now().Unix(), auth.User.UserId,
		departid)
	return
}

//列表
func ListDepart(depart *model.ShopDepart, data *common.Data) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	if depart.DepartName != "" {
		W = append(W, "role_name like '%"+depart.DepartName+"%'")
	}
	if depart.DepartDesc != "" {
		W = append(W, "role_desc like '%"+depart.DepartDesc+"%'")
	}
	if depart.State != "" {
		W = append(W, "state = '"+depart.State+"'")
	}
	sqlC := "select count(*) from shop_depart"
	sqlL := "select * from shop_depart"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.ShopDepart
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

//树形的menu
func TreeDepart(depart *model.DepartShop, data *common.Data) (err error) {
	sqlL := "select depart.*,shop.shop_name from shop_depart depart inner join shop on depart.shop_sn=shop.shop_sn where depart.is_del=0 and depart.shop_sn in('" + depart.ShopSn + "')" //这里查询的部门可能多个
	var list []model.DepartShop
	err = common.ListTable(&list, sqlL)
	if err != nil {
		return
	}
	var ids []uint64
	var topIds []uint64
	//获取需要显示的id切片
	for _, v := range list {
		if depart.State != "" {
			if v.State == depart.State {
				ids = append(ids, v.DepartId)
			}
		} else if depart.DepartName != "" {
			if strings.Contains(v.DepartName, depart.DepartName) {
				ids = append(ids, v.DepartId)
			}
		} else if depart.DepartDesc != "" {
			if strings.Contains(v.DepartDesc, depart.DepartDesc) {
				ids = append(ids, v.DepartId)
			}
		} else {
			ids = append(ids, v.DepartId)
		}
	}
	//获取需要显示的顶级ID切片
	for _, v := range ids {
		res := getTopItem(v, list)
		topIds = append(topIds, res.DepartId)
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
	tree := buildTree(0, seeIds, list)
	data.List = tree
	return
}
func buildTree(pid uint64, ids []uint64, list []model.DepartShop) (back []model.DepartTree) {
	for _, v := range list {
		if utils.NumInSlice(v.DepartId, ids) && v.DepartPid == pid {
			child := buildTree(v.DepartId, ids, list)
			tmp := model.DepartTree{
				DepartShop: v,
				Children:   child,
			}
			back = append(back, tmp)
		}
	}
	return
}
func getParentItem(id uint64, list []model.DepartShop) (ids []uint64) {
	for _, v := range list {
		if v.DepartId == id {
			ids = append(ids, v.DepartId)
			if v.DepartPid != 0 {
				ids = append(ids, getParentItem(v.DepartPid, list)...)
			}
			break
		}
	}
	return
}
func getTopItem(id uint64, list []model.DepartShop) (back model.DepartShop) {
	for _, v := range list {
		if v.DepartId == id {
			if v.DepartPid == 0 {
				back = v
			} else {
				back = getTopItem(v.DepartPid, list)
			}
			break
		}
	}
	return back
}
