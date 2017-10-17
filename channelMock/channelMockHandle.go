package channelMock

import (
	"net/http"
	"encoding/json"
	"github.com/CardInfoLink/bubble-gum/channelMock/model"
	"time"
	"github.com/CardInfoLink/log"
)

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
	time.Sleep(30000 * time.Millisecond)
	log.Debugf("[send resp]%+s", respBytes)
	w.Write(respBytes)
	return
}

func WxpHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("to do"))
	return
}
