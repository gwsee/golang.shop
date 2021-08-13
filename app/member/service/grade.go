package service

import (
	"errors"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/member/model"
	"strings"
	"time"
)

func AddGrade(grade *model.Grade, auth *authU.GlobalConfig) (id int64, err error) {
	if grade.ShopSn == "" {
		grade.ShopSn = auth.ShopSn
	}
	id, err = uniqueGrade(grade)
	if err != nil {
		return
	}
	if id > 0 {
		err = errors.New("当前等级已经被注册")
		return
	}
	if grade.GradeId > 0 {
		sqlstr := "UPDATE grade SET grade_name=?,grade_desc=?,grade_time=?,grade_price=?, id=?, " +
			" edit_time=?, edit_user=?, state=? " +
			" where grade_id=?"
		_, err = common.UpdateTable(sqlstr,
			grade.GradeName, grade.GradeDesc, grade.GradeTime, grade.GradePrice, grade.Id,
			time.Now().Unix(), auth.User.UserId, grade.State,
			grade.GradeId)
	} else {
		sqlstr := "INSERT INTO grade (add_time,add_user,state," +
			"grade_name,grade_desc,grade_time,grade_price,shop_sn,id) VALUES (?,?,?," +
			"?,?,?,?,?,?)"
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, grade.State,
			grade.GradeName, grade.GradeDesc, grade.GradeTime, grade.GradePrice, grade.ShopSn, grade.Id)
	}

	return
}
func uniqueGrade(grade *model.Grade) (id int64, err error) {
	var info model.Grade
	sqlstr := "select * from grade where is_del =0 and shop_sn = ? and id = ? and grade_id <> ?"
	err = common.FindTable(&info, sqlstr, grade.ShopSn, grade.Id, grade.GradeId)
	if err != nil {
		return 0, err
	}
	id = int64(info.GradeId)
	return
}
func DelGrade(gradeid string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE grade SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where grade_id = ?"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		gradeid)
	return
}
func ListGrade(grade *model.Grade, data *common.Data, auth *authU.GlobalConfig) (err error) {
	var W []string
	W = append(W, "grade.is_del = 0")
	W = append(W, "grade.shop_sn = '"+auth.ShopSn+"'")
	if grade.GradeName != "" {
		W = append(W, "grade.grade_name like '%"+grade.GradeName+"%'")
	}
	if grade.GradeDesc != "" {
		W = append(W, "grade.grade_desc like '%"+grade.GradeDesc+"%'")
	}
	if grade.State != "" {
		W = append(W, "grade.state = '"+grade.State+"'")
	}
	sqlC := "select count(*) from grade inner join grade_template template on grade.id=template.id"
	sqlL := "select grade.*,template.step,template.logo,template.pic from grade inner join grade_template template on grade.id=template.id"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.GradeList
	data.Total, err = common.CountTable(sqlC)
	if err != nil {
		return
	}
	limit := common.BuildCount(data)
	order := " order by template.step asc "
	sqlL = sqlL + order + limit
	err = common.ListTable(&list, sqlL)
	if err != nil {
		return
	}
	data.List = list
	return
}
