package persistence

import (
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
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
	return r.db.Create(warehouse).Error
}

func (r *warehouseRepository) Update(warehouse *domain.Warehouse) error {
	return r.db.Save(warehouse).Error
}

func (r *warehouseRepository) Delete(id int) error {
	return r.db.Delete(&domain.Warehouse{}, id).Error
}
