package repository

import "github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"

type StockRepository interface {
	GetByWarehouse(warehouseID int) ([]domain.Stock, error)
	GetByItem(itemID int) ([]domain.Stock, error)
	Upsert(stock *domain.Stock) error // Insert or Update
	DeleteByItemAndWarehouse(itemID, warehouseID int) error
}
