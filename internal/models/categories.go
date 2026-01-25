package models

type Category struct {
	Id           int64  `json:"id"`
	UserId       int64  `json:"user_id"`
	CategoryName string `json:"category_name"`
}
