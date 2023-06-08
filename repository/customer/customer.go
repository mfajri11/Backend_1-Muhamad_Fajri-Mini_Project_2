package customer

import (
	"fmt"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/entity"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository"
	"gorm.io/gorm"
)

//go:generate mockery --name ICustomerRepository
type ICustomerRepository interface {
	Create(customer *entity.Customer) error
	Update(customer *entity.Customer) (*entity.Customer, error)
	Delete(id uint) error
	Search(page int, name string, email string) ([]*entity.Customer, error)
}

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (repo *CustomerRepository) Create(customer *entity.Customer) error {
	err := repo.db.Create(&customer).Error
	if err != nil {
		return fmt.Errorf("repostiory.CustomerRepository.Create: error create customer: %w", err)
	}
	return nil
}

func (repo *CustomerRepository) Update(customer *entity.Customer) (*entity.Customer, error) {
	err := repo.db.Model(&entity.Customer{}).Where("id = ? ", customer.ID).Updates(customer).Error
	if err != nil {
		return nil, fmt.Errorf("repostiory.CustomerRepository.Update: error update customer: %w", err)
	}
	err = repo.db.Find(&customer).Error
	if err != nil {
		return nil, fmt.Errorf("repostiory.CustomerRepository.Update: error find customer: %w", err)
	}
	return customer, nil
}

func (repo *CustomerRepository) Delete(id uint) error {
	customer := entity.Customer{}
	err := repo.db.First(&customer, "id = ?", id).Error
	if err != nil {
		return fmt.Errorf("repostiory.CustomerRepository.Delete: error find customer: %w", err)
	}
	err = repo.db.Delete(&customer).Error
	if err != nil {
		return fmt.Errorf("repostiory.CustomerRepository.Delete: error delete customer: %w", err)
	}
	return nil
}

func (repo *CustomerRepository) Search(page int, name string, email string) ([]*entity.Customer, error) {
	customers := make([]*entity.Customer, 0)
	cust := entity.Customer{Email: email, FirstName: name}
	db := repo.db.Scopes(repository.Paginate(page)).Where(cust)
	err := db.Find(&customers).Error
	if err != nil {
		return nil, fmt.Errorf("repostiory.CustomerRepository.Search: error find account: %w", err)
	}

	return customers, nil
}
