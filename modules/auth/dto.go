package auth

import (
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/dto"
)

type LoginParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	AccessToken string `json:"access_token"`
	ExpiredAt   string `json:"expired_at"`
}

type LoginResponse struct {
	dto.ResponseMeta
	Data Login `json:"data"`
}
