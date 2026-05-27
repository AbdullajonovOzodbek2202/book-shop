package model

type Book struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int64  `json:"quantity"`
}
