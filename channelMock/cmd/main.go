package main

import (
	"flag"
	"net/http"
	"github.com/CardInfoLink/bubble-gum/channelMock"
	"fmt"
)

func main() {
	startMock()
}

func startMock() {
	flag.IntVar(&channelMock.MbpSleep, "mbpSleep", 0, "Sleep [mbpSleep] ms before mybank responds")
	flag.Parse()

	http.HandleFunc("/mock/alp", channelMock.AlpHandle)
	http.HandleFunc("/mock/unionalp", channelMock.UnionAlpHandle)
	http.HandleFunc("/mock/wxp/secapi/pay/refund", channelMock.WxpRefundHandle)
	http.HandleFunc("/mock/wxp/pay/micropay", channelMock.WxpPayHandle)
	http.HandleFunc("/mock/wxp/pay/unifiedorder", channelMock.WxpPrePayHandle)
	http.HandleFunc("/mock/wxp//pay/refundquery", channelMock.WxpRefundQueryHandle)
	http.HandleFunc("/mock/unionwxp/pay/refund", channelMock.UnionWxpRefundHandle)
	http.HandleFunc("/mock/unionwxp/pay/micropay", channelMock.UnionWxpPayHandle)
	http.HandleFunc("/mock/unionwxp//pay/refundquery", channelMock.UnionWxpRefundQueryHandle)	
	http.HandleFunc("/mock/mbp", channelMock.MbpHandle)
	http.HandleFunc("/mock/fyp/micropay", channelMock.FypHandle)
	http.HandleFunc("/mock/fyp/commonRefund", channelMock.FypHandle)
	if err := http.ListenAndServe(":9900", nil); err != nil {
		fmt.Printf("%s\n", err)
	}
}
