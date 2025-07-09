package persistence

import (
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	firebaseSync "github.com/Enryu5/Warehouse-Inventory-App/backend/internal/infrastructure/firebase"
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
	if err := r.db.Create(item).Error; err != nil {
		return err
	}
	return firebaseSync.SyncItemToFirebase(item)
}

func (r *itemRepository) Update(item *domain.Item) error {
	if err := r.db.Save(item).Error; err != nil {
		return err
	}
	return firebaseSync.SyncItemToFirebase(item)
}

func (r *itemRepository) Delete(id int) error {
	if err := r.db.Delete(&domain.Item{}, id).Error; err != nil {
		return err
	}
	return firebaseSync.DeleteItemFromFirebase(id)
}
