package channelMock

import (
	"github.com/CardInfoLink/bubble-gum/channelMock/model"
	"encoding/xml"
	"github.com/CardInfoLink/log"
)

func WxpRefundServive(req *model.MybankReq) []byte {
	log.Debugf("[rcv req]%+v", req)
	wxpPayResp :=  &model.WxpPayResp {
		WxpCommonBody: model.WxpCommonBody {
			ReturnCode : "SUCCESS" ,
			ResultCode : "FAIL" ,
			ErrCode : "SYSTEMERROR",
		},
	}

	bytes, _ := xml.Marshal(wxpPayResp)
	return bytes	
}


func WxpPayServive(req *model.WxpPayReq) []byte {
	log.Debugf("[rcv req]%+v", req)
	wxpPayResp :=  &model.WxpPayResp {
		WxpCommonBody: model.WxpCommonBody {
			ReturnCode : "SUCCESS" ,
			ResultCode : "FAIL" ,
			ErrCode : "SYSTEMERROR",
		},
	}

	bytes, _ := xml.Marshal(wxpPayResp)
	return bytes
}

func WxpRefundQueryServive(req *model.MybankReq) []byte {
	log.Debugf("[rcv req]%+v", req)
	wxpPayResp :=  &model.WxpPayResp {
		WxpCommonBody: model.WxpCommonBody {
			ReturnCode : "SUCCESS" ,
			ResultCode : "FAIL" ,
			ErrCode : "SYSTEMERROR",
		},
	}

	bytes, _ := xml.Marshal(wxpPayResp)
	return bytes	
}

func WxpPrePayServive(req *model.MybankReq) []byte {
	log.Debugf("[rcv req]%+v", req)
	wxpPayResp :=  &model.WxpPayResp {
		WxpCommonBody: model.WxpCommonBody {
			ReturnCode : "SUCCESS" ,
			ResultCode : "SUCCESS" ,
			ErrCode : "",
		},
	}

	bytes, _ := xml.Marshal(wxpPayResp)
	return bytes	
}

