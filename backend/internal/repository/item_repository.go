package repository

import "github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"

type ItemRepository interface {
	GetAll() ([]domain.Item, error)
	GetByID(id int) (*domain.Item, error)
	Create(item *domain.Item) error
	Update(item *domain.Item) error
	Delete(id int) error
}
