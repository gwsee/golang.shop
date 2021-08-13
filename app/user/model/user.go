package model

import (
	"gwsee.com.api/app/common"
)

// 用于获取个人基本信息的结构体
type User struct {
	common.DbColumn
	UserBase
	UserData
}

// var GlobalUser User
//不可修改的
type UserBase struct {
	UserId      uint64 `db:"user_id"        form:"userid" json:"userid"`
	UserAccount string `db:"user_account"   form:"account" json:"account"`
}

//可以修改的 --通过JSON绑定赋值
type UserData struct {
	UserName          string `db:"user_name"      form:"name" json:"name"`
	UserPassword      string `db:"user_password"  form:"password" json:"password"`
	UserHash          string `db:"user_hash"      json:"-"`
	UserAvatar        string `db:"user_avatar"    form:"avatar" json:"avatar"`
	UserMobile        string `db:"user_mobile"    form:"mobile" json:"mobile"`
	UserPoints        uint64 `db:"user_points"    form:"points" json:"points"`
	UserAddress       string `db:"user_address"   form:"address" json:"address"`
	UserAutograph     string `db:"user_autograph" form:"autograph" json:"autograph"`
	UserEmail         string `db:"user_email"     form:"email" json:"email"`
	UserSex           string `db:"user_sex"       form:"sex" json:"sex"`
	UserWechat        string `db:"user_wechat"    form:"wechat" json:"wechat"`
	UserWechatUnionid string `db:"user_wechat_unionid"    form:"wechat_unionid" json:"wechat_unionid"`
	Idcard            string `db:"idcard"         form:"idcard" json:"idcard"`
	IdcardName        string `db:"idcard_name"    form:"idcardname" json:"idcardname"`
	IdcardNum         string `db:"idcard_num"     form:"idcardnum" json:"idcardnum"`
}
