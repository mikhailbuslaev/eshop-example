package model

type DeliveryInfo struct {
	Address string `json:"address"`
	Zip string `json:zip`
	Receiver string `json:receiver`
}

type Order struct {
	Id string `json:"id"`
	UserId string `json:"-"`
	Delivery DeliveryInfo `json:"-"`
	Items []CartItem `json:"-"`
	Cost float64 `json:"-"`
	JsonData string `json:data`
}
