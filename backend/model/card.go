package model

type Card struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Firm        string `json:"firm"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Count       int    `json:"count"`
	PicturePath string `json:"picturepath"`
}
