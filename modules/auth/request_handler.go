package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/dto"
	account2 "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/account"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/utils/security"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

type AuthRequestHandler struct {
	authCtrl IAuthController
}

func NewAuthRequestHandler(db *gorm.DB) *AuthRequestHandler {
	return &AuthRequestHandler{
		authCtrl: &AuthController{
			AuthUC: &AuthUseCase{
				accountRepo:  account2.NewAccountRepository(db),
				tokenManager: security.NewTokenManager("secret")}}}
}

func (h *AuthRequestHandler) Login(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		log.Println("modules.AuthRequestHandler.Login: invalid basic token")
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	req := LoginParams{
		username,
		password,
	}
	resp, err := h.authCtrl.Login(req)
	if err != nil {
		log.Printf("modules.AuthRequestHandler.Login: error login: %w", err)
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *AuthRequestHandler) AuthorizationRequired(c *gin.Context) {
	bearerSchema := "Bearer "
	tokenStr := c.Request.Header.Get("Authorization")
	if !strings.HasPrefix(tokenStr, bearerSchema) {
		log.Println("modules.AuthRequestHandler.AuthorizationRequired: invalid Bearer schema in header")
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	token := tokenStr[len(bearerSchema):]
	payload, err := h.authCtrl.ValidateToken(token)
	if err != nil {
		log.Println("modules.AuthRequestHandler.AuthorizationRequired: error validate token: %w", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}
	c.Set("Authorization", payload)
	c.Next()
}

func (r AuthRouter) UseAuthorizationRequired(router *gin.Engine) *gin.Engine {
	router.Use(r.authHandler.AuthorizationRequired)
	return router
}
