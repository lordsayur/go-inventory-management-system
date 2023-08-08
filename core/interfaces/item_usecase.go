package interfaces

import "ims/core/entities"

type ItemUsecase interface {
	CreateItem(name string) error
	GetAllItems(sortField string, sortOrder string) ([]entities.Item, error)
	UpdateItem(item *entities.Item) error
	DeleteItem(id int) error
}
