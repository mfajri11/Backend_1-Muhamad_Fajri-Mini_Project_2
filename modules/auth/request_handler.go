package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/dto"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/modules/auth"
	accountRepo "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/account"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type AuthRequestHandler struct {
	authCtrl     IAuthController
	tokenManager auth.ITokenManager
}

func NewAuthRequestHandler(db *gorm.DB) *AuthRequestHandler {
	return &AuthRequestHandler{
		authCtrl: AuthController{
			AuthUC: AuthUseCase{
				accountRepo: accountRepo.NewUserRepository(db),
			},
		},
	}
}

func (h AuthRequestHandler) Login(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	req := LoginParams{
		username,
		password,
	}
	resp, err := h.authCtrl.Login(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h AuthRequestHandler) AuthorizationRequired(c *gin.Context) {
	bearerSchema := "bearer "
	tokenStr := c.Request.Header.Get("Authorization")
	if !strings.HasPrefix(tokenStr, bearerSchema) {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	token := tokenStr[len(bearerSchema):]
	payload, err := h.tokenManager.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}
	c.Set("Authorization", payload)
	c.Next()
}
