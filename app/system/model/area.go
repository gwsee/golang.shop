package model

import "gwsee.com.api/app/common"

type Area struct {
	common.DbColumn
	AreaId   uint64 `db:"area_id" form:"areaid" json:"areaid"`
	AreaName string `db:"area_name" form:"name" json:"name"`
	AreaPid  string `db:"area_pid" form:"pid" json:"pid"`
}
