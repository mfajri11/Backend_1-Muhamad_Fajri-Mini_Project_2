package account

import (
	"context"
	"fmt"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/dto"
)

type IAccountController interface {
	Create(ctx context.Context, req AccountParams) (AccountResponse, error)
	Update(ctx context.Context, req AccountUpdateParams) (AccountResponse, error)
	Delete(ctx context.Context, id uint) error
	FindByUsername(page int, username string) (AccountResponse, error)
	UpdateActivatedAccount(ctx context.Context, id uint, activated bool) error
}

type AccountController struct {
	AccountUC IAccountUseCase
}

func NewAccountController(accountUC IAccountUseCase) *AccountController {
	return &AccountController{AccountUC: accountUC}
}

func (ctrl *AccountController) Create(ctx context.Context, req AccountParams) (AccountResponse, error) {
	account, err := ctrl.AccountUC.Create(ctx, req)
	if err != nil {
		return AccountResponse{}, fmt.Errorf("modules.AccountController.Create: error create account: %w", err)
	}
	res := AccountResponse{
		dto.ResponseMeta{
			Success:      true,
			MessageTitle: "success crete account",
			Message:      "account already created but in pending status need approval",
			ResponseTime: "",
		},
		AccountParams{
			UserName: account.Username,
			RoleName: account.Role.Name,
		},
	}
	return res, nil
}

func (ctrl *AccountController) Update(ctx context.Context, req AccountUpdateParams) (AccountResponse, error) {
	account, err := ctrl.AccountUC.Update(ctx, req)
	if err != nil {
		return AccountResponse{}, fmt.Errorf("modules.AccountController.Update: error update account: %w", err)
	}
	res := AccountResponse{
		dto.ResponseMeta{
			Success:      true,
			MessageTitle: "success update account",
			Message:      "account already updated",
			ResponseTime: "",
		},
		AccountParams{
			UserName: account.Username,
			RoleName: account.Role.Name,
		},
	}
	return res, nil
}

func (ctrl *AccountController) Delete(ctx context.Context, id uint) error {
	err := ctrl.AccountUC.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("modules.AccountController.Delete: error delete account: %w", err)
	}

	return nil
}

func (ctrl *AccountController) UpdateActivatedAccount(ctx context.Context, id uint, activated bool) error {
	err := ctrl.AccountUC.UpdateActivatedAccount(ctx, id, activated)
	if err != nil {
		return fmt.Errorf("modules.AccountController.UpdateActivatedAccount: error update account: %w", err)
	}

	return nil
}

func (ctrl *AccountController) FindByUsername(page int, username string) (AccountResponse, error) {
	account, err := ctrl.AccountUC.FindByUsername(page, username)
	if err != nil {
		return AccountResponse{}, fmt.Errorf("modules.AccountController.Update: error update account: %w", err)
	}
	res := AccountResponse{
		dto.ResponseMeta{
			Success:      true,
			MessageTitle: "success find account",
			Message:      "account already found",
			ResponseTime: "",
		},
		AccountParams{
			UserName: account.Username,
			RoleName: account.Role.Name,
		},
	}
	return res, nil
}
