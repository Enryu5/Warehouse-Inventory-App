package repository

import "github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"

type WarehouseRepository interface {
	GetAll() ([]domain.Warehouse, error)
	GetByID(id int) (*domain.Warehouse, error)
	Create(warehouse *domain.Warehouse) error
	Update(warehouse *domain.Warehouse) error
	Delete(id int) error
}
