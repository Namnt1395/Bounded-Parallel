package handle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"rtp_demo/object"
	"sort"
	"strconv"
	"time"
)

//TODO for bid
func ActionBidding(w http.ResponseWriter, r *http.Request) {
	urls := []string{"http://localhost:9083/api-delay-100", "http://localhost:9083/api-delay-50"}
	// Init bidding request
	// init json
	biddingRequest := object.BiddingRequest{}
	biddingRequest.InitBiddingRequest()

	device := object.Device{}
	device.InitDevice(r)
	biddingRequest.Device = device
	bidJson, _ := json.Marshal(biddingRequest)
	resultsChan := make(chan *object.BiddingResponse)
	defer func() {
		close(resultsChan)
	}()
	for i, url := range urls {
		go func(i int, url string) {
			client := http.Client{
				Timeout: 120 * time.Millisecond,
			}
			//res, err := client.Get(url)
			res, err := client.Post(url, "application/json", bytes.NewBuffer(bidJson))
			data := &object.BiddingResponse{}
			if err ==  nil {
				json.NewDecoder(res.Body).Decode(&data)
			}
			resultsChan <- &object.BiddingResponse{Id: data.Id, BidId: data.BidId, Price: data.Price, CodeAds: data.CodeAds, Error: err}
		}(i, url)
	}
	var resultsSuccess []object.BiddingResponse
	var result []object.BiddingResponse
L:
	for {
		select {
		case data := <-resultsChan:
			result = append(result, *data)
			if data.Error == nil {
				resultsSuccess = append(resultsSuccess, *data)
			}
			if len(result) == len(urls) {
				break L
			}
		case <-time.After(time.Millisecond * 120):
			fmt.Println("time.Millisecond * 120", time.Millisecond*120)
			break L
		}
	}
	if len(resultsSuccess) <= 0{
		w.Write([]byte("pass pack"))
		return
	}
	// let's sort these results real quick
	sort.Slice(resultsSuccess, func(i, j int) bool {
		priceInt, _ :=  strconv.Atoi(resultsSuccess[i].Price)
		priceNextInt, _ :=  strconv.Atoi(resultsSuccess[j].Price)
		return priceInt > priceNextInt
	})
	fmt.Println("resultsSuccess", resultsSuccess)
	// display ads
	res, _ := json.Marshal(resultsSuccess[0].CodeAds)
	w.Write([]byte(res))
}
