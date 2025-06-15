package usecase

import (
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/repository"
)

type WarehouseUsecase interface {
	GetAll() ([]domain.Warehouse, error)
	GetByID(id int) (*domain.Warehouse, error)
	Create(warehouse *domain.Warehouse) error
	Update(warehouse *domain.Warehouse) error
	Delete(id int) error
}

type warehouseUsecase struct {
	repo repository.WarehouseRepository
}

func NewWarehouseUsecase(r repository.WarehouseRepository) WarehouseUsecase {
	return &warehouseUsecase{repo: r}
}

func (u *warehouseUsecase) GetAll() ([]domain.Warehouse, error) {
	return u.repo.GetAll()
}

func (u *warehouseUsecase) GetByID(id int) (*domain.Warehouse, error) {
	return u.repo.GetByID(id)
}

func (u *warehouseUsecase) Create(w *domain.Warehouse) error {
	return u.repo.Create(w)
}

func (u *warehouseUsecase) Update(w *domain.Warehouse) error {
	return u.repo.Update(w)
}

func (u *warehouseUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}
