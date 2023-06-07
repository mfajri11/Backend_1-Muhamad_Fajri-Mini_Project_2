package customer

import (
	"fmt"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/dto"
)

type ICustomerController interface {
	Create(req CustomerParams) (CustomerResponse, error)
	Update(req CustomerUpdateParams) (CustomerResponse, error)
	Delete(id uint) error
	Search(page int, name string, email string) (CustomerResponse, error)
}

type CustomerController struct {
	customerUC ICustomerUseCase
}

func NewCustomerController(customerUC ICustomerUseCase) *CustomerController {
	return &CustomerController{customerUC: customerUC}
}

func (ctrl *CustomerController) Create(req CustomerParams) (CustomerResponse, error) {
	customer, err := ctrl.customerUC.Create(req)
	if err != nil {
		return CustomerResponse{}, fmt.Errorf("modules.CustomerController.Create: error create customer: %w", err)
	}
	res := CustomerResponse{
		dto.ResponseMeta{
			Success:      true,
			MessageTitle: "success crete account",
			Message:      "account already created but in pending status need approval",
			ResponseTime: "",
		},
		customer,
	}
	return res, nil
}

func (ctrl *CustomerController) Update(req CustomerUpdateParams) (CustomerResponse, error) {
	account, err := ctrl.customerUC.Update(req)
	if err != nil {
		return CustomerResponse{}, fmt.Errorf("modules.CustomerController.Update: error update customer: %w", err)
	}
	res := CustomerResponse{
		dto.ResponseMeta{
			Success:      true,
			MessageTitle: "success update account",
			Message:      "account already updated",
			ResponseTime: "",
		},
		account,
	}
	return res, nil
}

func (ctrl *CustomerController) Delete(id uint) error {
	err := ctrl.customerUC.Delete(id)
	if err != nil {
		return fmt.Errorf("modules.CustomerController.Delete: error delete customer: %w", err)
	}

	return nil
}

func (ctrl *CustomerController) Search(page int, name string, email string) (CustomerResponse, error) {
	customers, err := ctrl.customerUC.Search(page, name, email)
	if err != nil {
		return CustomerResponse{}, fmt.Errorf("modules.AccountController.Update: error update account: %w", err)
	}
	res := CustomerResponse{
		dto.ResponseMeta{
			Success:      true,
			MessageTitle: "success find account",
			Message:      "account already found",
			ResponseTime: "",
		},
		customers,
	}
	return res, nil
}
