package persistence

import (
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	firebaseSync "github.com/Enryu5/Warehouse-Inventory-App/backend/internal/infrastructure/firebase"
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
	if err := r.db.Save(stock).Error; err != nil {
		return err
	}
	return firebaseSync.SyncStockToFirebase(stock)
}

func (r *stockRepository) DeleteByItemAndWarehouse(itemID, warehouseID int) error {
	if err := r.db.Where("item_id = ? AND warehouse_id = ?", itemID, warehouseID).Delete(&domain.Stock{}).Error; err != nil {
		return err
	}
	return firebaseSync.DeleteStockFromFirebase(warehouseID, itemID)
}
