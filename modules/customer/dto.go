package customer

import "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/dto"

type CustomerResponse struct {
	dto.ResponseMeta
	Data any `json:"data,omitempty"`
}

type CustomerParams struct {
	FirstName string `json:"first_name" form:"name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" form:"email"`
	Avatar    string `json:"avatar"`
	Page      int    `form:"page"`
}
