package nsq

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nsqio/go-nsq"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	"gwsee.com.api/app/test/model"
	"log"
	"time"
)

type nsqHandler struct {
	nsqConsumer      *nsq.Consumer
	messagesReceived int
	nsqAction        interface{}
	data             *nsqData
}
type nsqData struct {
	topic   string
	channel string
	addr    string
	auth    *authU.GlobalConfig
}

//主要用于判断哪些是已经添加到消息里面 避免重复添加
var nsqMap map[uint64]*nsqHandler

//in one shell, start nsqlookupd:
//
//$ nsqlookupd
//in another shell, start nsqd:
//
//$ nsqd --lookupd-tcp-address=127.0.0.1:4160
//in another shell, start nsqadmin:
//
//$ nsqadmin --lookupd-http-address=127.0.0.1:4161
//publish an initial message (creates the topic in the cluster, too):
//
//$ curl -d 'hello world 1' 'http://127.0.0.1:4151/pub?topic=test'
//finally, in another shell, start nsq_to_file:
//
//$ nsq_to_file --topic=test --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161

//处理消息
func (nh *nsqHandler) HandleMessage(msg *nsq.Message) (err error) {
	nh.messagesReceived++
	if nh.data.topic == "testMessage" {
		action, ok := nsqMap[nh.data.auth.User.UserId].nsqAction.(*websocket.Conn)
		if ok {
			messageid := string(msg.Body)
			if messageid == "" {
				return
			}
			var message model.Message
			sqlstr := "select * from  test_message where message_id = ? and receiver_id=?  and is_del=0 and add_time+delay_seconds<=?"
			err = common.FindTable(&message, sqlstr, messageid, nh.data.auth.User.UserId, time.Now().Unix())
			if err != nil {
				fmt.Println(err.Error(), "消息错误")
				return
			}
			if message.MessageId > 0 {
				jsonStr := gin.H{"code": 1, "data": message, "msg": "ok"}
				res, _ := json.Marshal(&jsonStr)
				fmt.Println("写入消息：" + string(res))
				err = action.WriteMessage(1, res)
				if err != nil {
					fmt.Println("写入消息失败：" + err.Error())
					return
				}
			} else {
				fmt.Println("暂无消息：")
			}
			msg.Finish()
		}
	}
	return nil
}

func InitConsumer(topic, channel string, action interface{}, auth *authU.GlobalConfig) error {
	var nsqConnect nsqData
	nsqConnect.topic = topic
	nsqConnect.addr = "127.0.0.1:4161"
	nsqConnect.channel = channel
	nsqConnect.auth = auth
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = 3 * time.Second
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		log.Println("init Consumer NewConsumer error:", err)
		return err
	}
	_, ok := nsqMap[auth.User.UserId]
	handler := &nsqHandler{nsqConsumer: c, nsqAction: action, data: &nsqConnect}

	if ok {
		nsqMap[auth.User.UserId] = handler
		fmt.Println("已经添加进去过了")
	} else {
		c.AddHandler(handler)
		nsqMap[auth.User.UserId] = handler
		err = c.ConnectToNSQLookupd("127.0.0.1:4161")
		if err != nil {
			log.Println("init Consumer ConnectToNSQLookupd error:", err)
			return err
		}
	}

	return nil
}
