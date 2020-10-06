package object

type BiddingResponse struct {
	Id      string `json:"id"`
	BidId   string `json:"bidid"`
	Price   string    `json:"price"`
	CodeAds string `json:"code_ads"`
	Error error `json:"error"`
}
