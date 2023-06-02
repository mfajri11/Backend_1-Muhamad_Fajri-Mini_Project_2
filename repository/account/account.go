package repository

import (
	"fmt"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/entity"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	Create(account entity.Account) error
	Delete(id uint) error
	Update(account *entity.Account) (*entity.Account, error)
}

type AccountRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IAccountRepository {
	return &AccountRepository{db: db}
}

func (repo *AccountRepository) Create(account entity.Account) error {
	err := repo.db.Create(&account).Error
	if err != nil {
		return fmt.Errorf("repostiory.AccountRepository.Create: error create account: %w", err)
	}
	return nil
}

func (repo *AccountRepository) Delete(id uint) error {
	account := entity.Account{}
	err := repo.db.First(&account, "id = ?", id).Error
	if err != nil {
		return fmt.Errorf("repostiory.AccountRepository.Delete: error find account: %w", err)
	}
	err = repo.db.Delete(&account).Error
	if err != nil {
		return fmt.Errorf("repostiory.AccountRepository.Delete: error delete account: %w", err)
	}
	return nil
}

func (repo *AccountRepository) Update(account *entity.Account) (*entity.Account, error) {
	err := repo.db.Model(&entity.Account{}).Where("id = ? ", account.ID).Updates(account).Error
	if err != nil {
		return nil, fmt.Errorf("repostiory.AccountRepository.Update: error update account: %w", err)
	}
	err = repo.db.Find(&account).Error
	if err != nil {
		return nil, fmt.Errorf("repostiory.AccountRepository.Update: error find account: %w", err)
	}
	return account, nil
}

func (repo *AccountRepository) Search(page int, username string, email string) ([]*entity.Account, error) {
	whereMap := make(map[string]string)
	accounts := make([]*entity.Account, 0)
	if len(username) > 0 {
		whereMap["username"] = username
	}
	if len(email) > 0 {
		whereMap["email"] = email
	}

	repo.db.Scopes(paginate(page)).Where(whereMap).Find(&accounts)
	return accounts, nil
}
