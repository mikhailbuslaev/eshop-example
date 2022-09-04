package model

type Card struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Count       int    `json:"count"`
	PicturePath string `json:"picturepath"`
}
