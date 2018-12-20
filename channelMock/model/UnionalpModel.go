package model

type UnionAlpComonRequest struct {
	Method string `json:"method"`
}

type UnionAlpCommonResponse struct {
	Sign        string `json:"sign" bson:"sign"`                             // 签名
	Code        string `json:"code" bson:"code"`                             // 结果码
	Msg         string `json:"msg" bson:"msg"`                               // 结果码描述
	SubCode     string `json:"sub_code,omitempty" bson:"sub_code,omitempty"` // 错误子代码
	SubMsg      string `json:"sub_msg,omitempty"  bson:"sub_msg,omitempty"`  // 错误子代码描述
	TradeStatus string `json:"trade_status" bson:"trade_status"`			//
	TradeNo     string `json:"trade_no" bson:"trade_no"`                     //渠道訂單號
}
type UnionAlpExceptionResponse struct {
	Sign        string `json:"sign" bson:"sign"`                             // 签名
	Code        string `json:"code" bson:"code"`                             // 结果码
	Msg         string `json:"msg" bson:"msg"`                               // 结果码描述
	SubCode     string `json:"sub_code,omitempty" bson:"sub_code,omitempty"` // 错误子代码
	SubMsg      string `json:"sub_msg,omitempty"  bson:"sub_msg,omitempty"`  // 错误子代码描述
	TradeStatus string `json:"trade_status" bson:"trade_status"`
	TradeNo     string `json:"trade_no" bson:"trade_no"`                     //渠道訂單號
}