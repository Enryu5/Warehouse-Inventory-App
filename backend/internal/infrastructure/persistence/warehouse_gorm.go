package persistence

import (
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	firebaseSync "github.com/Enryu5/Warehouse-Inventory-App/backend/internal/infrastructure/firebase"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/repository"
	"gorm.io/gorm"
)

type warehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) repository.WarehouseRepository {
	return &warehouseRepository{db}
}

func (r *warehouseRepository) GetAll() ([]domain.Warehouse, error) {
	var warehouses []domain.Warehouse
	err := r.db.Find(&warehouses).Error
	return warehouses, err
}

func (r *warehouseRepository) GetByID(id int) (*domain.Warehouse, error) {
	var warehouse domain.Warehouse
	err := r.db.First(&warehouse, id).Error
	if err != nil {
		return nil, err
	}
	return &warehouse, nil
}

func (r *warehouseRepository) Create(warehouse *domain.Warehouse) error {
	if err := r.db.Create(warehouse).Error; err != nil {
		return err
	}
	return firebaseSync.SyncWarehouseToFirebase(warehouse)
}

func (r *warehouseRepository) Update(warehouse *domain.Warehouse) error {
	if err := r.db.Save(warehouse).Error; err != nil {
		return err
	}
	return firebaseSync.SyncWarehouseToFirebase(warehouse)
}

func (r *warehouseRepository) Delete(id int) error {
	if err := r.db.Delete(&domain.Warehouse{}, id).Error; err != nil {
		return err
	}
	return firebaseSync.DeleteWarehouseFromFirebase(id)
}
