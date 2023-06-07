package account

import "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/dto"

type AccountParams struct {
	UserName  string `json:"user_name" binding:"required"`
	Password  string `json:"password,omitempty" binding:"required,min=8"`
	RoleName  string `json:"role_name" binding:"oneof=admin 'super admin'"`
	Page      int    `json:"-" form:"page" binding:"numeric"`
	Activated bool   `json:"activated,omitempty"`
	Verified  bool   `json:"verified,omitempty"`
}

type AccountUpdateParams struct {
	ID        uint
	UserName  string `json:"user_name" form:"username"`
	Password  string `json:"password,omitempty" binding:"omitempty,min=8"`
	RoleName  string `json:"role_name" binding:"omitempty,oneof=admin 'super admin'"`
	Page      int    `json:"-" form:"page" binding:"omitempty,numeric"`
	Activated bool   `json:"activated,omitempty"`
	Verified  bool   `json:"verified,omitempty"`
}

type AccountResponse struct {
	dto.ResponseMeta
	Data any `json:"data,omitempty"`
}
