package repositories

import (
	"ims/core/entities"
	"sync"
)

type MemoryItemRepository struct {
	mu     sync.Mutex
	items  map[int]entities.Item
	lastID int
}

func NewMemoryItemRepository() *MemoryItemRepository {
	return &MemoryItemRepository{
		items:  make(map[int]entities.Item),
		lastID: 0,
	}
}

func (r *MemoryItemRepository) Create(item *entities.Item) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.lastID++
	item.ID = r.lastID
	r.items[item.ID] = *item

	return nil
}

func (r *MemoryItemRepository) ReadAll() ([]entities.Item, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var items []entities.Item
	for _, item := range r.items {
		items = append(items, item)
	}

	return items, nil
}
