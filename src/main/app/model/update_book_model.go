package model

type UpdateBookModel struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int64  `josn:"quantity"`
}
