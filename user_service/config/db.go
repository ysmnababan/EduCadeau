package config

import (
	"fmt"
	"log"
	"user_service/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		helper.LOCALHOST,
		helper.USER,
		helper.PASSWORD,
		helper.DBNAME,
		helper.DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("unable connect to database", err)
	}

	fmt.Println("CONNECTED TO USER DATABASE")
	return db
}
