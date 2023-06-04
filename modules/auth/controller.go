package auth

import (
	"fmt"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/dto"
)

type IAuthController interface {
	Login(req LoginParams) (LoginResponse, error)
}

type AuthController struct {
	AuthUC IAuthUseCase
}

func NewAuthController(authUC IAuthUseCase) *AuthController {
	return &AuthController{AuthUC: authUC}
}

func (ctrl *AuthController) Login(req LoginParams) (LoginResponse, error) {
	token, exp, err := ctrl.AuthUC.Login(req.Username, req.Password)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("modules.LoginResponse.Login: error login: %w", err)
	}

	resp := LoginResponse{
		dto.ResponseMeta{
			Success:      true,
			MessageTitle: "success login account",
			Message:      "account already logged in",
			ResponseTime: "",
		},
		Login{
			AccessToken: token,
			ExpiredAt:   exp.String(),
		},
	}

	return resp, nil
}
