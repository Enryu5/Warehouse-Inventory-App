package usecase

import (
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/repository"
)

type ItemUsecase interface {
	GetAll() ([]domain.Item, error)
	GetByID(id int) (*domain.Item, error)
	Create(item *domain.Item) error
	Update(item *domain.Item) error
	Delete(id int) error
}

type itemUsecase struct {
	repo repository.ItemRepository
}

func NewItemUsecase(r repository.ItemRepository) ItemUsecase {
	return &itemUsecase{repo: r}
}

func (u *itemUsecase) GetAll() ([]domain.Item, error) {
	return u.repo.GetAll()
}

func (u *itemUsecase) GetByID(id int) (*domain.Item, error) {
	return u.repo.GetByID(id)
}

func (u *itemUsecase) Create(item *domain.Item) error {
	return u.repo.Create(item)
}

func (u *itemUsecase) Update(item *domain.Item) error {
	return u.repo.Update(item)
}

func (u *itemUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}
