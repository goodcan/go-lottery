package models

type GiftPrize struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	PrizeNum     int    `json:"-"`
	LeftNum      int    `json:"-"`
	PrizeCodeA   int    `json:"-"`
	PrizeCodeB   int    `json:"-"`
	Img          string `json:"img"`
	DisplayOrder int    `json:"display_order"`
	Gtype        int    `json:"gtype"`
	Gdata        string `json:"gdata"`
}
