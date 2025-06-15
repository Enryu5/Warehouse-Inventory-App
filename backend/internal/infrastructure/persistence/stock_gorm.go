package persistence

import (
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/repository"
	"gorm.io/gorm"
)

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) repository.StockRepository {
	return &stockRepository{db}
}

func (r *stockRepository) GetByWarehouse(warehouseID int) ([]domain.Stock, error) {
	var stocks []domain.Stock
	err := r.db.Where("warehouse_id = ?", warehouseID).Find(&stocks).Error
	return stocks, err
}

func (r *stockRepository) GetByItem(itemID int) ([]domain.Stock, error) {
	var stocks []domain.Stock
	err := r.db.Where("item_id = ?", itemID).Find(&stocks).Error
	return stocks, err
}

func (r *stockRepository) Upsert(stock *domain.Stock) error {
	// This implements "upsert" (insert or update) logic
	return r.db.Save(stock).Error
	// Alternative implementation if you need more control:
	// var existing domain.Stock
	// err := r.db.Where("item_id = ? AND warehouse_id = ?", stock.ItemID, stock.WarehouseID).First(&existing).Error
	// if err == gorm.ErrRecordNotFound {
	//     return r.db.Create(stock).Error
	// }
	// return r.db.Model(&existing).Updates(stock).Error
}

func (r *stockRepository) DeleteByItemAndWarehouse(itemID, warehouseID int) error {
	return r.db.Where("item_id = ? AND warehouse_id = ?", itemID, warehouseID).Delete(&domain.Stock{}).Error
}
