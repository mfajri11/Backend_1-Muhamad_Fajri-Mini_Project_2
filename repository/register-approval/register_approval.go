package register_approval

import (
	"fmt"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/entity"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository"
	"gorm.io/gorm"
)

type IRegisterApprovalRepository interface {
	FindAll(page int) ([]*entity.RegisterApproval, error)
	UpdateApprovalStatus(id uint, val string) error
}

type RegisterApprovalRepository struct {
	db *gorm.DB
}

func NewRegisterApproval(db *gorm.DB) *RegisterApprovalRepository {
	return &RegisterApprovalRepository{db: db}
}

func (repo *RegisterApprovalRepository) FindAll(page int) ([]*entity.RegisterApproval, error) {
	approvs := make([]*entity.RegisterApproval, 0)
	err := repo.db.Scopes(repository.Paginate(page)).Find(&approvs).Error
	if err != nil {
		return nil, fmt.Errorf("repository.RegsiterApproval.FindAll: %w", err)
	}

	return approvs, nil
}

func (repo *RegisterApprovalRepository) UpdateApprovalStatus(id uint, val string) error {
	var err error
	//  if id == 0 update all pending admin
	if id == uint(0) {
		err = repo.db.Model(&entity.RegisterApproval{}).Where("status = ?", "pending").Update("status", val).Error
		if err != nil {
			return fmt.Errorf("repository.RegsiterApproval.UpdateApprovalStatus: %w", err)
		}
	}

	err = repo.db.Model(&entity.RegisterApproval{}).Where("id = ?", id).Update("status", val).Error
	if err != nil {
		return fmt.Errorf("repository.RegsiterApproval.UpdateApprovalStatus: %w", err)
	}
	return err
}
