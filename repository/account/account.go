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
	UpdateActivatedAccount(id uint, activated bool) error
	FirstByUsername(username string) (*entity.Account, error)
}

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) IAccountRepository {
	return &AccountRepository{db: db}
}

func (repo *AccountRepository) Create(account entity.Account) error {
	err := repo.db.Create(&account).Error
	if err != nil {
		return fmt.Errorf("repostiory.AccountRepository.Create: error create account: %w", err)
	}
	account.RegisterApproval.AdminID = account.ID
	err = repo.db.Model(&entity.RegisterApproval{}).Where("id = ?", account.RegisterApprovalID).Update("admin_id", account.ID).Error
	if err != nil {
		return fmt.Errorf("repostiory.AccountRepository.Create: error update `register_approval`.`admin_id`: %w", err)
	}
	//err = repo.db.Create(&entity.RegisterApproval{
	//	AdminID:      account.ID,
	//	SuperAdminID: 1,
	//	Status:       "pending",
	//}).Error
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
	err := repo.db.Model(&entity.Account{}).Where("id = ? ", account.ID).Updates(&account).Error
	if err != nil {
		return nil, fmt.Errorf("repostiory.AccountRepository.Update: error update account: %w", err)
	}
	err = repo.db.Preload("Role").First(&account, account.ID).Error
	if err != nil {
		return nil, fmt.Errorf("repostiory.AccountRepository.Update: error find account: %w", err)
	}
	return account, nil
}

//func (repo *AccountRepository) Search(page int, email string) ([]*entity.Account, error) {
//	accounts := make([]*entity.Account, 0)
//
//	err := repo.db.Scopes(repository.Paginate(page)).Where("email = ? ", email).Find(&accounts).Error
//	if err != nil {
//		return nil, fmt.Errorf("repostiory.AccountRepository.Search: error find account: %w", err)
//	}
//	return accounts, nil
//}

func (repo *AccountRepository) FindByUsername(page int, username string) (*entity.Account, error) {
	account := entity.Account{}
	err := repo.db.Preload("Role").Scopes(repository.Paginate(page)).Where("username = ? ", username).Find(&account).Error
	if err != nil {
		return nil, fmt.Errorf("repostiory.AccountRepository.FindByUsername: error find account: %w", err)
	}
	return &account, nil
}

func (repo *AccountRepository) FirstByUsername(username string) (*entity.Account, error) {
	account := entity.Account{}
	err := repo.db.Preload("Role").First(&account, "username = ?", username).Error
	if err != nil {
		return nil, fmt.Errorf("repostiory.AccountRepository.FirstByUsername: error find account: %w", err)
	}
	return &account, nil
}

func (repo *AccountRepository) UpdateActivatedAccount(id uint, activated bool) error {
	err := repo.db.
		Model(&entity.Account{}).
		Where("id = ?", id).
		Update("activated", activated).
		Error
	if err != nil {
		return fmt.Errorf("repostiory.AccountRepository.UpdateActivateAccount: error update account: %w", err)

	}
	return nil
}
