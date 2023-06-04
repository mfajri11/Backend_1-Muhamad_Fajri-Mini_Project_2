package account

import "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/dto"

type AccountParams struct {
	UserName       string `json:"user_name" form:"username"`
	Password       string `json:"password,omitempty"`
	RoleName       string `json:"role_name"`
	Page           int    `form:"page"`
	ActivatedValue string `json:"activated_value"`
}

type AccountResponse struct {
	dto.ResponseMeta
	Data any `json:"data,omitempty"`
}
