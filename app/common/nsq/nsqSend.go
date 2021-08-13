package nsq

import (
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

type producer struct {
	producer *nsq.Producer
}

var NsqProducers *producer

//服务启动
func InitNsqProducer() {
	NsqProducers = &producer{}
	nsqMap = make(map[uint64]*nsqHandler, 2)
	NsqProducers.producer, _ = initProducer("127.0.0.1:4150")
}
func initProducer(addr string) (p *nsq.Producer, err error) {
	var (
		config *nsq.Config
	)
	config = nsq.NewConfig()
	if p, err = nsq.NewProducer(addr, config); err != nil {
		return nil, err
	}
	return p, nil
}

//消息发布
func (p *producer) Publish(topic string, delay time.Duration, message string) (err error) {
	if message == "" {
		return errors.New("message is empty")
	}
	if delay > 0 {
		if err = p.producer.DeferredPublish(topic, delay, []byte(message)); err != nil {
			fmt.Println(err)
			return err
		}
	} else {
		if err = p.producer.Publish(topic, []byte(message)); err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
