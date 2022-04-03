package controllers

import (
	"Final-Project-JCC-Golang-2022/utils"
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "production" {
		username := os.Getenv("DATABASE_USERNAME")
		password := os.Getenv("DATABASE_PASSWORD")
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABASE_PORT")
		database := os.Getenv("DATABASE_NAME")

		psqlInfo := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=require"
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

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
