package service

import (
	"errors"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/system/model"
	"gwsee.com.api/extend/aliyun/sms"
	"gwsee.com.api/utils"
	"strconv"
	"time"
)

//新增
func AddMessage(message *model.SystemMessage, auth *authU.GlobalConfig) (id int64, err error) {
	if message.MessageType < 1 {
		message.MessageType = 1
	}

	//先查询是否有短信
	var info model.SystemMessage
	err = FindMessage(&info, message.MessageMobile, message.MessageType)

	if err != nil {
		return 0, err
	}
	if info.MessageId > 0 {
		err = errors.New("短信已发送 敬请期待")
		return
	}

	sqlstr := "INSERT INTO shop_message (add_time,add_user,state," +
		"message_type,message_mobile,message_code,message_expire) VALUES (?,?,?," +
		"?,?,?,?)"
	if message.MessageType == 1 {
		//短信验证码
		m, _ := time.ParseDuration("5m")
		result := time.Now().Add(m)
		message.MessageExpire = uint64(result.Unix())
	} else {
		m, _ := time.ParseDuration("10m")
		result := time.Now().Add(m)
		message.MessageExpire = uint64(result.Unix())
	}
	//生成code  并且短信发送
	code := utils.Krand(4, 0)
	//短信发送成功再加入数据库
	var obj sms.SmsObject
	obj.Code = code
	obj.Template = "SMS_211494683"
	obj.Sign = "动想随行"
	obj.Mobile = message.MessageMobile

	err = sms.Send(&obj)
	if err != nil {
		return 0, err
	}
	id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, 1,
		message.MessageType, message.MessageMobile, code, message.MessageExpire)

	return
}

func ValidMessage(message *model.SystemMessage) (err error) {
	var info model.SystemMessage
	err = FindMessage(&info, message.MessageMobile, message.MessageType)
	if err != nil {
		return err
	}
	if info.MessageId < 1 {
		err = errors.New("短信验证失败")
		return
	}
	if info.MessageCode != message.MessageCode {
		err = errors.New("短信验证不匹配")
		return
	}
	//存在这个短信 就进行短信状态修改
	err = SetMessage(strconv.FormatUint(info.MessageId, 10), "2")
	//验证成功
	return
}

//查询
func FindMessage(message *model.SystemMessage, messageMobile string, messageType uint64) (err error) {
	sqlstr := "select * from shop_message where message_mobile=? and message_type=?  and state=? and message_expire>?"
	err = common.FindTable(message, sqlstr, messageMobile, messageType, 1, time.Now().Unix()) //过期短信不用查出来
	return
}

//删除
func DelMessage(messageid string) (err error) {
	sqlstr := "UPDATE shop_message SET " +
		"is_del=? " +
		", edit_time=? " +
		"where message_id = ?"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(),
		messageid)
	return
}

//状态修改
func SetMessage(messageid, state string) (err error) {
	sqlstr := "UPDATE shop_message SET " +
		"state=? " +
		", edit_time=? " +
		"where message_id = ?"
	_, err = common.UpdateTable(sqlstr,
		state,
		time.Now().Unix(),
		messageid)
	return
}
