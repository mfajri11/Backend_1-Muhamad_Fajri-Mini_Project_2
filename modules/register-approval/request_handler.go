package register_approval

import (
	"github.com/gin-gonic/gin"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/dto"
	register_approval "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/register-approval"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RegisterApprovalRequestHandler struct {
	registerApprovalController IRegisterApprovalController
}

func NewRegisterApprovalRequestHandler(db *gorm.DB) *RegisterApprovalRequestHandler {
	return &RegisterApprovalRequestHandler{
		registerApprovalController: RegisterApprovalController{
			registerApprovalUseCase: RegisterApprovalUseCase{
				registerApprovalRepo: register_approval.NewRegisterApproval(db),
			},
		},
	}
}

func (h RegisterApprovalRequestHandler) FindAll(c *gin.Context) {
	approvalQuery := RegisterApprovalParams{}
	err := c.ShouldBindQuery(&approvalQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	resp, err := h.registerApprovalController.FindAll(approvalQuery.Page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h RegisterApprovalRequestHandler) UpdateApprovalStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	approvalQuery := RegisterApprovalParams{}
	err = c.ShouldBindJSON(&approvalQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	err = h.registerApprovalController.UpdateApprovalStatus(uint(id), approvalQuery.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponseWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, nil)

}
