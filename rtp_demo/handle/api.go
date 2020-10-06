package handle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rtp_demo/object"
	"time"
)

func ApiDelay50MS(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data object.BiddingRequest
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error decode", err.Error())
	}
	time.Sleep(time.Millisecond * 50)
	var response object.BiddingResponse
	response.Price = "1000"
	response.CodeAds = "Delay 50 ms"
	response.BidId = "800"
	//response := "delay 50 ms"
	bidResponse, _ := json.Marshal(response)
    w.Write([]byte(bidResponse))
}
func ApiDelay100MS(w http.ResponseWriter,req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data object.BiddingRequest
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error decode", err.Error())
	}
	time.Sleep(time.Millisecond * 80)
	var response object.BiddingResponse
	response.Price = "6000"
	response.CodeAds = "Delay 100 ms"
	response.BidId = "800"
	bidResponse, _ := json.Marshal(response)
	w.Write([]byte(bidResponse))
}
