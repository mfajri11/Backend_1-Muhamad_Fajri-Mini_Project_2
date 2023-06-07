package account

import "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/dto"

type AccountParams struct {
	UserName       string `json:"user_name" form:"username" binding:"required"`
	Password       string `json:"password,omitempty" binding:"required,min=8"`
	RoleName       string `json:"role_name" binding:"oneof=admin 'super admin'"`
	Page           int    `json:"-" form:"page" binding:"numeric"`
	ActivatedValue string `json:"activated_valuse,omitempty" binding:"oneof=true false True False TRUE FALSE"`
}

type AccountResponse struct {
	dto.ResponseMeta
	Data any `json:"data,omitempty"`
}
