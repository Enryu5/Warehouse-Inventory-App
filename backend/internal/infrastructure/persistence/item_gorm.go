package persistence

import (
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/repository"
	"gorm.io/gorm"
)

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) repository.ItemRepository {
	return &itemRepository{db}
}

func (r *itemRepository) GetAll() ([]domain.Item, error) {
	var items []domain.Item
	err := r.db.Find(&items).Error
	return items, err
}

func (r *itemRepository) GetByID(id int) (*domain.Item, error) {
	var item domain.Item
	err := r.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) Create(item *domain.Item) error {
	return r.db.Create(item).Error
}

func (r *itemRepository) Update(item *domain.Item) error {
	return r.db.Save(item).Error
}

func (r *itemRepository) Delete(id int) error {
	return r.db.Delete(&domain.Item{}, id).Error
}
