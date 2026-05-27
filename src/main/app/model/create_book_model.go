package model

type CreateBookModel struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int64  `json:"quantity"`
}
