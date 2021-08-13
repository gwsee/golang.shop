package service

import (
	"errors"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	modelSystem "gwsee.com.api/app/system/model"
	serviceSystem "gwsee.com.api/app/system/service"
	"gwsee.com.api/app/user/model"
	"gwsee.com.api/utils"
	"strconv"
	"time"
)

// GetInfo 获取个人信息的基本方法
func GetInfo(info *model.User, i string) (err error) {
	sqlstr := "select * from  user where is_del =0 AND user_account = ?"
	err = common.FindTable(info, sqlstr, i)
	return
}

// GetInfoByMobile 获取个人信息的基本方法
func GetInfoByMobile(info *model.User, i string) (err error) {
	sqlstr := "select * from  user where is_del =0 AND user_mobile = ?"
	err = common.FindTable(info, sqlstr, i)
	return
}

// CheckAccount 登陆 查询此人的账号密码 是否正确
func CheckAccount(account, password string) (user model.User, err error) {
	//查询账号
	err = GetInfo(&user, account)
	if err != nil {
		return
	}
	if user.UserId < 1 {
		err = errors.New("账号不存在")
		return
	}
	//验证密码
	screct := utils.HmacSha256(password, user.UserHash)
	if screct != user.UserPassword {
		err = errors.New("密码不正确")
	}
	return
}

// CheckCode 登陆，电话号码进行登陆的时候是不需要验证的
func CheckCode(mobile, code string) (user model.User, err error) {
	err = GetInfoByMobile(&user, mobile)
	if err != nil {
		return
	}
	if user.UserId < 1 {
		//代表要注册
		err = GetInfo(&user, mobile)
		if user.UserId > 0 {
			err = errors.New("请联系商家验证个人信息")
			return
		}
		if user.UserId < 1 {
			user.UserMobile = mobile
			user.UserAccount = mobile
			id, err := AddUser(&user, nil, false)
			if err != nil {
				return model.User{}, err
			}
			user.UserId = uint64(id)
		}

	}
	var mess modelSystem.SystemMessage
	mess.MessageCode, _ = strconv.ParseUint(code, 0, 64)
	mess.MessageMobile = mobile
	mess.MessageType = 1
	err = serviceSystem.ValidMessage(&mess)
	return
}

// CheckUnicode wxunicode 在一个微信开放平台（一个开发者平台）下的多个小程序 都是相同的unicode，所以；只需要把他丢到user表里就行了，一个公司应该不会有多个微信开放平台
func CheckUnicode(mobile, unicode string) (user model.User, err error) {
	err = GetInfoByMobile(&user, mobile)
	if err != nil {
		return
	}
	if user.UserId < 1 {
		//代表要注册
		err = GetInfo(&user, mobile)
		if user.UserId > 0 {
			err = errors.New("请联系商家验证个人信息")
			return
		}
		if user.UserId < 1 {
			user.UserName = "靓仔"+utils.Krand(2,3)
			user.UserMobile = mobile
			user.UserAccount = mobile
			user.UserWechatUnionid = unicode
			id, err := AddUser(&user, nil, false)
			if err != nil {
				return model.User{}, err
			}
			user.UserId = uint64(id)
		}
	}

	if user.UserWechatUnionid == ""&&unicode!="" {
		// 将 unionid 添加到用户信息中
		sqlstr := "UPDATE user set " +
			"user_wechat_unionid=?," +
			"edit_time =?  " +
			"where user_account = ?"
		_, err = common.UpdateTable(sqlstr,
			unicode,
			time.Now().Unix(),
			user.UserAccount)
	} else if user.UserWechatUnionid != unicode {
		//信息异常 需要提示
		err = errors.New("unicode异常，请联系商家")
	}
	return
}

//  创建用户信息 INSERT INTO user_info(username,sex,email)VALUES (?,?,?)  f代表是否需要判定
func AddUser(user *model.User, auth *authU.GlobalConfig, f bool) (id int64, err error) {
	//检测账号是否存在
	var info model.User
	if f {
		err = GetInfo(&info, user.UserAccount)
		if err != nil {
			return
		}
	}
	if info.UserId > 0 {
		err = EditUser(user, info.UserId, auth)
		// uint64 转 int64
		strInt64 := strconv.FormatUint(info.UserId, 10)
		id, _ = strconv.ParseInt(strInt64, 10, 64)
	} else {
		//然后进行插入
		sqlstr := "INSERT INTO user " +
			"(add_time,user_account,user_name,user_password," +
			"user_hash,user_avatar,user_mobile,user_address," +
			"user_autograph,user_email,user_sex,user_wechat,user_wechat_unionid," +
			"idcard,idcard_num,idcard_name)" +
			"VALUES (?,?,?,?," +
			"?,?,?,?," +
			"?,?,?,?,?," +
			"?,?,?)"
		hash := utils.Krand(6, 3)
		password := user.UserPassword
		if user.UserPassword == "" {
			password = "888888"
		}
		password = utils.HmacSha256(password, hash)
		id, err = common.InsertTable(sqlstr,
			time.Now().Unix(), user.UserAccount, user.UserName, password,
			hash, user.UserAvatar, user.UserMobile, user.UserAddress,
			user.UserAutograph, user.UserEmail, user.UserSex, user.UserWechat, user.UserWechatUnionid,
			user.Idcard, user.IdcardNum, user.IdcardName)
	}
	return
}

//  修改用户信息
func EditUser(data *model.User, userid uint64, auth *authU.GlobalConfig) (err error) {
	if data.UserPassword != "" {
		hash := utils.Krand(6, 3)
		password := utils.HmacSha256(data.UserPassword, hash)
		sqlstr := "UPDATE user set " +
			"user_name=? , user_avatar =? , user_mobile = ? , user_points=? " +
			", user_address=? , user_autograph =? , user_email=? , user_wechat=? " +
			", user_sex=? , idcard=? ,idcard_num =? , idcard_name=? " +
			", edit_time= ? , edit_user= ? ,user_password= ? ,user_hash= ? " +
			"where user_id = ?"
		_, err = common.UpdateTable(sqlstr,
			data.UserName, data.UserAvatar, data.UserMobile, data.UserPoints,
			data.UserAddress, data.UserAutograph, data.UserEmail, data.UserWechat,
			data.UserSex, data.Idcard, data.IdcardNum, data.IdcardName,
			time.Now().Unix(), auth.User.UserId, password, hash,
			userid)
	} else {
		sqlstr := "UPDATE user set " +
			"user_name=? , user_avatar =? , user_mobile = ? , user_points=? " +
			", user_address=? , user_autograph =? , user_email=? , user_wechat=? " +
			", user_sex=? , idcard=? ,idcard_num =? , idcard_name=? " +
			", edit_time = ? , edit_user=? " +
			"where user_id = ?"
		_, err = common.UpdateTable(sqlstr,
			data.UserName, data.UserAvatar, data.UserMobile, data.UserPoints,
			data.UserAddress, data.UserAutograph, data.UserEmail, data.UserWechat,
			data.UserSex, data.Idcard, data.IdcardNum, data.IdcardName,
			time.Now().Unix(), auth.User.UserId,
			userid)
	}

	return
}

//  修改用户密码
func EditPs(password, account, mobile, code string, auth *authU.GlobalConfig) (err error) {

	//验证电话号码和账户是否是一个人
	var user model.User
	err = GetInfo(&user, account)
	if err != nil {
		return
	}
	if user.UserMobile != mobile {
		err = errors.New("账户与电话号码不匹配")
		return
	}
	//然后验证验证码是否正确
	var mess modelSystem.SystemMessage

	mess.MessageCode, _ = strconv.ParseUint(code, 0, 64)
	mess.MessageMobile = mobile
	mess.MessageType = 1

	err = serviceSystem.ValidMessage(&mess)
	if err != nil {
		return
	}

	//如果是没修改人的信息的话 就修改自己
	hash := utils.Krand(6, 3)
	password = utils.HmacSha256(password, hash)
	sqlstr := "UPDATE user set " +
		"user_password=? , user_hash =? ," +
		"edit_time =? , edit_user=?  " +
		"where user_account = ?"
	_, err = common.UpdateTable(sqlstr,
		password, hash,
		time.Now().Unix(), auth.User.UserId,
		account)
	return
}
