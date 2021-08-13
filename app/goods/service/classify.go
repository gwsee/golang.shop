package service

import (
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/goods/model"
	modelSys "gwsee.com.api/app/system/model"
	"sort"
	"strings"
	"time"
)

//删除
func DelClassify(classifyid string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE goods_classify SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where classify_id = ?"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		classifyid)
	return
}
func AddClassify(classify *model.Classify, auth *authU.GlobalConfig) (id int64, err error) {
	if classify.ClassifyId > 0 {
		sqlstr := "UPDATE goods_classify SET classify_name=?,classify_alias=?,classify_pic=?,classify_sort=?," +
			" edit_time=?, edit_user=?, state=? " +
			" where classify_id=?"
		_, err = common.UpdateTable(sqlstr,
			classify.ClassifyName, classify.ClassifyAlias, classify.ClassifyPic, classify.ClassifySort,
			time.Now().Unix(), auth.User.UserId, classify.State,
			classify.ClassifyId)
	} else {
		sqlstr := "INSERT INTO goods_classify (add_time,add_user,state," +
			"classify_name,classify_alias,classify_pic,classify_sort,shop_sn) VALUES (?,?,?," +
			"?,?,?,?,?)"
		if classify.ShopSn == "" {
			classify.ShopSn = auth.ShopSn
		}
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, classify.State,
			classify.ClassifyName, classify.ClassifyAlias, classify.ClassifyPic, classify.ClassifySort, classify.ShopSn)
	}

	return
}

func ListClassify(classify *model.Classify, data *common.Data, auth *authU.GlobalConfig) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	W = append(W, "shop_sn = '"+auth.ShopSn+"'")
	if classify.ClassifyName != "" {
		W = append(W, "classify_name like '%"+classify.ClassifyName+"%'")
	}
	if classify.State != "" {
		W = append(W, "state = '"+classify.State+"'")
	}
	sqlC := "select count(*) from goods_classify"
	sqlL := "select * from goods_classify"
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
	sqlL = sqlL + " order by classify_sort desc " + limit
	err = common.ListTable(&list, sqlL)
	if err != nil {
		return
	}
	data.List = list
	return
}

// --- 对其做排序 待优化 待看懂  -----https://blog.csdn.net/qq_17308321/article/details/94998236
type ByClassifyPid []*modelSys.Classify

func (a ByClassifyPid) Len() int           { return len(a) }
func (a ByClassifyPid) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByClassifyPid) Less(i, j int) bool { return a[i].ClassifyPid < a[j].ClassifyPid }

func FindAllClassifys(id uint64, list *[]*modelSys.Classify) (err error) {
	idNow := id
	for idNow > 0 {
		var info modelSys.Classify
		sqlstr := "select * from system_classify where classify_id=?"
		err = common.FindTable(&info, sqlstr, idNow)
		if err != nil {
			return
		}
		idNow = info.ClassifyPid
		*list = append(*list, &info)
	}
	sort.Sort(ByClassifyPid(*list))
	return
}
