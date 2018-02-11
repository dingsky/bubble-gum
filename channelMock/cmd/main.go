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
	http.HandleFunc("/mock/wxp", channelMock.WxpHandle)
	http.HandleFunc("/mock/mbp", channelMock.MbpHandle)
	http.HandleFunc("/mock/fyp/micropay", channelMock.FypHandle)
	if err := http.ListenAndServe(":9900", nil); err != nil {
		fmt.Printf("%s\n", err)
	}
}
