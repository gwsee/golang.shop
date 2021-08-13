package model

import "gwsee.com.api/app/common"

type Timer struct {
	common.DbColumn
	TimerId      uint64 `db:"timer_id" form:"timerid" json:"timerid"`
	TimerName    string `db:"timer_name"  form:"name" json:"name"`
	TimerDesc    string `db:"timer_desc"  form:"desc" json:"desc"`
	TimerSign    string `db:"timer_sign"  form:"sign" json:"sign"`
	TimerExec    string `db:"timer_exec"  form:"exec" json:"exec"`
	TimerCommand string `db:"timer_command"   form:"command" json:"command"`
	ClassifyId   uint64 `db:"classify_id"   form:"classifyid" json:"classifyid"`
}

type TimerLog struct {
	common.DbColumn
	LogId        uint64 `db:"log_id" form:"logid" json:"logid"`
	TimerId      uint64 `db:"timer_id" form:"timerid" json:"timerid"`
	LogRes       string `db:"log_res" form:"res" json:"res"`
	LogCommand   string `db:"log_command" form:"command" json:"command"`
	TimerName    string `db:"timer_name"  form:"name" json:"name"`
	TimerSign    string `db:"timer_sign"  form:"sign" json:"sign"`
	ClassifyName string `db:"classify_name"  form:"classifyname" json:"classifyname"`
	ClassifySign string `db:"classify_sign"  form:"classifysign" json:"classifysign"`
}
type TimerClassify struct {
	common.DbColumn
	ClassifyId   uint64 `db:"classify_id" form:"classifyid" json:"classifyid"`
	ClassifyName string `db:"classify_name"  form:"name" json:"name"`
	ClassifyDesc string `db:"classify_desc"  form:"desc" json:"desc"`
	ClassifySign string `db:"classify_sign"  form:"sign" json:"sign"`
}

type TimerData struct {
	Timer
	ClassifyName string `db:"classify_name"  form:"classifyname" json:"classifyname"`
	ClassifyDesc string `db:"classify_desc"  form:"classifydesc" json:"classifydesc"`
	ClassifySign string `db:"classify_sign"  form:"classifysign" json:"classifysign"`
}
