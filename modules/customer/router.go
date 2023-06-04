package customer

import "github.com/gin-gonic/gin"

type CustomerRouter struct {
	customerHandler CustomerRequestHandler
}

func NewCustomerRouter(customerHandler CustomerRequestHandler) *CustomerRouter {
	return &CustomerRouter{customerHandler: customerHandler}
}

func (r CustomerRouter) Handle(e *gin.Engine) {
	basePath := "/customer"
	account := e.Group(basePath)
	account.POST("/", r.customerHandler.Create)
	account.PUT("/:id", r.customerHandler.Update)
	account.DELETE("/:id", r.customerHandler.Delete)
	account.GET("/search", r.customerHandler.Search)
}
