package channelMock

import (
	"github.com/CardInfoLink/bubble-gum/channelMock/model"
	"encoding/json"
	"github.com/CardInfoLink/log"
)

type alpPayResp struct {
	Resp model.AlpCommonResponse `json:"alipay_trade_pay_response"`
}

type alpPreCreateResp struct {
	Resp model.AlpCommonResponse `json:"alipay_trade_precreate_response"`
}

type alpQueryResp struct {
	Resp model.AlpCommonResponse `json:"alipay_trade_query_response"`
}
//by sunny
//定义支付宝退款应答结构体
type alpRefundResp struct {
	Resp model.AlpCommonResponse `json:"alipay_trade_refund_response"`
}
var alpCommonResp_Success = model.AlpCommonResponse{
	Code:"10000",
	Msg:"Success",
	TradeStatus:"TRADE_SUCCESS",
}
//by sunny
//定义支付宝回服务器异常的应答
var alpResp_Exception = model.AlpExceptionResponse{
	Code:"20000",
	Msg:"支付宝服务器异常",
	TradeStatus:"TRADE_Error",
}	
var alpPayResp_Success = alpPayResp{
	Resp:alpCommonResp_Success,
}

var alpPreCreateResp_Process = alpPreCreateResp{
	Resp:alpCommonResp_Success,
}

var alpQueryResp_Process = alpQueryResp{
	Resp:alpCommonResp_Success,
}
//by sunny
//定义支付宝退款异常应答处理
var alpRefundResp_Process = alpRefundResp{
	Resp:alpResp_Exception,
}
// json: 	alipay_trade_pay_response
func alpServive(req *model.AlpComonRequest) []byte {
	log.Debugf("[rcv req]%+v", req)

	switch req.Method {
	case "alipay.trade.query":
		return alpQueryService(req)
	case "alipay.trade.precreate":
		return alpPreCreateService(req)
	case "alipay.trade.pay":
		return alpPayService(req)
    //by sunny
	case "alipay.trade.refund":
	    return alpRefundService(req)
	}
	return nil
}

func alpPayService(req *model.AlpComonRequest) []byte {
	bytes, _ := json.Marshal(alpPayResp_Success)
	return bytes
}

func alpPreCreateService(req *model.AlpComonRequest) []byte {
	bytes, _ := json.Marshal(alpPreCreateResp_Process)
	return bytes
}

func alpQueryService(req *model.AlpComonRequest) []byte {
	bytes, _ := json.Marshal(alpQueryResp_Process)
	return bytes
}
//by sunny
//增加支付宝退款接口处理
func alpRefundService(req *model.AlpComonRequest) []byte {
	bytes, _ := json.Marshal(alpRefundResp_Process)
	return bytes
}