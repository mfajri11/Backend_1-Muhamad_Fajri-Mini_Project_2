package account

import (
	"github.com/gin-gonic/gin"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/dto"
	accountRepo "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/account"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type AccountRequestHandler struct {
	accountController IAccountController
}

func NewAccountRequestHandler(db *gorm.DB) *AccountRequestHandler {
	return &AccountRequestHandler{
		accountController: &AccountController{
			AccountUC: &AccountUseCase{
				accountRepo: accountRepo.NewAccountRepository(db),
			},
		},
	}
}

func (h *AccountRequestHandler) Create(c *gin.Context) {
	req := AccountParams{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("modules.AccountRequestHandler.Create: error bind json: %s", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	resp, err := h.accountController.Create(c, req)
	if err != nil {
		log.Printf("modules.AccountRequestHandler.Create: error create account: %s", err)
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *AccountRequestHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		log.Printf("modules.AccountRequestHandler.Update: error parsing path param: %s", err)
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	req := AccountUpdateParams{
		ID: uint(id),
	}
	err = c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("modules.AccountRequestHandler.Update: error bad request: %s", err)
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	resp, err := h.accountController.Update(c, req)
	if err != nil {
		log.Printf("modules.AccountRequestHandler.Update: error update account: %s", err)
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, resp)
}
func (h *AccountRequestHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Printf("modules.AccountRequestHandler.Delete: error parsing path param: %s", err)
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	err = h.accountController.Delete(c, uint(id))
	if err != nil {
		log.Printf("modules.AccountRequestHandler.Delete: error delete account: %s", err)
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *AccountRequestHandler) FindByUsername(c *gin.Context) {
	accountQuery := AccountUpdateParams{}
	err := c.ShouldBindQuery(&accountQuery)
	if err != nil {
		log.Printf("modules.AccountRequestHandler.FindByUsername: error bind query: %s", err)
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	resp, err := h.accountController.FindByUsername(accountQuery.Page, accountQuery.UserName)
	if err != nil {
		log.Printf("modules.AccountRequestHandler.FindByUsername: error find account: %s", err)
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *AccountRequestHandler) UpdateActivatedAccount(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Printf("modules.AccountRequestHandler.UpdateActivatedAccount: error parse path param: %s", err)
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	accountQuery := AccountUpdateParams{}
	err = c.ShouldBindJSON(&accountQuery)
	if err != nil {
		log.Printf("modules.AccountRequestHandler.UpdateActivatedAccount: error bind json: %s", err)
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	err = h.accountController.UpdateActivatedAccount(c, uint(id), accountQuery.Activated)
	if err != nil {
		log.Printf("modules.AccountRequestHandler.UpdateActivatedAccount: error update activated value: %s", err)
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, nil)
}
