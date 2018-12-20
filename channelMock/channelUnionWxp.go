package channelMock

import (
	"github.com/CardInfoLink/bubble-gum/channelMock/model"
	"encoding/xml"
	"github.com/CardInfoLink/log"
)

func UnionWxpRefundServive(req *model.MybankReq) []byte {
	log.Debugf("[rcv req]%+v", req)
	UnionwxpPayResp :=  &model.UnionWxpPayResp {
		UnionWxpCommonBody: model.UnionWxpCommonBody {
			ReturnCode : "SUCCESS" ,
			ResultCode : "FAIL" ,
			ErrCode : "SYSTEMERROR",
		},
	}

	bytes, _ := xml.Marshal(UnionwxpPayResp)
	return bytes	
}


func UnionWxpPayServive(req *model.WxpPayReq) []byte {
	log.Debugf("[rcv req]%+v", req)
	UnionwxpPayResp :=  &model.UnionWxpPayResp {
		UnionWxpCommonBody: model.UnionWxpCommonBody {
			ReturnCode : "SUCCESS" ,
			ResultCode : "SUCCESS" ,
			ErrCode : "",
		},
		TradeState : "SUCCESS",		
	}

	bytes, _ := xml.Marshal(UnionwxpPayResp)
	return bytes
}

func UnionWxpRefundQueryServive(req *model.MybankReq) []byte {
	log.Debugf("[rcv req]%+v", req)
	UnionwxpPayResp :=  &model.UnionWxpPayResp {
		UnionWxpCommonBody: model.UnionWxpCommonBody {
			ReturnCode : "SUCCESS" ,
			ResultCode : "FAIL" ,
			ErrCode : "SYSTEMERROR",
		},
	}

	bytes, _ := xml.Marshal(UnionwxpPayResp)
	return bytes	
}
