package controllers

import (
	"Final-Project-JCC-Golang-2022/utils"
	"database/sql"
	"fmt"
	"log"

	model "Final-Project-JCC-Golang-2022/model"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *sql.DB {
	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "production" {
		username := "evyruldwjcspkh"
		password := "4f7dbdf0e9e57f0d97f9f2e162fac8eca230840eedf71d0625b6817f6e622db8"
		host := "ec2-54-157-79-121.compute-1.amazonaws.com"
		port := 5432
		database := "daga015a6rd6v8"

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, username, password, database)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

		db2, _ := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
		db2.AutoMigrate(&model.User{}, &model.Transaction{}, &model.Store{}, &model.Cart{}, &model.Product{}, &model.ProductReview{}, &model.Feedback{})

		return db

	} else {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/jcc_ecommerce?charset=utf8mb4&parseTime=True&loc=Local")

		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Success Connect DB")
		}

		return db
	}

}
