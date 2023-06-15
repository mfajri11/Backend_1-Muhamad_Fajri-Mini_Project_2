package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/dto"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/customer"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type CustomerRequestHandler struct {
	customerController ICustomerController
}

func NewCustomerRequestHandler(db *gorm.DB) *CustomerRequestHandler {
	return &CustomerRequestHandler{
		customerController: &CustomerController{
			customerUC: &CustomerUseCase{
				customerRepo: customer.NewCustomerRepository(db),
			},
		}}
}

func (h CustomerRequestHandler) Create(c *gin.Context) {
	req := CustomerParams{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("modules.CustomerRequestHandler.Create: error bind json: %s", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	resp, err := h.customerController.Create(req)
	if err != nil {
		log.Printf("modules.CustomerRequestHandler.Create: error create customer: %s", err)
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h CustomerRequestHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Printf("modules.CustomerRequestHandler.Update: error parse path params id: %s", err)
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	req := CustomerUpdateParams{ID: uint(id)}
	err = c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("modules.CustomerRequestHandler.Update: error bind json: %s", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	resp, err := h.customerController.Update(req)
	if err != nil {
		log.Printf("modules.CustomerRequestHandler.Update: error update customer: %s", err)
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h CustomerRequestHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Printf("modules.CustomerRequestHandler.Delete: error parse path params id: %s", err)
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	err = h.customerController.Delete(uint(id))
	if err != nil {
		log.Printf("modules.CustomerRequestHandler.Delete: error delete customer: %s", err)
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, nil)

}

func (h CustomerRequestHandler) Search(c *gin.Context) {
	req := CustomerUpdateParams{}
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Printf("modules.CustomerRequestHandler.Search: error bind query: %s", err)
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	resp, err := h.customerController.Search(req.Page, req.FirstName, req.Email)
	if err != nil {
		log.Printf("modules.CustomerRequestHandler.Search: error search customer: %s", err)
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, resp)
}
