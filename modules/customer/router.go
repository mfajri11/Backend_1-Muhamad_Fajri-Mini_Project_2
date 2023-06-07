package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/customer"
	"gorm.io/gorm"
)

type CustomerRouter struct {
	customerHandler CustomerRequestHandler
}

func NewCustomerRouter(db *gorm.DB) *CustomerRouter {
	return &CustomerRouter{
		customerHandler: CustomerRequestHandler{
			customerController: &CustomerController{
				customerUC: &CustomerUseCase{
					customerRepo: customer.NewCustomerRepository(db),
				},
			},
		}}
}

func (r CustomerRouter) Handle(e *gin.Engine) {
	basePath := "/customer"
	account := e.Group(basePath)
	account.POST("/", r.customerHandler.Create)
	account.PATCH("/:id", r.customerHandler.Update)
	account.DELETE("/:id", r.customerHandler.Delete)
	account.GET("/search", r.customerHandler.Search)
}
