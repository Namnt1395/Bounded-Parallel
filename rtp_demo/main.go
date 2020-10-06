package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"rtp_demo/handle"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api-delay-50", handle.ApiDelay50MS).Methods("POST")
	router.HandleFunc("/api-delay-100", handle.ApiDelay100MS).Methods("POST")
	router.HandleFunc("/bidding", handle.ActionBidding)

	err := http.ListenAndServe(":9083", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
