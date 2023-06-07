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
	router.Use(gin.Recovery())
	store := db.MustOpenGormMysql()

	accountRouter := account.NewAccountRouter(store)
	customerRouter := customer.NewCustomerRouter(store)
	approvalRouter := register_approval.NewRegisterApprovalRouter(store)
	authRouter := auth.NewAuthRouter(store)
	authRouter.Handle(router)
	// wrapper for middleware
	authorizeWrapper := authRouter.UseAuthorizationRequired(router)
	accountRouter.Handle(authorizeWrapper)
	customerRouter.Handle(authorizeWrapper)
	approvalRouter.Handle(authorizeWrapper)

	log.Fatal(router.Run(":8080"))
}
