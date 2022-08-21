package model

type CartItem struct {
	Id string `json:"id"`
	CardId string `json:"cardid"`
	Count int `json:"count"`
}

type ShoppingCart struct {
	Id string `json:"id"`
	Items []CartItem `json:"items"`
	JsonItems []byte
}