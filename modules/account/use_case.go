package account

import (
	"context"
	"fmt"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/entity"
	accountRepo "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/account"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/utils/security"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type IAccountUseCase interface {
	Create(ctx context.Context, accountParams AccountParams) (entity.Account, error)
	Update(ctx context.Context, accountParams AccountParams) (*entity.Account, error)
	Delete(ctx context.Context, id uint) error
	FindByUsername(page int, username string) (*entity.Account, error)
	UpdateActivatedAccount(ctx context.Context, id uint, activatedValue string) error
}

type AccountUseCase struct {
	accountRepo accountRepo.IAccountRepository
}

func NewAccountUseCase(accountRepo accountRepo.IAccountRepository) *AccountUseCase {
	return &AccountUseCase{accountRepo: accountRepo}
}

func (uc *AccountUseCase) Create(ctx context.Context, accountParams AccountParams) (entity.Account, error) {
	if !uc.isSuperAdmin(ctx) {
		return entity.Account{}, fmt.Errorf("modules.AccountUseCase.Create: error unauthorized")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(accountParams.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.Account{}, fmt.Errorf("modules.AccountUseCase.Create: error hash password %w", err)
	}
	newAccount := entity.Account{
		Username:       accountParams.UserName,
		HashedPassword: string(hashedPassword),
		Role:           entity.Role{Name: accountParams.RoleName},
		Verified:       false,
		Activated:      false,
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
	//TODO: create side effect populate newAccount with role_id and approval_id
	err = uc.accountRepo.Create(newAccount)
	if err != nil {
		return entity.Account{}, fmt.Errorf("modules.AccountUseCase.Create: error create account %w", err)
	}

	return newAccount, nil
}

func (uc *AccountUseCase) Update(ctx context.Context, accountParams AccountParams) (*entity.Account, error) {
	if !uc.isSuperAdmin(ctx) {
		return &entity.Account{}, fmt.Errorf("modules.AccountUseCase.Create: error unauthorized")
	}
	newAccount := entity.Account{
		Username:  accountParams.UserName,
		Role:      entity.Role{Name: accountParams.RoleName},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	acc, err := uc.accountRepo.Update(&newAccount)
	if err != nil {
		return &entity.Account{}, fmt.Errorf("modules.AccountUseCase.Update: error update account %w", err)
	}

	return acc, nil
}

func (uc *AccountUseCase) FindByUsername(page int, username string) (*entity.Account, error) {
	acc, err := uc.accountRepo.FindByUsername(page, username)
	if err != nil {
		return nil, fmt.Errorf("modules.AccountUseCase.FindByUsername: error find account %w", err)
	}

	return acc, nil
}

func (uc *AccountUseCase) UpdateActivatedAccount(ctx context.Context, id uint, activatedValue string) error {
	if !uc.isSuperAdmin(ctx) {
		return fmt.Errorf("modules.AccountUseCase.Create: error unauthorized")
	}
	err := uc.accountRepo.UpdateActivatedAccount(id, activatedValue)
	if err != nil {
		return fmt.Errorf("modules.AccountUseCase.UpdateActivatedAccount: error update account %w", err)
	}

	return nil
}

func (_ *AccountUseCase) isSuperAdmin(ctx context.Context) bool {
	payload := ctx.Value("Authorization")
	if payload == nil {
		return false
	}
	acc := payload.(security.JWTClaims)
	return acc.Role == "super admin"
}

func (uc *AccountUseCase) Delete(ctx context.Context, id uint) error {
	if !uc.isSuperAdmin(ctx) {
		return fmt.Errorf("modules.AccountUseCase.Delete: error delete account")
	}
	err := uc.accountRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("modules.AccountUseCase.Delete: error delete account %w", err)
	}

	return nil
}
