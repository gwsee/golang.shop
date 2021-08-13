package service

import (
	"errors"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/system/model"
	"gwsee.com.api/utils"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func EditMenu(menu *model.SystemMenu, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE system_menu SET " +
		"menu_pid=? , menu_name=? , menu_icon=? , menu_target=? " +
		",menu_cache=? , menu_layout=? , menu_sort=? , menu_desc=? , url_id=? " +
		",edit_time=? ,edit_user=? ,state=? " +
		"where menu_id = ?"
	_, err = common.UpdateTable(sqlstr,
		menu.MenuPid, menu.MenuName, menu.MenuIcon, menu.MenuTarget,
		menu.MenuCache, menu.MenuLayout, menu.MenuSort, menu.MenuDesc, menu.UrlId,
		time.Now().Unix(), auth.User.UserId, menu.State,
		menu.MenuId)
	if err != nil {
		return
	}
	return
}
func AddMenu(menu *model.SystemMenu, auth *authU.GlobalConfig) (id int64, err error) {
	if menu.MenuId > 0 {
		sqlstr := "UPDATE system_menu SET " +
			"menu_pid=? , menu_name=? , menu_icon=? , menu_target=? " +
			",menu_cache=? , menu_layout=? , menu_sort=? , menu_desc=? , url_id=? " +
			",edit_time=? ,edit_user=? ,state=? " +
			"where menu_id = ?"
		//这里应该做判断 不能把自己家到自己下面--前端做了处理
		_, err = common.UpdateTable(sqlstr,
			menu.MenuPid, menu.MenuName, menu.MenuIcon, menu.MenuTarget,
			menu.MenuCache, menu.MenuLayout, menu.MenuSort, menu.MenuDesc, menu.UrlId,
			time.Now().Unix(), auth.User.UserId, menu.State,
			menu.MenuId)
	} else {
		sqlstr := "INSERT INTO system_menu (add_time,add_user,state," +
			"menu_pid,menu_name,menu_icon,menu_target" +
			",menu_cache,menu_layout,menu_sort,menu_desc,url_id) VALUES (?,?,?" +
			",?,?,?,?,?" +
			",?,?,?,?)"
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, menu.State,
			menu.MenuPid, menu.MenuName, menu.MenuIcon, menu.MenuTarget,
			menu.MenuCache, menu.MenuLayout, menu.MenuSort, menu.MenuDesc, menu.UrlId)
	}
	if err != nil {
		return
	}
	return
}
func ListMenu() {
}
func DelMenu(menuid string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE system_menu SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where is_del=0 and menu_id = ?"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		menuid)
	if err != nil {
		return
	}
	return
}
func SetMenu(menuid, state string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE system_menu SET " +
		"state=? " +
		", edit_time=? ,edit_user=? " +
		"where menu_id = ?"
	_, err = common.UpdateTable(sqlstr,
		state,
		time.Now().Unix(), auth.User.UserId,
		menuid)
	if err != nil {
		return
	}
	return
}
func FindMenu(menu *model.SystemMenu, menuid string) (err error) {
	sqlstr := "select * from  system_menu where menu_id = ?"
	err = common.FindTable(menu, sqlstr, menuid)
	if err != nil {
		return
	}
	return
}

// 提供一个数组 这里生成child类型的树形结构
func TreeMenu(menu *model.MenuUrl, data *common.Data) (err error) {
	//第一个是查询条件
	//第二个是需要组装的菜单有哪些
	var list []model.MenuUrl
	sqlL := "select menu.*,IFNULL(url.url_name, '')as url_name from system_menu menu left join system_url url on menu.url_id=url.url_id where menu.is_del = 0 order by menu.menu_sort desc"
	data.Total = 0 //有多少菜单是循环出来的
	err = common.ListTable(&list, sqlL)
	if err != nil {
		return
	}
	//1获取需要显示的菜单
	var menuids []uint64
	var menuTopids []uint64
	for _, v := range list {
		if menu.State != "" {
			if v.State == menu.State {
				menuids = append(menuids, v.MenuId)
			}
		} else if menu.MenuDesc != "" {
			if strings.Contains(v.MenuDesc, menu.MenuDesc) {
				menuids = append(menuids, v.MenuId)
			}
		} else if menu.MenuName != "" {
			if strings.Contains(v.MenuName, menu.MenuName) {
				menuids = append(menuids, v.MenuId)
			}
		} else if menu.UrlName != "" {
			if strings.Contains(v.UrlName, menu.UrlName) {
				menuids = append(menuids, v.MenuId)
			}
		} else {
			menuids = append(menuids, v.MenuId)
		}
	}
	//获取顶级ID
	//根据子项获取顶级数据
	for _, v := range menuids {
		res := getTopMenu(v, list)
		menuTopids = append(menuTopids, res.MenuId)
	}
	menuTopids = utils.UniqueSlice(menuTopids)
	//2计算有多少菜单 data.Total（一级菜单） 并且计算取的是那些进行展示（一级菜单）
	data.Total = len(menuTopids)
	if data.PageSize < 1 {
		data.PageSize = 15
	}
	data.PageTotal = int(math.Ceil(float64(data.Total) / float64(data.PageSize)))
	if data.PageNo < 1 || data.PageNo > data.PageTotal {
		data.PageNo = 1
	}
	menuTop := 0
	if data.PageNo*data.PageSize > data.Total {
		menuTop = data.Total
	} else {
		menuTop = data.PageNo * data.PageSize
	}
	showTop := menuTopids[(data.PageNo-1)*data.PageSize : menuTop]
	//获取需要展示的菜单id数组
	var menus []uint64
	for _, v := range menuids {
		var ids = getParentMenu(v, list)
		flag := false
		for _, i := range showTop {
			flag = utils.NumInSlice(i, ids)
			if flag {
				//切片添加切片
				menus = append(menus, ids...)
				break
			}
		}
	}
	//将菜单组装成需要的child类型并且去重
	menus = utils.UniqueSlice(menus)
	//然后组装成为想要的形式
	tree := buildTree(0, menus, list)
	data.List = tree
	//end
	return
}
func buildTree(pid uint64, ids []uint64, list []model.MenuUrl) (back []model.MenuTree) {
	for _, v := range list {
		if utils.NumInSlice(v.MenuId, ids) && v.MenuPid == pid {
			child := buildTree(v.MenuId, ids, list)
			tmp := model.MenuTree{
				MenuUrl:  v,
				Children: child,
			}
			back = append(back, tmp)
		}
	}
	return back
}

//返回的是切片数组  最后这个切片数组中如果包含
func getParentMenu(id uint64, list []model.MenuUrl) (ids []uint64) {
	for _, v := range list {
		if v.MenuId == id {
			ids = append(ids, v.MenuId)
			if v.MenuPid != 0 {
				ids = append(ids, getParentMenu(v.MenuPid, list)...)
			}
			break
		}
	}
	return
}
func getTopMenu(id uint64, list []model.MenuUrl) (back model.MenuUrl) {
	for _, v := range list {
		if v.MenuId == id {
			if v.MenuPid == 0 {
				back = v
			} else {
				back = getTopMenu(v.MenuPid, list)
			}
			break
		}
	}
	return
}
func UserMenu(id string) (menuid uint64, topIds []uint64, first model.MenuUrl, topMenus []model.MenuUrl, menu []model.MenuData, err error) {
	menuid, _ = strconv.ParseUint(id, 10, 64)
	//0 获取顶级菜单有哪些
	var list []model.MenuUrl
	sqlL := "select menu.*,IFNULL(url.url_name, '')as url_name,IFNULL(url.url_path, '')as url_path from system_menu menu left join system_url url on menu.url_id=url.url_id where menu.state = 1 and menu.is_del = 0 order by menu.menu_sort desc"
	err = common.ListTable(&list, sqlL)
	if err != nil {
		return
	}
	//1 查询这个人有哪些菜单
	//  --先假定是所有
	for _, v := range list {
		if reflect.DeepEqual(first, model.MenuUrl{}) && v.UrlPath != "" {
			first = v
		}
		if v.MenuPid == 0 {
			topIds = append(topIds, v.MenuId)
			topMenus = append(topMenus, v)
		}
	}
	//  --没有传的时候默认第一个
	if menuid == 0 {
		menuid = topIds[0]
	}
	if menuid == 0 {
		err = errors.New("暂无可访问应用")
		return
	}
	//2 (先查询所有菜单)在看这个人在id这个大类下面有哪些菜单
	menu = buildMenuTree(menuid, list)
	return
}
func buildMenuTree(pid uint64, list []model.MenuUrl) (tree []model.MenuData) {
	for _, v := range list {
		if pid == v.MenuPid {
			child := buildMenuTree(v.MenuId, list)
			component := ""
			if v.UrlPath != "" {
				component = "." + v.UrlPath
			}
			tmp := model.MenuData{
				MenuUrl:      v,
				UrlComponent: component,
				Children:     child,
			}
			tree = append(tree, tmp)
		}
	}
	return
}
