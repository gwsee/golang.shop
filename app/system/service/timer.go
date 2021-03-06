package service

import (
	"errors"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/system/model"
	"strconv"
	"strings"
	"time"
)

func AddTimerClassify(classify *model.TimerClassify, auth *authU.GlobalConfig) (id int64, err error) {
	if classify.ClassifyId > 0 {
		sqlstr := "UPDATE system_timer_classify SET " +
			"classify_name=?,classify_desc=?,classify_sign=?," +
			"edit_time=?,edit_user=?,state=? " +
			"where classify_id=?"
		_, err = common.UpdateTable(sqlstr,
			classify.ClassifyName, classify.ClassifyDesc, classify.ClassifySign,
			time.Now().Unix(), auth.User.UserId, classify.State,
			classify.ClassifyId)
	} else {
		sqlstr := "INSERT INTO system_timer_classify (add_time,add_user,state," +
			"classify_name,classify_desc,classify_sign) VALUES (?,?,?," +
			"?,?,?)"
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, classify.State,
			classify.ClassifyName, classify.ClassifyDesc, classify.ClassifySign)
	}
	return
}
func ListTimerClassify(classify *model.TimerClassify, data *common.Data) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	if classify.ClassifyName != "" {
		W = append(W, "classify_name like '%"+classify.ClassifyName+"%'")
	}
	if classify.ClassifyDesc != "" {
		W = append(W, "classify_desc like '%"+classify.ClassifyDesc+"%'")
	}
	if classify.ClassifySign != "" {
		W = append(W, "classify_desc like '%"+classify.ClassifySign+"%'")
	}
	if classify.State != "" {
		W = append(W, "state = '"+classify.State+"'")
	}
	if classify.ClassifyId > 0 {
		W = append(W, "classify_id = '"+strconv.FormatUint(classify.ClassifyId, 10)+"'")
	}
	sqlC := "select count(*) from system_timer_classify"
	sqlL := "select * from system_timer_classify"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.TimerClassify
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
	return
}
func DelTimerClassify(id string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE system_timer_classify SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where classify_id = ? and is_del = 0"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		id)
	if err != nil {
		return
	}
	return
}

func AddTimerLog(log *model.TimerLog, auth *authU.GlobalConfig) (id int64, err error) {
	if log.LogId > 0 {
		sqlstr := "UPDATE system_timer_log SET " +
			"timer_id=?,log_res=?,log_command=?," +
			"edit_time=?,edit_user=?,state=? " +
			"where log_id=?"
		_, err = common.UpdateTable(sqlstr,
			log.TimerId, log.LogRes, log.LogCommand,
			time.Now().Unix(), auth.User.UserId, log.State,
			log.LogId)
	} else {
		sqlstr := "INSERT INTO system_timer_log (add_time,add_user,state," +
			"timer_id,log_res,log_command) VALUES (?,?,?," +
			"?,?,?)"
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, log.State,
			log.TimerId, log.LogRes, log.LogCommand)
	}
	return
}
func ListTimerLog(obj *model.TimerLog, data *common.Data) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	// W = append(W, "shop_user.shop_sn = '"+auth.ShopSn+"'")
	if obj.TimerName != "" {
		W = append(W, "timer_name like '%"+obj.TimerName+"%'")
	}
	//?????? ??????????????????????????????findinset
	if obj.TimerSign != "" {
		W = append(W, "timer_sign like '%"+obj.TimerSign+"%'")
	}

	if obj.ClassifyName != "" {
		W = append(W, "classify_name like '%"+obj.ClassifyName+"%'")
	}
	if obj.ClassifySign != "" {
		W = append(W, "classify_sign = '"+obj.ClassifySign+"'")
	}
	sqlC := "select count(*) from system_timer_log"
	sqlL := "select * from system_timer_log"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	//start
	var list []*model.TimerLog
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
func DelTimerLog(id string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE system_timer_log SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where log_id = ? and is_del = 0"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		id)
	if err != nil {
		return
	}
	return
}

func AddTimer(timer *model.Timer, auth *authU.GlobalConfig) (id int64, err error) {
	if timer.TimerId > 0 {
		sqlstr := "UPDATE system_timer SET " +
			"timer_name=?,timer_desc=?,timer_sign=?,timer_command=?,timer_exec=?,classify_id=?," +
			"edit_time=?,edit_user=?,state=? " +
			"where timer_id=?"
		_, err = common.UpdateTable(sqlstr,
			timer.TimerName, timer.TimerDesc, timer.TimerSign, timer.TimerCommand, timer.TimerExec, timer.ClassifyId,
			time.Now().Unix(), auth.User.UserId, timer.State,
			timer.TimerId)
	} else {
		sqlstr := "INSERT INTO system_timer (add_time,add_user,state," +
			"timer_name,timer_desc,timer_sign,timer_command,timer_exec,classify_id) VALUES (?,?,?," +
			"?,?,?,?,?,?)"
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, timer.State,
			timer.TimerName, timer.TimerDesc, timer.TimerSign, timer.TimerCommand, timer.TimerExec, timer.ClassifyId)
	}
	// ??????
	var classify model.TimerClassify
	sqlstr1 := "select * from  system_timer_classify where classify_id = ?"
	common.FindTable(&classify, sqlstr1, timer.ClassifyId)
	if classify.ClassifyId > 0 {
		err = handleTimer(0, classify.ClassifySign, timer.TimerSign, timer.TimerCommand, auth)
	}
	return
}
func ListTimer(obj *model.TimerData, data *common.Data) (err error) {
	var W []string
	W = append(W, "timer.is_del = 0")
	if obj.State != "" {
		W = append(W, "timer.state = '"+obj.State+"'")
	}
	if obj.TimerName != "" {
		W = append(W, "timer.timer_name like '%"+obj.TimerName+"%'")
	}
	if obj.TimerDesc != "" {
		W = append(W, "timer.timer_desc like '%"+obj.TimerDesc+"%'")
	}
	//?????? ??????????????????????????????findinset
	if obj.TimerSign != "" {
		W = append(W, "timer.timer_sign like '%"+obj.TimerSign+"%'")
	}

	if obj.ClassifyName != "" {
		W = append(W, "classify.classify_name like '%"+obj.ClassifyName+"%'")
	}
	if obj.ClassifyDesc != "" {
		W = append(W, "classify.classify_desc like '%"+obj.ClassifyDesc+"%'")
	}
	if obj.ClassifySign != "" {
		W = append(W, "classify.classify_sign = '"+obj.ClassifySign+"'")
	}

	sqlC := "select count(*) from system_timer timer " +
		"left join system_timer_classify classify on timer.classify_id=classify.classify_id "

	sqlL := "select timer.*,IFNULL(classify.classify_name, '')as classify_name," +
		"IFNULL(classify.classify_desc, '')as classify_desc,IFNULL(classify.classify_sign, '')as classify_sign " +
		"from system_timer timer " +
		"left join system_timer_classify classify on timer.classify_id=classify.classify_id "
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	//start
	var list []*model.TimerData
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
func DelTimer(id string, auth *authU.GlobalConfig) (err error) {
	var timer model.Timer
	sqlstr1 := "select * from  system_timer where timer_id = ?"
	common.FindTable(&timer, sqlstr1, id)

	var classify model.TimerClassify
	sqlstr2 := "select * from  system_timer_classify where classify_id = ?"
	common.FindTable(&classify, sqlstr2, timer.ClassifyId)

	sqlstr := "UPDATE system_timer SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where timer_id = ? and is_del = 0"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		id)
	if err != nil {
		return
	}
	if timer.State == "1" {
		err = statusTimer(4, classify.ClassifySign, timer.TimerSign)
	} else {
		err = statusTimer(5, classify.ClassifySign, timer.TimerSign)
	}
	return
}

/**
step ???0????????? ????????????/????????????????????? ?????????????????????????????????????????????????????????
step ???1????????? ????????????/????????????????????? ???????????????????????????????????????????????????????????? ???????????????.stop

*/
func handleTimer(step int, path, name, content string, auth *authU.GlobalConfig) (err error) {
	cont := strings.Split(content, ",")
	var contentArr []string
	contentArr = append(contentArr, "#!/bin/bash")
	contentArr = append(contentArr, "if [ -f ~/.bash_profile ];")
	contentArr = append(contentArr, "then")
	contentArr = append(contentArr, "  . ~/.bash_profile")
	contentArr = append(contentArr, "fi")
	for _, v := range cont {
		contentArr = append(contentArr, v)
	}
	ext := "sh"
	if step == 1 {
		ext = "sh.stop"
	}
	err = common.BuildCommandFile(path, name, ext, contentArr, "CrontabConfig")
	//if err !=nil {
	//	return
	//}
	//err,_ = common.ExecCommand("service crond restart")//???????????????
	if err != nil {
		return
	}
	err = reStartTimer()
	return
}

/**
step ???2????????? ?????????????????????  .sh??????????????????.sh.stop??????
step ???3????????? ?????????????????????  .sh.stop?????????????????????.sh
step ???4????????? ?????????????????????  .sh ????????????.del??????????????????
step ???5????????? ?????????????????????  .sh.stop ????????????.del??????????????????
*/
func statusTimer(step int, path, name string) (err error) {
	ext := "sh"
	newExt := "sh.stop"
	if step == 3 {
		ext = "sh"
		newExt = "sh.stop"
	} else if step == 4 {
		nowStr := time.Now().Format("02_150405")
		ext = "sh"
		newExt = "sh.del" + nowStr
	} else if step == 5 {
		nowStr := time.Now().Format("02_150405")
		ext = "sh.stop"
		newExt = "sh.stop.del" + nowStr
	}
	err = common.RenameCommandFile(path, name, ext, newExt, "CrontabConfig")
	if err != nil {
		return
	}
	err = reStartTimer()
	return
}

//????????????????????????????????????crontab
func reStartTimer() (err error) {
	//?????? ??????????????????????????? ??????????????????.sh??????
	//?????????
	var obj model.TimerData
	obj.State = "1" //???????????????1??????????????????
	var data common.Data
	err = ListTimer(&obj, &data)
	list := make(map[string]string)
	//????????????
	listValue, ok := data.List.([]*model.TimerData)
	if ok {
		for k, v := range listValue {
			list[string(k)+"__"+v.TimerExec] = v.ClassifySign + "/" + v.TimerSign + ".sh"
		}
	} else {
		err = errors.New("???????????????????????????")
		if err != nil {
			return
		}
	}
	err = common.RestartCrontab(&list)
	return
}
