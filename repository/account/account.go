package account

import (
	"fmt"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/entity"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	Create(account entity.Account) error
	Delete(id uint) error
	Update(account *entity.Account) (*entity.Account, error)
	//Search(page int, username string) ([]*entity.Account, error)
	FindByUsername(page int, username string) (*entity.Account, error)
	UpdateActivateAccount(id uint, activateValue string) error
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

func (repo *AccountRepository) Search(page int, email string) ([]*entity.Account, error) {
	accounts := make([]*entity.Account, 0)

	err := repo.db.Scopes(repository.Paginate(page)).Where("email = ? ", email).Find(&accounts).Error
	if err != nil {
		return nil, fmt.Errorf("repostiory.AccountRepository.Search: error find account: %w", err)
	}
	return accounts, nil
}

func (repo *AccountRepository) FindByUsername(page int, username string) (*entity.Account, error) {
	account := entity.Account{}
	err := repo.db.Scopes(repository.Paginate(page)).Where("username = ? ", username).Find(&account).Error
	if err != nil {
		return nil, fmt.Errorf("repostiory.AccountRepository.FindByUsername: error find account: %w", err)
	}
	return &account, nil
}

func (repo *AccountRepository) UpdateActivateAccount(id uint, activateValue string) error {
	err := repo.db.
		Model(&entity.Account{}).
		Where("id = ?", id).
		Update("activated", activateValue).
		Error
	if err != nil {
		return fmt.Errorf("repostiory.AccountRepository.UpdateActivateAccount: error update account: %w", err)

	}
	return nil
}
