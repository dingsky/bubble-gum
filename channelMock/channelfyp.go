package channelMock

import (
	"fmt"
	"github.com/CardInfoLink/bubble-gum/channelMock/model"
	"encoding/xml"
	"github.com/CardInfoLink/log"
	"time"
)

func mbpServive(req *model.FypPayReq) []byte {
	log.Debugf("[rcv req]%+v", req)
	return fypPayService(req)
}

func fypPayService(req *model.FypPayReq) []byte {
	fypResp := &model.FypPayResp {
		CommonBody:model.CommonBody{
			ResultCode:"000000",
			ResultMsg:"SUCCESS",
			InsCD: req.CommonParams.InsCD,
			MchntCD:req.CommonParams.MchntCD,
			TermID:req.CommonParams.TermID,
			RandomStr:req.CommonParams.RandomStr,
		}

		OrderType:req.OrderType,
		TotalAmount:req.OrderAmt,
		AddnInf：req.AddnInf,
		ReservedMchntOrderNo:req.MchntOrderNo,
	}
	bytes, _ := xml.Marshal(fypResp)
	return bytes
}
