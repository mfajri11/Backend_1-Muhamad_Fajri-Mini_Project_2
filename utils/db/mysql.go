package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func MustOpenGormMysql() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/mini_project?charset=utf8mb4&parseTime=True&loc=UTC"))
	if err != nil {
		log.Fatal("db.MustOpenGormMysql: error open connection to mysql: %w", err)
	}

	return db

}
