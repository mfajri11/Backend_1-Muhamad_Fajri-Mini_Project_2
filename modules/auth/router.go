package auth

import (
	"github.com/gin-gonic/gin"
	accountRepo "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/account"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/utils/security"
	"gorm.io/gorm"
)

type AuthRouter struct {
	authHandler AuthRequestHandler
}

func NewAuthRouter(db *gorm.DB) *AuthRouter {
	return &AuthRouter{
		authHandler: AuthRequestHandler{
			authCtrl: &AuthController{
				AuthUC: &AuthUseCase{
					accountRepo:  accountRepo.NewUserRepository(db),
					tokenManager: security.NewTokenManager("secret"),
				},
			},
		}}
}

func (r AuthRouter) Handle(router *gin.Engine) {
	basePath := "/auth"
	auth := router.Group(basePath)
	auth.POST("/login", r.authHandler.Login)
}
