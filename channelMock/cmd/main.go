package main

import (
	"net/http"
	"github.com/CardInfoLink/bubble-gum/channelMock"
	"fmt"
)

func main() {
	startMock()
}

func startMock() {
	http.HandleFunc("/mock/alp", channelMock.AlpHandle)
	http.HandleFunc("/mock/wxp", channelMock.WxpHandle)
	if err := http.ListenAndServe(":9900", nil); err != nil {
		fmt.Printf("%s\n", err)
	}
}
