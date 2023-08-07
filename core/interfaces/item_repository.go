package interfaces

import (
	"ims/core/entities"
)

type ItemRepository interface {
	Create(item *entities.Item) error
	ReadAll() ([]entities.Item, error)
}
