package repositories

import (
	"ims/core/entities"
	"sort"
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

func (r *MemoryItemRepository) ReadById(id int) (*entities.Item, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	item, ok := r.items[id]

	if !ok {
		return nil, entities.ErrNotFound
	}

	return &item, nil
}

func (r *MemoryItemRepository) ReadAll(sortField string, sortOrder string) ([]entities.Item, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var items []entities.Item
	for _, item := range r.items {
		items = append(items, item)
	}

	switch sortField {
	case "name":
		sort.Slice(items, func(i, j int) bool {
			if sortOrder == "desc" {
				return items[i].Name > items[j].Name
			}

			return items[i].Name < items[j].Name
		})
	}

	return items, nil
}

func (r *MemoryItemRepository) Update(item *entities.Item) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.items[item.ID]; ok {
		r.items[item.ID] = *item
		return nil
	}

	return entities.ErrNotFound
}

func (r *MemoryItemRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.items[id]; ok {
		delete(r.items, id)
		return nil
	}

	return entities.ErrNotFound
}
