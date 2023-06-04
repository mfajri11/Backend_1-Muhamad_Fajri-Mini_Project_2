package register_approval

import (
	"fmt"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/entity"
	registerApprovalRepo "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/register-approval"
)

type IRegisterApprovalUseCase interface {
	FindAll(page int) ([]*entity.RegisterApproval, error)
	UpdateApprovalStatus(id uint, val string) error
}

type RegisterApprovalUseCase struct {
	registerApprovalRepo registerApprovalRepo.IRegisterApprovalRepository
}

func NewRegisterApprovalUseCase(registerApprovalRepo registerApprovalRepo.IRegisterApprovalRepository) *RegisterApprovalUseCase {
	return &RegisterApprovalUseCase{registerApprovalRepo: registerApprovalRepo}
}

func (uc RegisterApprovalUseCase) FindAll(page int) ([]*entity.RegisterApproval, error) {
	approvs, err := uc.registerApprovalRepo.FindAll(page)
	if err != nil {
		return nil, fmt.Errorf("modules.RegisterApprovalUseCase.FindAll: error find approval: %w", err)
	}

	return approvs, nil
}

func (uc RegisterApprovalUseCase) UpdateApprovalStatus(id uint, val string) error {
	err := uc.registerApprovalRepo.UpdateApprovalStatus(id, val)
	if err != nil {
		return fmt.Errorf("modules.RegisterApprovalUseCase.UpdateApprovalStatus: error update approval: %w", err)
	}

	return nil
}
