package model

import(
	"encoding/xml"
)

// CommonParams 组装请求公共参数
type CommonParams struct {
	// 公共字段
	Version       string                `xml:"version" url:"version" validate:"nonzero"`       //版本号。固定为1.0
	InsCD         string                `xml:"ins_cd"  url:"ins_cd"  validate:"nonzero"`       //富友给咱的机构号
	MchntCD       string                `xml:"mchnt_cd" url:"mchnt_cd" validate:"nonzero"`     //富友给商户的商户号
	TermID        string                `xml:"term_id"  url:"term_id" validate:"nonzero"`      //富友给的终端号
	RandomStr     string                `xml:"random_str" url:"random_str" validate:"nonzero"` // 随机字符串
	Sign          string                `xml:"sign" url:"-" `                                  // 签名
	ReqTimeString string                `xml:"-" url:"-" bson:"-"` // 请求发送时间
}




//富友返回报文公共部分
type CommonBody struct {
	ResultCode     string `xml:"result_code" url:"result_code"` // 错误代码  000000成功
	ResultMsg      string `xml:"result_msg" url:"result_msg"`   // 错误代码描述
	InsCD          string `xml:"ins_cd"  url:"ins_cd"`          //富友给咱的机构号
	MchntCD        string `xml:"mchnt_cd" url:"mchnt_cd"`       //富友给商户的商户号
	TermID         string `xml:"term_id"  url:"term_id"`        //富友给的终端号
	RandomStr      string `xml:"random_str" url:"random_str"`   // 随机字符串
	Sign           string `xml:"sign" url:"-" `                 // 签名
	RespTimeString string `xml:"-" url:"-" bson:"-"`            // 请求应答时间
}



//PayReq 富友支付条码支付请求
type FypPayReq struct{
	XMLName   xml.Name `xml:"xml" url:"-"`
	CommonParams
	OrderType string `xml:"order_type" url:"order_type" validate:"nonzero"`      //订单类型ALIPAY,WECHAT,JD,QQ,UNIONPAY
	GoodsDes  string `xml:"goods_des"  url:"goods_des"  validate:"nonzero"`      //商品描述
	GoodsDetail string `xml:"goods_detail" url:"goods_detail"`						 //商品详情
	AddnInf    string `xml:"addn_inf" url:"addn_inf"` 							 //附加数据
	MchntOrderNo string `xml:"mchnt_order_no" url:"mchnt_order_no" validate:"nonzero"` //商户订单号
	CurrType string `xml:"curr_type" url:"curr_type"`							 //货币种类 默认人民币:CNY
	OrderAmt string `xml:"order_amt" url:"order_amt" validate:"nonzero"`         //订单总金额
	TermIP string `xml:"term_ip" url:"term_ip" validate:"nonzero"`               //终端IP            
	TxnBeginTs string `xml:"txn_begin_ts" url:"txn_begin_ts" validate:"nonzero"` //交易开始时间  yyyyMMddHHmmss
	GoodsTag string `xml:"goods_tag" url:"goods_tag" `							 //商品标记
	AuthCode string `xml:"auth_code" url:"auth_code" validate:"nonzero"`         //授权码
	Sence string `xml:"sence" url:"sence"`										 //支付场景，默认1.1:扫码支付2:声波支付
	ReservedSubAppid string `xml:"reserved_sub_appid" url:"-"`					 //子商户公众号
	ReservedLimitPay string `xml:"reserved_limit_pay" url:"-"` 					 //限制支付no_credit：不能使用信用卡
	ReservedExpireMinute string `xml:"reserved_expire_minute" url:"-" validate:"nonzero"` //交易关闭时间单位分钟.不设置默认填0
	ReservedFyTermID string `xml:"reserved_fy_term_id" url:"-"`           //富友终端号
}


//PayResp 富友条码支付应答
type FypPayResp struct{
	XMLName   xml.Name `xml:"xml" url:"-"`
	CommonBody
	OrderType string `xml:"order_type" url:"order_type"`				         //订单类型ALIPAY,WECHAT,JD,QQ,UNIONPAY
	TotalAmount string `xml:"total_amount" url:"total_amount"`				     //订单金额 分为单位
	BuyerID   string `xml:"buyer_id" url:"buyer_id"`				             //买家在渠道账户
	TransactionID string `xml:"transaction_id" url:"transaction_id"`		     //渠道交易流水号
	AddnInf    string `xml:"addn_inf" url:"addn_inf"` 							 //附加数据
	ReservedFyOrderNo string `xml:"reserved_fy_order_no" url:"-"` 	             //富友订单号
	ReservedMchntOrderNo string `xml:"reserved_mchnt_order_no" url:"-"`			 //咱们的订单号
	ReservedFySettleDt string `xml:"reserved_fy_settle_dt" url:"-"`  		  	 //富友清算日
	ReservedCouponFee string `xml:"reserved_coupon_fee" url:"-"`                 //优惠金额（分）            
	ReservedBuyerLogonID string `xml:"reserved_buyer_logon_id" url:"-" ` 	     //买家在渠道登录的账号
	ReservedFundBillList string `xml:"reserved_fund_bill_list" url:"-" `		 //支付宝交易资金渠道
	ReservedFyTraceNo string `xml:"reserved_fy_trace_no" url:"-" `               //富友系统内部追踪号
	ReservedChannelOrderID string `xml:"reserved_channel_order_id" url:"-"`		 //条码流水号，用户账单二维码对应的流水								 
	ReservedIsCredit string `xml:"reserved_is_credit" url:"-"`				     //1标识信用卡或者花呗 0标识其他  不填标识未知
}