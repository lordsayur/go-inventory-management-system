package interfaces

import (
	"ims/core/entities"
)

type ItemRepository interface {
	Create(item *entities.Item) error
	ReadAll(sortField string, sortOrder string) ([]entities.Item, error)
	Update(item *entities.Item) error
	Delete(id int) error
}
