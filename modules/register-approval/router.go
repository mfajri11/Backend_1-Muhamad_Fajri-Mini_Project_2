package register_approval

import "github.com/gin-gonic/gin"

type RegisterApprovalRouter struct {
	registerApprovalHandler RegisterApprovalRequestHandler
}

func NewRegisterApprovalRouter(registerApprovalHandler RegisterApprovalRequestHandler) *RegisterApprovalRouter {
	return &RegisterApprovalRouter{registerApprovalHandler: registerApprovalHandler}
}

func (r RegisterApprovalRouter) Handle(router *gin.Engine) {
	basePath := "/approve"
	approve := router.Group(basePath)
	approve.GET("/", r.registerApprovalHandler.FindAll)
	approve.PUT("/:id", r.registerApprovalHandler.UpdateApprovalStatus)
}
