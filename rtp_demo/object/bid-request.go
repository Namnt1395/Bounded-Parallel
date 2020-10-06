package object
/**
- IP của user
- cookie của user
- Tên đầy đủ trình duyệt đang sử dụng
 */
type BiddingRequest struct {
	Id     string   `json:"id"`
	At     int      `json:"at"`
	Cur    []string `json:"cur"`
	Device Device   `json:"device"`
	Tmax   int      `json:"tmax"`
}

func (b *BiddingRequest) InitBiddingRequest() {
	b.At = 1
	b.Id = "5700"
	b.Cur = []string{"USD"}
	b.Tmax = 2500
}
