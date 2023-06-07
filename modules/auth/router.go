package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthRouter struct {
	authHandler *AuthRequestHandler
}

func NewAuthRouter(db *gorm.DB) *AuthRouter {
	return &AuthRouter{
		authHandler: NewAuthRequestHandler(db),
	}
}

func (r AuthRouter) Handle(router *gin.Engine) {
	basePath := "/auth"
	auth := router.Group(basePath)
	auth.POST("/login", r.authHandler.Login)
}
