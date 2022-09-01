package model

type DeliveryInfo struct {
	Address string `json:"address"`
	Zip string `json:zip`
	Receiver string `json:receiver`
}

type Order struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
	Delivery DeliveryInfo `json:"delivery"`
	Items []CartItem `json:"items"`
	Cost float64 `json:"cost"`
	JsonData string `json:-`
}
