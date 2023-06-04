package account

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountRouter struct {
	Accounthandler *AccountRequestHandler
}

func NewAccountRouter(db *gorm.DB) *AccountRouter {
	return &AccountRouter{Accounthandler: NewAccountRequestHandler(db)}
}

func (r *AccountRouter) Handle(e *gin.Engine) {
	basePath := "/account"
	account := e.Group(basePath)
	account.POST("/", r.Accounthandler.Create)
	account.PUT("/:id", r.Accounthandler.Update)
	account.DELETE("/:id", r.Accounthandler.Delete)
	account.GET("/search", r.Accounthandler.FindByUsername)
	account.PUT("/:id/activated/", r.Accounthandler.UpdateActivatedAccount)
}
