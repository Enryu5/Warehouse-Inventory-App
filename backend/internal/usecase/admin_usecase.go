package usecase

import (
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/domain"
	"github.com/Enryu5/Warehouse-Inventory-App/backend/internal/repository"
)

type AdminUsecase interface {
	GetByUsername(username string) (*domain.Admin, error)
	Create(user *domain.Admin) error
}

type adminUsecase struct {
	repo repository.AdminRepository
}

func NewAdminUsecase(r repository.AdminRepository) AdminUsecase {
	return &adminUsecase{repo: r}
}

func (u *adminUsecase) GetByUsername(username string) (*domain.Admin, error) {
	return u.repo.GetByUsername(username)
}

func (u *adminUsecase) Create(user *domain.Admin) error {
	return u.repo.Create(user)
}
