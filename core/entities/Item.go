package entities

import "errors"

var (
	ErrNotFound = errors.New("item not found")
)

type Item struct {
	ID          int     `json:"id" xorm:"'id' pk autoincr"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float32 `json:"price"`
}
