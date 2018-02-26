package model

import (
	"encoding/xml"
)

//公共字段
type WxpCommonParams struct {
	XMLName xml.Name `xml:"xml"`

	// 公共字段
	Appid    string `xml:"appid,omitempty"`           // 微信分配的公众账号ID validate:"len=18
	SubAppid string `xml:"sub_appid,omitempty"`   // 微信分配的子商户公众账号ID
	MchID    string `xml:"mch_id,omitempty"`         // 微信支付分配的商户号
	SubMchId string `xml:"sub_mch_id,omitempty""` // 微信支付分配的子商户号，开发者模式下必填
	NonceStr string `xml:"nonce_str"`    // 随机字符串
	Sign     string `xml:"sign"`                                    // 签名

	WeixinMD5Key string `xml:"-"`

	ClientCert []byte                `xml:"-"` // HTTPS 双向认证证书
	ClientKey  []byte                `xml:"-"` // HTTPS 双向认证密钥
	ReqTimeString  string	         `xml:"-"`// 请求发送时间

}

//公共返回字段
type WxpCommonBody struct {
	XMLName xml.Name `xml:"xml"`

	ReturnCode string `xml:"return_code"`                   // 返回状态码
	ReturnMsg  string `xml:"return_msg,omitempty"` // 返回信息

	// 当 return_code 为 SUCCESS 的时候，还会包括以下字段：
	Appid      string `xml:"appid,omitempty"`               // 公众账号ID,小程序下单的话是返回的小程序id
	MchID      string `xml:"mch_id,omitempty"`             // 商户号
	SubMchId   string `xml:"sub_mch_id,omitempty"`     // 子商户号
	SubAppid   string `xml:"sub_appid,omitempty"`       // 子商户公众账号 ID
	NonceStr   string `xml:"nonce_str"`                           // 随机字符串
	Sign       string `xml:"sign"`                                        // 签名
	ResultCode string `xml:"result_code"`                       // 业务结果
	ErrCode    string `xml:"err_code,omitempty"`         // 错误代码
	ErrCodeDes string `xml:"err_code_des,omitempty"` // 错误代码描述
	RespTimeString string                  `xml:"-"`// 请求应答时间
}

type WxpResqDetail struct {
	XmlDetail      DetailData  `xml:"detail,omitempty"`                    // 商品详情 XML数据
	UrlDetail      string     `xml:"-"` 			//签名数据  有效数据
}
type DetailData struct{
	DetailDataString  string   `xml:",cdata"`  //<!CDATA[...]>
}

type WxpPayReq struct {
	WxpCommonParams

	DeviceInfo     string `xml:"device_info,omitempty"`          // 设备号
	Body           string `xml:"body"`                         	// 商品描述
	WxpResqDetail
	Attach         string `xml:"attach,omitempty"`              // 附加数据
	OutTradeNo     string `xml:"out_trade_no"`         			// 商户订单号
	TotalFee       string `xml:"total_fee"`               // 总金额
	FeeType        string `xml:"fee_type,omitempty"`                // 货币类型
	SpbillCreateIP string `xml:"spbill_create_ip"` // 终端IP
	GoodsGag       string `xml:"goods_tag,omitempty"`              // 商品标记
	LimitPay       string `xml:"limit_pay,omitempty"`              // 指定支付方式
	AuthCode       string `xml:"auth_code"`               // 授权码
	// AuthCode       string `xml:"auth_code" url:"auth_code" validate:"regexp=^1\\d{17}$"` // 授权码

	TimeStart  string `xml:"time_start,omitempty"`   // 交易起始时间
	TimeExpire string `xml:"time_expire,omitempty"` // 交易结束时间
	Version    string `xml:"version,omitempty"` // 接口版本号 固定填写1.0
}




// PayResp 被扫支付提交Post数据给到API之后，API会返回XML格式的数据，这个类用来装这些数据
type WxpPayResp struct {
	WxpCommonBody

	DeviceInfo     string `xml:"device_info,omitempty"`     // 设备号
	Openid         string `xml:"openid"`                         // 用户标识
	IsSubscribe    string `xml:"is_subscribe"`             // 是否关注公众账号
	TradeType      string `xml:"trade_type"`                 // 交易类型
	BankType       string `xml:"bank_type"`                   // 付款银行
	FeeType        string `xml:"fee_type,omitempty"`           // 货币类型
	TotalFee       string `xml:"total_fee"`                   // 总金额
	CashFeeType    string `xml:"cash_fee_type,omitempty"` // 现金支付货币类型
	CashFee        string `xml:"cash_fee"`                     // 现金支付金额
	CouponFee      string `xml:"coupon_fee,omitempty"`       // 代金券或立减优惠金额
	TransactionId  string `xml:"transaction_id"`         // 微信支付订单号
	OutTradeNo     string `xml:"out_trade_no"`             // 商户订单号
	Attach         string `xml:"attach,omitempty"`               // 商家数据包
	TimeEnd        string `xml:"time_end"`                     // 支付完成时间
	CouponId0      string `xml:"coupon_id_0,omitempty"`     // 代金券或立减优惠ID
	CouponFee0     string `xml:"coupon_fee_0,omitempty"`   // 代金券或立减优惠退款金额
	CouponId1      string `xml:"coupon_id_1,omitempty"`     // 代金券或立减优惠ID
	CouponFee1     string `xml:"coupon_fee_1,omitempty"`   // 代金券或立减优惠退款金额
	CouponId2      string `xml:"coupon_id_2,omitempty"`     // 代金券或立减优惠ID
	CouponFee2     string `xml:"coupon_fee_2,omitempty"`   // 代金券或立减优惠退款金额
	SubOpenid      string `xml:"sub_openid,omitempty"`       // 子商户 Open ID
	SubIsSubscribe string `xml:"sub_is_subscribe"`     // 是否关注子商户公众账号
	SettlementTotalFee string `xml:"settlement_total_fee"` //应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	PromotionDetail    string `xml:"promotion_detail,omitempty"`         // 营销详情列表

}