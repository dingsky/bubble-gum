package channelMock

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"encoding/xml"
	"github.com/CardInfoLink/bubble-gum/channelMock/model"
	"time"
	"github.com/CardInfoLink/log"
)

var MbpSleep = 0

func AlpHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vals := r.Form
	m := make(map[string]string, len(vals))
	for k, v := range vals {
		m[k] = v[0]
	}
	log.Debugf("[rcv req]%+v", m)

	bytes, _ := json.Marshal(m)
	req := model.AlpComonRequest{}
	//log.Debugf("[rcv req]%s", bytes)
	json.Unmarshal(bytes, &req)
	respBytes := alpServive(&req)
	time.Sleep(300 * time.Millisecond)
	log.Debugf("[send resp]%+s", respBytes)
	w.Write(respBytes)
	return
}

func WxpRefundHandle(w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("wxp mock error"))
		return
	}
	defer r.Body.Close()
	log.Debugf("[rcv req]%s", string(reqData))
	req := &model.MybankReq {
	}
	// err = xml.Unmarshal(reqData, req)
	// if err != nil {
	// 	w.Write([]byte("wxp Unmarshal mock error"))
	// 	return
	// }
	respBytes := WxpRefundServive(req)
	log.Debugf("[send resp]%+s", respBytes)
	w.Write(respBytes)

	return
}

func WxpPayHandle(w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("wxp mock error"))
		return
	}
	defer r.Body.Close()
	log.Debugf("[rcv req]%s", string(reqData))
	req := &model.WxpPayReq {
	}
	err = xml.Unmarshal(reqData, req)
	if err != nil {
		w.Write([]byte("wxp Unmarshal mock error"))
		return
	}
	respBytes := WxpPayServive(req)
	log.Debugf("[send resp]%+s", respBytes)
	w.Write(respBytes)

	return
}

func WxpPrePayHandle(w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("wxp mock error"))
		return
	}
	defer r.Body.Close()
	log.Debugf("[rcv req]%s", string(reqData))
	req := &model.MybankReq {
	}
	// err = xml.Unmarshal(reqData, req)
	// if err != nil {
	// 	w.Write([]byte("wxp Unmarshal mock error"))
	// 	return
	// }
	respBytes := WxpPrePayServive(req)
	log.Debugf("[send resp]%+s", respBytes)
	w.Write(respBytes)

	return
}

func WxpRefundQueryHandle(w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("wxp mock error"))
		return
	}
	defer r.Body.Close()
	log.Debugf("[rcv req]%s", string(reqData))
	req := &model.MybankReq {
	}
	// err = xml.Unmarshal(reqData, req)
	// if err != nil {
	// 	w.Write([]byte("wxp Unmarshal mock error"))
	// 	return
	// }
	respBytes := WxpRefundQueryServive(req)
	log.Debugf("[send resp]%+s", respBytes)
	w.Write(respBytes)

	return
}

func MbpHandle(w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("mybank mock error"))
		return
	}
	defer r.Body.Close()
	log.Debugf("[rcv req]%s", string(reqData))
	req := &model.MybankReq {
		Sign: &model.MybankSignature{},
	}
	err = xml.Unmarshal(reqData, req)
	if err != nil {
		w.Write([]byte("mybank Unmarshal mock error"))
		return
	}
	respBytes := mbpServive(req)
	if MbpSleep > 0 {
		time.Sleep(time.Duration(MbpSleep) * time.Millisecond)
	}
	log.Debugf("[send resp]%+s", respBytes)
	w.Write(respBytes)
	return
}


func FypHandle(w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("fuiou mock error"))
		return
	}
	defer r.Body.Close()
	log.Debugf("[rcv req]%s", string(reqData))
	req := &model.FypPayReq{
	}
	err = xml.Unmarshal(reqData, req)
	if err != nil {
		w.Write([]byte("fuiou Unmarshal mock error"))
		return
	}
	respBytes := fypServive(req)
	if MbpSleep > 0 {
		time.Sleep(time.Duration(MbpSleep) * time.Millisecond)
	}
	log.Debugf("[send resp]%+s", respBytes)
	w.Write(respBytes)
	return
}
