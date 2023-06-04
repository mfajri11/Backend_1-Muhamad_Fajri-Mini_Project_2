package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/modules/account"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/modules/auth"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/modules/customer"
	register_approval "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/modules/register-approval"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/utils/db"
	"log"
)

func main() {
	router := gin.New()
	store := db.MustOpenGormMysql()

	accountRouter := account.NewAccountRouter(store)
	customerRouter := customer.NewCustomerRouter(store)
	approvalRouter := register_approval.NewRegisterApprovalRouter(store)
	authRouter := auth.NewAuthRouter(store)

	accountRouter.Handle(router)
	customerRouter.Handle(router)
	approvalRouter.Handle(router)
	authRouter.Handle(router)

	log.Fatal(router.Run(":8080"))
}
