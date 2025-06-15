package persistence

import (
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/repository"
	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) repository.AdminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) GetByUsername(username string) (*domain.Admin, error) {
	var admin domain.Admin
	err := r.db.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) Create(admin *domain.Admin) error {
	return r.db.Create(admin).Error
}
