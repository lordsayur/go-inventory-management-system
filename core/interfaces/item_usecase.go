package interfaces

import "ims/core/entities"

type ItemUsecase interface {
	CreateItem(name string) error
	GetAllItems() ([]entities.Item, error)
}
