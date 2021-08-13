package service

import (
	"errors"
	"fmt"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/common/nsq"
	"gwsee.com.api/app/test/model"
	"strconv"
	"strings"
	"time"
)

//删除
func DelMessage(messageid string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE test_message SET " +
		"is_del=? " +
		", edit_time=? ,edit_user=? " +
		"where message_id = ?"
	_, err = common.UpdateTable(sqlstr,
		time.Now().Unix(),
		time.Now().Unix(), auth.User.UserId,
		messageid)
	//消息删除之后 就不能再次发送消息了
	return
}

/**
DelaySeconds  使用nsq的延迟处理进行消息 延迟加载到数据库（延迟调用addMessage
然后插入成功之后 使用websocket进行消息发送到前端的消息列表上面
*/
func AddMessage(message *model.Message, auth *authU.GlobalConfig) (id int64, err error) {
	message.SenderId = auth.User.UserId
	fmt.Println(message.ReceiverId, "message.ReceiverId改变之前")
	if message.ReceiverId < 1 {
		message.ReceiverId = auth.User.UserId
	}
	fmt.Println(message.ReceiverId, "message.ReceiverId改变之后")
	if message.MessageId > 0 {
		//消息不允许修改
		return 0, errors.New("消息不能修改")
		sqlstr := "UPDATE test_message SET message=?,sender_id=?,receiver_id=?,delay_seconds=?," +
			" edit_time=?, edit_user=?, state=? " +
			" where message_id=?"
		_, err = common.UpdateTable(sqlstr,
			message.Message, message.SenderId, message.ReceiverId, message.DelaySeconds,
			time.Now().Unix(), auth.User.UserId, message.State,
			message.MessageId)
	} else {
		sqlstr := "INSERT INTO test_message (add_time,add_user,state," +
			"message,sender_id,receiver_id,delay_seconds,shop_sn) VALUES (?,?,?," +
			"?,?,?,?,?)"
		if message.ShopSn == "" {
			message.ShopSn = auth.ShopSn
		}
		id, err = common.InsertTable(sqlstr, time.Now().Unix(), auth.User.UserId, message.State,
			message.Message, message.SenderId, message.ReceiverId, message.DelaySeconds, message.ShopSn)
		if id > 0 {
			var t = time.Duration(message.DelaySeconds) * time.Second
			nsq.NsqProducers.Publish("testMessage", t, strconv.FormatInt(id, 10))
		}
	}
	//插入成功之后要发送消息给前端
	return
}
func ReadMessage(messageid string, auth *authU.GlobalConfig) (err error) {
	sqlstr := "UPDATE test_message SET " +
		"status=? " +
		", edit_time=? ,edit_user=? " +
		"where message_id = ?"
	_, err = common.UpdateTable(sqlstr,
		'2',
		time.Now().Unix(), auth.User.UserId,
		messageid)
	return
}

func ListMessage(message *model.MessageLoad, data *common.Data, auth *authU.GlobalConfig) (err error) {
	var W []string
	W = append(W, "is_del = 0")
	W = append(W, "shop_sn = '"+auth.ShopSn+"'")
	W = append(W, "(receiver_id = '"+strconv.FormatUint(auth.User.UserId, 10)+"' or sender_id='"+strconv.FormatUint(auth.User.UserId, 10)+"')") //只能读取本人的消息
	if message.Message != "" {
		W = append(W, "message like '%"+message.Message+"%'")
	}
	if message.State != "" {
		W = append(W, "state = '%"+message.State+"%'")
	}
	order := " order by add_time desc "
	if message.MessageId > 0 {
		if message.MessageType == "before" { //拉取这个数据之前的数据
			order = " order by add_time desc "
			W = append(W, "message_id < '"+strconv.FormatUint(message.MessageId, 10)+"'")
		} else { //默认是拉取最新的
			order = " order by add_time asc "
			W = append(W, "message_id > '"+strconv.FormatUint(message.MessageId, 10)+"'")
		}
	}

	sqlC := "select count(*) from test_message"
	sqlL := "select * from test_message"
	if W != nil {
		sqlC = sqlC + " where " + strings.Join(W, " and ")
		sqlL = sqlL + " where " + strings.Join(W, " and ")
	}
	var list []*model.Message
	data.Total, err = common.CountTable(sqlC)
	if err != nil {
		return
	}
	limit := common.BuildCount(data)
	sqlL = sqlL + order + limit
	err = common.ListTable(&list, sqlL)
	//list数据要根据时间进行升序排序  正顺序 最新的在最下面方便前端加载

	if err != nil {
		return
	}
	data.List = list
	return
}

//  获取单个消息
func FindMessage(info *model.Message, i string) (err error) {
	sqlstr := "select * from  test_message where message_id = ? and is_del=0 and add_time+delay_seconds<?"
	err = common.FindTable(info, sqlstr, i, time.Now().Unix())
	if err != nil {
		return
	}
	return
}
