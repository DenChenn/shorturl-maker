package model

type UrlEntity struct {
	Id       string `json:"id"`
	ShortUrl string `json:"shortUrl"`
}

type UrlEntityInDB struct {
	Id       string `json:"id"`
	Url      string `json:"url"`
	ExpireAt string `json:"expireAt"`
	CreateAt string `json:"createAt"`
}
