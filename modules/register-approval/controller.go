package register_approval

import (
	"fmt"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/dto"
)

type IRegisterApprovalController interface {
	FindAll(page int) (RegisterApprovalResponse, error)
	UpdateApprovalStatus(id uint, val string) error
}

type RegisterApprovalController struct {
	registerApprovalUseCase IRegisterApprovalUseCase
}

func NewRegisterApprovalController(registerApprovalUseCase IRegisterApprovalUseCase) *RegisterApprovalController {
	return &RegisterApprovalController{registerApprovalUseCase: registerApprovalUseCase}
}

func (ctrl *RegisterApprovalController) FindAll(page int) (RegisterApprovalResponse, error) {
	approvs, err := ctrl.registerApprovalUseCase.FindAll(page)
	if err != nil {
		return RegisterApprovalResponse{}, fmt.Errorf("modules.RegisterApprovalController.FindAll: error find approval: %w", err)
	}

	res := RegisterApprovalResponse{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "success find approvals",
			Message:      "approvals already found",
			ResponseTime: "",
		},
	}
	apvs := make([]RegisterApprovalParams, 0)
	for _, apv := range approvs {
		apvs = append(apvs, RegisterApprovalParams{
			ID:           apv.ID,
			AdminID:      apv.AdminID,
			SuperAdminID: apv.SuperAdminID,
			Status:       apv.Status,
		})
	}
	res.Data = apvs

	return res, nil

}

func (ctrl *RegisterApprovalController) UpdateApprovalStatus(id uint, val string) error {
	err := ctrl.registerApprovalUseCase.UpdateApprovalStatus(id, val)
	if err != nil {
		return fmt.Errorf("modules.RegisterApprovalController.UpdateApprovalStatus: error update approval: %w", err)
	}

	return nil
}
