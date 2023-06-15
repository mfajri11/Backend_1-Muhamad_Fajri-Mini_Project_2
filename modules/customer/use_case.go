package customer

import (
	"fmt"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/entity"
	customerRepository "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/customer"
)

type ICustomerUseCase interface {
	Create(params CustomerParams) (entity.Customer, error)
	Update(params CustomerUpdateParams) (*entity.Customer, error)
	Delete(id uint) error
	Search(page int, name string, email string) ([]*entity.Customer, error)
}

type CustomerUseCase struct {
	customerRepo customerRepository.ICustomerRepository
}

func NewCustomerUseCase(customerRepo customerRepository.ICustomerRepository) *CustomerUseCase {
	return &CustomerUseCase{customerRepo: customerRepo}
}

func (uc *CustomerUseCase) Create(customer CustomerParams) (entity.Customer, error) {
	newCustomer := entity.Customer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
	}

	err := uc.customerRepo.Create(&newCustomer)
	if err != nil {
		return entity.Customer{}, fmt.Errorf("modules.CustomerUseCase.Create: error create customer %w", err)
	}

	return newCustomer, nil
}

func (uc *CustomerUseCase) Update(customer CustomerUpdateParams) (*entity.Customer, error) {
	newCustomer := entity.Customer{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
	}

	customerUpdated, err := uc.customerRepo.Update(&newCustomer)
	if err != nil {
		return &entity.Customer{}, fmt.Errorf("modules.CustomerUseCase.Update: error update customer %w", err)
	}

	return customerUpdated, nil
}

func (uc *CustomerUseCase) Delete(id uint) error {
	err := uc.customerRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("modules.CustomerUseCase.Create: error create customer %w", err)
	}

	return nil
}

func (uc *CustomerUseCase) Search(page int, name string, email string) ([]*entity.Customer, error) {
	customers, err := uc.customerRepo.Search(page, name, email)
	if err != nil {
		return nil, fmt.Errorf("modules.CustomerUseCase.Create: error create customer %w", err)
	}

	return customers, nil
}
