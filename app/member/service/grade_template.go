package service

import (
	"errors"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/member/model"
	"strings"
	"time"
)

func AddGradeTemplate(template *model.GradeTemplate, auth *authU.GlobalConfig) (id int64, err error) {
	if template.ShopSn == "" {
		template.ShopSn = auth.ShopSn
	}
	id, err = uniqueTemplate(template)
	if err != nil {
		return
	}
	if id > 0 {
		err = errors.New("当前等级的会员模板已经存在")
		return
	}
	if template.Id > 0 {
		sqlstr := "UPDATE grade_template SET name=?,step=?,logo=?,pic=?, " +
			" edit_time=?, edit_user=?, state=? " +
			" where id=?"
		_, err = common.UpdateTable(sqlstr,
			template.Name, template.Step, template.Logo, template.Pic,
			time.Now().Unix(), auth.User.UserId, template.State,
			template.Id)
	} else {
		sqlstr := "INSERT INTO grade_template (add_time,add_user,state," +
			"name,step,logo,pic,shop_sn) VALUES (?,?,?," +
			"?,?,?,?,?)"
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, template.State,
			template.Name, template.Step, template.Logo, template.Pic, template.ShopSn)
	}

	return
}
func uniqueTemplate(template *model.GradeTemplate) (templateid int64, err error) {
	var info model.GradeTemplate
	sqlstr := "select * from grade_template where is_del =0 and shop_sn = ? and step = ? and id <> ?"
	err = common.FindTable(&info, sqlstr, template.ShopSn, template.Step, template.Id)
	if err != nil {
		return 0, err
	}
	templateid = int64(info.Id)
	return

}
func DelGradeTemplate(id string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE grade_template SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where id = ?"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		id)
	return
}
func ListGradeTemplate(template *model.GradeTemplate, data *common.Data, auth *authU.GlobalConfig) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	W = append(W, "shop_sn = '"+auth.ShopSn+"'")
	if template.Name != "" {
		W = append(W, "name like '%"+template.Name+"%'")
	}
	if template.State != "" {
		W = append(W, "state = '"+template.State+"'")
	}
	sqlC := "select count(*) from grade_template"
	sqlL := "select * from grade_template"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.GradeTemplate
	data.Total, err = common.CountTable(sqlC)
	if err != nil {
		return
	}
	limit := common.BuildCount(data)
	order := " order by step asc "
	sqlL = sqlL + order + limit
	err = common.ListTable(&list, sqlL)
	if err != nil {
		return
	}
	data.List = list
	return
}
