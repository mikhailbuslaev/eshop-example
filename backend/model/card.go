package model

type Card struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Count       int    `json:"count"`
	PicturePath string `json:"picturepath"`
}