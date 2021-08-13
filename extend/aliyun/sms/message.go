package sms

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type smsOss struct {
	regionId        string
	accessKeyId     string
	accessKeySecret string
}
type SmsObject struct {
	Mobile   string
	Sign     string
	Template string
	Code     string
}

func Send(obj *SmsObject) (err error) {
	var sms smsOss
	sms.regionId = "cn-hangzhou" //非必须 默认就行
	sms.accessKeyId = ""
	sms.accessKeySecret = ""
	client, _ := dysmsapi.NewClientWithAccessKey(sms.regionId, sms.accessKeyId, sms.accessKeySecret)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = obj.Mobile
	request.SignName = obj.Sign         //动想随行
	request.TemplateCode = obj.Template //"SMS_211494683"
	request.TemplateParam = "{code:" + obj.Code + "}"
	response, err := client.SendSms(request)
	if err != nil {
		return err
	}
	fmt.Printf("response is %#v\n", response)
	return
}
