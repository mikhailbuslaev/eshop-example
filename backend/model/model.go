package model

type Card struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Count       int    `json:"count"`
	PicturePath string `json:"picturepath"`
}

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