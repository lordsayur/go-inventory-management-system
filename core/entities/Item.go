package entities

import "errors"

var (
	ErrNotFound = errors.New("item not found")
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
