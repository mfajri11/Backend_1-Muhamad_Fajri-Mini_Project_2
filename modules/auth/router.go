package auth

import "github.com/gin-gonic/gin"

type AuthRouter struct {
	authHandler AuthRequestHandler
}

func NewAuthRouter(authHandler AuthRequestHandler) *AuthRouter {
	return &AuthRouter{authHandler: authHandler}
}

func (r AuthRouter) Handle(router *gin.Engine) {
	basePath := "/auth"
	auth := router.Group(basePath)
	auth.POST("/login", r.authHandler.Login)
}
