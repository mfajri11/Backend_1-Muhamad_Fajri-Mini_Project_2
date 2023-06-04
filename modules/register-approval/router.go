package register_approval

import (
	"github.com/gin-gonic/gin"
	registerApprovalRepo "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/register-approval"
	"gorm.io/gorm"
)

type RegisterApprovalRouter struct {
	registerApprovalHandler RegisterApprovalRequestHandler
}

func NewRegisterApprovalRouter(db *gorm.DB) *RegisterApprovalRouter {
	return &RegisterApprovalRouter{
		registerApprovalHandler: RegisterApprovalRequestHandler{
			registerApprovalController: &RegisterApprovalController{
				registerApprovalUseCase: &RegisterApprovalUseCase{
					registerApprovalRepo: registerApprovalRepo.NewRegisterApproval(db),
				},
			},
		}}
}

func (r RegisterApprovalRouter) Handle(router *gin.Engine) {
	basePath := "/approve"
	approve := router.Group(basePath)
	approve.GET("/", r.registerApprovalHandler.FindAll)
	approve.PUT("/:id", r.registerApprovalHandler.UpdateApprovalStatus)
}
