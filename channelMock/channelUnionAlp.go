package channelMock

import (
	"strconv"
	"time"
	"github.com/CardInfoLink/bubble-gum/channelMock/model"
	"encoding/json"
	"github.com/CardInfoLink/log"
)

type UnionAlpPayResp struct {
	Resp model.UnionAlpCommonResponse `json:"alipay_trade_pay_response"`
}

type UnionAlpPreCreateResp struct {
	Resp model.UnionAlpCommonResponse `json:"alipay_trade_precreate_response"`
}

type UnionAlpQueryResp struct {
	Resp model.UnionAlpCommonResponse `json:"alipay_trade_query_response"`
}
//by sunny
//定义支付宝退款应答结构体
type UnionAlpRefundResp struct {
	Resp model.UnionAlpExceptionResponse `json:"alipay_trade_refund_response"`
}
var UnionAlpCommonResp_Success = model.UnionAlpCommonResponse{
	Code:"10000",
	Msg:"Success",
	TradeStatus:"TRADE_SUCCESS",
	TradeNo:strconv.FormatInt(time.Now().Unix(), 10),
}
//by sunny
//定义支付宝回服务器异常的应答
var UnionAlpResp_Exception = model.UnionAlpExceptionResponse{
	Code:"20000",
	Msg:"支付宝服务器异常",
	TradeStatus:"TRADE_Error",
}	
var UnionAlpPayResp_Success = UnionAlpPayResp{
	Resp:UnionAlpCommonResp_Success,
}

var UnionAlpPreCreateResp_Process = UnionAlpPreCreateResp{
	Resp:UnionAlpCommonResp_Success,
}

var UnionAlpQueryResp_Process = UnionAlpQueryResp{
	Resp:UnionAlpCommonResp_Success,
}
//by sunny
//定义支付宝退款异常应答处理
var UnionAlpRefundResp_Process = UnionAlpRefundResp{
	Resp:UnionAlpResp_Exception,
}
// json: 	alipay_trade_pay_response
func UnionAlpServive(req *model.UnionAlpComonRequest) []byte {
	log.Debugf("[rcv req]%+v", req)

	switch req.Method {
	case "alipay.trade.query":
		return UnionAlpQueryService(req)
	case "alipay.trade.precreate":
		return UnionAlpPreCreateService(req)
	case "alipay.trade.pay":
		return UnionAlpPayService(req)
    //by sunny
	case "alipay.trade.refund":
	    return UnionAlpRefundService(req)
	}
	return nil
}

func UnionAlpPayService(req *model.UnionAlpComonRequest) []byte {
	bytes, _ := json.Marshal(UnionAlpPayResp_Success)
	return bytes
}

func UnionAlpPreCreateService(req *model.UnionAlpComonRequest) []byte {
	bytes, _ := json.Marshal(UnionAlpPreCreateResp_Process)
	return bytes
}

func UnionAlpQueryService(req *model.UnionAlpComonRequest) []byte {
	bytes, _ := json.Marshal(UnionAlpQueryResp_Process)
	return bytes
}
//by sunny
//增加支付宝退款接口处理
func UnionAlpRefundService(req *model.UnionAlpComonRequest) []byte {
	bytes, _ := json.Marshal(UnionAlpRefundResp_Process)
	return bytes
}