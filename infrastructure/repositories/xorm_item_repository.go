package repositories

import (
	"ims/core/entities"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type XORMItemRepository struct {
	engine *xorm.Engine
}

func NewXORMItemRepository() *XORMItemRepository {
	engine, err := xorm.NewEngine("sqlite3", "ims.db")
	if err != nil {
		log.Fatal("Error creating XORM engine:", err)
	}

	repo := &XORMItemRepository{
		engine: engine,
	}
	err = repo.CreateTable()
	if err != nil {
		log.Fatal("Error creating table", err)
	}

	return repo
}

func (r *XORMItemRepository) Create(item *entities.Item) error {
	_, err := r.engine.Insert(item)

	if err != nil {
		log.Println("Error creating item:", err)
		return err
	}

	return nil
}

func (r *XORMItemRepository) ReadAll(sortField string, sortOrder string) ([]entities.Item, error) {
	var items []entities.Item
	err := r.engine.Find(&items)

	if err != nil {
		log.Println("Error reading items:", err)
		return nil, err
	}

	return items, nil
}

func (r *XORMItemRepository) Update(item *entities.Item) error {

	return entities.ErrNotFound
}

func (r *XORMItemRepository) Delete(id int) error {

	return entities.ErrNotFound
}

func (r *XORMItemRepository) CreateTable() error {
	err := r.engine.Sync2(new(entities.Item))
	if err != nil {
		log.Println("Error creating table:", err)
		return err
	}

	return nil
}
