package model

type Welcome struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Status bool   `json:"status"`
}
