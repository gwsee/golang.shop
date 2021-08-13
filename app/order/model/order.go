package model

import "gwsee.com.api/app/common"

// 订单状态
var OrderState = map[string]string{
	"to_be_paid":             "待支付",
	"to_be_delivered":        "待发货", //已支付
	"to_be_received":         "待收货",
	"to_be_served":           "待服务", //已支付
	"refund_ing":             "退款中",
	"refund_failed":          "退款失败",
	"refund_success":         "退款成功",
	"transaction_finished":   "交易完成",
	"transaction_settlement": "交易结算", //不可退款了
}

// 订单类型
var DetailType = map[string]string{
	"service":  "服务",
	"pay":      "支付",
	"sign":     "报名",
	"activity": "活动",
	"vip":      "会员",
	"charge":   "充值",
}

type Order struct {
	common.DbColumn

	OrderId      uint64  `db:"order_id" form:"orderid" json:"orderid"`
	OrderNo      string  `db:"order_no" form:"orderno" json:"orderno"`
	OrderGrade   uint64  `db:"order_grade" form:"ordergrade" json:"ordergrade"`
	OrderAmount  float64 `db:"order_amount" form:"orderamount" json:"orderamount"`
	OrderPay     float64 `db:"order_pay" form:"orderpay" json:"orderpay"`
	OrderType    string  `db:"order_type" form:"ordertype" json:"ordertype"`
	OrderReward  float64 `db:"order_reward" form:"orderreward" json:"orderreward"`
	OrderComment uint64  `db:"order_comment" form:"ordercomment" json:"ordercomment"`
	OrderRemarks string  `db:"order_remarks" form:"orderremarks" json:"orderremarks"`
	OrderVerify  uint64  `db:"order_verify" form:"orderverify" json:"orderverify"`
	OrderFrom    string  `db:"order_from" form:"orderfrom" json:"orderfrom"`

	UserId      uint64 `db:"user_id" form:"userid" json:"userid"`
	UserReferee uint64 `db:"user_referee" form:"userreferee" json:"userreferee"`
	UserVerify  uint64 `db:"user_verify" form:"userverify" json:"userverify"`

	PayTime   uint64 `db:"pay_time" form:"paytime" json:"paytime"`
	PayType   string `db:"pay_type" form:"paytype" json:"paytype"`
	PayNo     string `db:"pay_no" form:"payno" json:"payno"`
	PayMobile string `db:"pay_mobile" form:"paymobile" json:"paymobile"`

	ShopSn     string `db:"shop_sn"   form:"shopsn" json:"shopsn"`
	FinishTime uint64 `db:"finish_time" form:"finishtime" json:"finishtime"`
	VerifyTime uint64 `db:"verify_time" form:"verifytime" json:"verifytime"`
}

type OrderDetail struct {
	common.DbColumn
	DetailId        uint64  `db:"detail_id" form:"detailid" json:"detailid"`
	ShopSn          string  `db:"shop_sn"   form:"shopsn" json:"shopsn"`
	OrderNo         string  `db:"order_no" form:"orderno" json:"orderno"`
	GoodsSn         string  `db:"goods_sn" form:"goodssn" json:"goodssn"`
	GoodsTitle      string  `db:"goods_title" form:"goodstitle" json:"goodstitle"`
	GoodsPrice      float64 `db:"goods_price" form:"goodsprice" json:"goodsprice"`
	GoodsNum        uint64  `db:"goods_num" form:"goodsnum" json:"goodsnum"`
	DetailPrice     float64 `db:"detail_price" form:"detailprice" json:"detailprice"`
	DetailPay       float64 `db:"detail_pay" form:"detailpay" json:"detailpay"`
	DetailType      string  `db:"detail_type" form:"detailtype" json:"detailtype"`
	DetailTransport uint64  `db:"detail_transport" form:"detailtransport" json:"detailtransport"`
}

type OrderService struct {
	common.DbColumn
	ServiceId      uint64 `db:"service_id" form:"serviceid" json:"serviceid"`
	ShopSn         string `db:"shop_sn"   form:"shopsn" json:"shopsn"`
	OrderNo        string `db:"order_no" form:"orderno" json:"orderno"`
	GoodsSn        string `db:"goods_sn" form:"goodssn" json:"goodssn"`
	UserName       string `db:"user_name" form:"username" json:"username"`
	UserMobile     string `db:"user_mobile" form:"usermobile" json:"usermobile"`
	ServiceServer  string `db:"service_server" form:"serviceserver" json:"serviceserver"`
	ServiceMinutes uint64 `db:"service_minutes" form:"serviceminutes" json:"serviceminutes"`
	ServiceBegin   uint64 `db:"service_begin" form:"servicebegin" json:"servicebegin"`
	ServiceEnd     uint64 `db:"service_end" form:"serviceend" json:"serviceend"`
}

type OrderData struct {
	OrderDetail
	OrderGrade   uint64  `db:"order_grade" form:"ordergrade" json:"ordergrade"`
	OrderRemarks string  `db:"order_remarks" form:"orderremarks" json:"orderremarks"`
	OrderPay     float64 `db:"order_pay" form:"orderpay" json:"orderpay"`
	OrderFrom    string  `db:"order_from" form:"orderfrom" json:"orderfrom"`
	UserId       uint64  `db:"user_id" form:"userid" json:"userid"`
	UserName     string  `db:"user_name" form:"username" json:"username"`
	PayTime      uint64  `db:"pay_time" form:"paytime" json:"paytime"`
	PayType      string  `db:"pay_type" form:"paytype" json:"paytype"`
	StateName    string  `json:"statusname"`
	TypeName     string  `json:"typename"`
}

type OrderServiceData struct {
	OrderService
	GoodsTitle  string  `db:"goods_title" form:"goodstitle" json:"goodstitle"`
	GoodsPrice  float64 `db:"goods_price" form:"goodsprice" json:"goodsprice"`
	GoodsNum    uint64  `db:"goods_num" form:"goodsnum" json:"goodsnum"`
	DetailPay   float64 `db:"detail_pay" form:"detailpay" json:"detailpay"`
	DetailType  string  `db:"detail_type" form:"detailtype" json:"detailtype"`
	DetailState string  `db:"detail_tate" form:"detailtate" json:"detailtate"`
	StateName   string  `json:"statusname"`
	TypeName    string  `json:"typename"`
}
