package controllers

import (
	"Final-Project-JCC-Golang-2022/utils"
	"database/sql"
	"fmt"
	"log"
	"os"

	model "Final-Project-JCC-Golang-2022/model"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *sql.DB {
	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "production" {
		username := os.Getenv("DATABASE_USERNAME")
		password := os.Getenv("DATABASE_PASSWORD")
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABASE_PORT")
		database := os.Getenv("DATABASE_NAME")

		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, username, password, database)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

		db2, _ := gorm.Open(mysql.Open(psqlInfo), &gorm.Config{})
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
