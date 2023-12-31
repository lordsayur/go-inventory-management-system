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

func (u *ItemUsecase) CreateItem(item *entities.Item) error {
	return u.itemRepo.Create(item)
}

func (u *ItemUsecase) GetById(id int) (*entities.Item, error) {
	return u.itemRepo.ReadById(id)
}

func (u *ItemUsecase) GetAllItems(sortField string, sortOrder string) ([]entities.Item, error) {
	return u.itemRepo.ReadAll(sortField, sortOrder)
}

func (u *ItemUsecase) UpdateItem(item *entities.Item) error {

	return u.itemRepo.Update(item)
}

func (u *ItemUsecase) DeleteItem(id int) error {
	return u.itemRepo.Delete(id)
}
