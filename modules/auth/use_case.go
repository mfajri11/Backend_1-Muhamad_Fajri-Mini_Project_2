package auth

import (
	"fmt"
	accountRepo "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/account"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type IAuthUseCase interface {
	Login(username, password string) (token string, exp time.Time, err error)
	ValidateToken(token string) (any, error)
}

//go:generate mockery --name ITokenManager
type ITokenManager interface {
	GenerateToken(username string, role string, exp time.Time) (string, error)
	ValidateToken(token string) (any, error)
}

type AuthUseCase struct {
	accountRepo  accountRepo.IAccountRepository
	tokenManager ITokenManager
}

func NewAuthUseCase(accountRepo accountRepo.IAccountRepository) *AuthUseCase {
	return &AuthUseCase{accountRepo: accountRepo}
}

func (uc *AuthUseCase) Login(username, password string) (token string, exp time.Time, err error) {
	acc, err := uc.accountRepo.FirstByUsername(username)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("modules.AuthUseCase.Login: error find account: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(acc.HashedPassword), []byte(password))
	if err != nil || acc.Username != username {
		return "", time.Time{}, fmt.Errorf("modules.AuthUseCase.Login: error authenticate: %w", err)
	}
	expireTime := time.Now()
	token, err = uc.tokenManager.GenerateToken(acc.Username, acc.Role.Name, expireTime)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("modules.AuthUseCase.Login: error generate token %w", err)
	}
	return token, expireTime.Add(48 * time.Hour), nil
}

func (uc AuthUseCase) ValidateToken(token string) (any, error) {
	return uc.tokenManager.ValidateToken(token)

}
