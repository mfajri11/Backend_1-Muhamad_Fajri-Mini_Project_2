package register_approval

import "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/dto"

type RegisterApprovalResponse struct {
	dto.ResponseMeta
	Data any `json:"data,omitempty"`
}

type RegisterApprovalParams struct {
	ID           uint   `json:"id"`
	AdminID      uint   `json:"admin_id"`
	SuperAdminID uint   `json:"super_admin_id"`
	Status       string `json:"status"`
	Page         int    `form:"page"`
}
