package model

import "gwsee.com.api/app/common"

type SystemMessage struct {
	common.DbColumn
	MessageId     uint64 `db:"message_id"    form:"id" json:"id"`
	MessageType   uint64 `db:"message_type"  form:"type" json:"type"`
	MessageCode   uint64 `db:"message_code"  form:"code" json:"code"`
	MessageMobile string `db:"message_mobile"   form:"mobile" json:"mobile"`
	MessageExpire uint64 `db:"message_expire"   form:"expire" json:"expire"`
}
