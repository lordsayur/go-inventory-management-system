package usecases

import (
	"ims/core/entities"
	"ims/core/interfaces"
)

type ItemUsecase struct {
	itemRepo interfaces.ItemRepository
}

func NewItemUsecase(itemRepo interfaces.ItemRepository) *ItemUsecase {
	return &ItemUsecase{
		itemRepo: itemRepo,
	}
}

func (u *ItemUsecase) CreateItem(name string) error {
	item := &entities.Item{
		Name: name,
	}

	return u.itemRepo.Create(item)
}

func (u *ItemUsecase) GetAllItems() ([]entities.Item, error) {
	return u.itemRepo.ReadAll()
}
