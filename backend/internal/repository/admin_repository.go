package repository

import "github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"

type AdminRepository interface {
	GetByUsername(username string) (*domain.Admin, error)
	Create(user *domain.Admin) error
}
