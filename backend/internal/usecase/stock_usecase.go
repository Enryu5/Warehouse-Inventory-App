package usecase

import (
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/repository"
)

type StockUsecase interface {
	GetByWarehouse(warehouseID int) ([]domain.Stock, error)
	GetByItem(itemID int) ([]domain.Stock, error)
	Upsert(stock *domain.Stock) error
	DeleteByItemAndWarehouse(itemID, warehouseID int) error
}

type stockUsecase struct {
	repo repository.StockRepository
}

func NewStockUsecase(r repository.StockRepository) StockUsecase {
	return &stockUsecase{repo: r}
}

func (u *stockUsecase) GetByWarehouse(warehouseID int) ([]domain.Stock, error) {
	return u.repo.GetByWarehouse(warehouseID)
}

func (u *stockUsecase) GetByItem(itemID int) ([]domain.Stock, error) {
	return u.repo.GetByItem(itemID)
}

func (u *stockUsecase) Upsert(stock *domain.Stock) error {
	return u.repo.Upsert(stock)
}

func (u *stockUsecase) DeleteByItemAndWarehouse(itemID, warehouseID int) error {
	return u.repo.DeleteByItemAndWarehouse(itemID, warehouseID)
}
