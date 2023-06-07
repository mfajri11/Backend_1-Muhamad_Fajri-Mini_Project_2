package main

import (
	"fmt"
	register_approval "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/register-approval"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/mini_project?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		log.Fatal(err)
	}

	//acc := account.NewUserRepository(db)
	//err = acc.Create(entity.Account{Username: "test new 1", Role: entity.Role{Name: "test"}, Password: "123"})
	//err = acc.Create(entity.Account{Username: "test new 1", Role: entity.Role{Name: "test"}, Password: "123"})
	//err = acc.Create(entity.Account{Username: "test new 1", Role: entity.Role{Name: "test"}, Password: "123"})

	//err = acc.Delete(2)
	//_, err = acc.Update(&entity.Account{ID: 4, Username: "test updated"})
	//accs, err := acc.FindByUsername(1, "test updated")
	//err = acc.UpdateActivateAccount(4, "true")
	//cust := customer.NewCustomerRepository(db)
	//_, err = cust.Update(&entity.Customer{
	//	ID:        1,
	//	FirstName: "test updated",
	//	LastName:  "in dev",
	//	Email:     "test@example.com",
	//	Avatar:    "haaha",
	//})
	//cs, err := cust.Search(1, "test updated", "")
	//err = cust.Delete(1)
	appv := register_approval.NewRegisterApproval(db)
	appv.UpdateApprovalStatus(0, "approved")
	//if err != nil {
	//	log.Println(err)
	//}

	//fmt.Printf("data: %s", cs)
	fmt.Println("success")
}
