package controllers

import (
	"Final-Project-JCC-Golang-2022/utils"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *sql.DB {
	environment := utils.Getenv("ENVIRONMENT", "development")
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaa" + environment)
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

		db2, _ := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

		db2.AutoMigrate(User{}, Transaction{}, Store{}, Cart{}, Product{}, ProductReview{}, Feedback{})

		return db

	} else {
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_tes_jcc?charset=utf8mb4&parseTime=True&loc=Local")

		username := "root"
		password := ""
		host := "tcp(127.0.0.1:3306)"
		database := "db_tes_jcc"

		dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)
		db2, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		db2.AutoMigrate(&User{}, &Transaction{}, &Store{}, &Cart{}, &Product{}, &ProductReview{}, &Feedback{})

		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Success Connect DB")
		}

		return db
	}

}

type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	Address   string `json:"address"`
	User_Type int    `json:"usertype"`
}

type Transaction struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	User_Id    int    `json:"userId"`
	Product_Id int    `json:"productId"`
	Date       string `json:"date"`
	Quantity   int    `json:"qty"`
}

type Store struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	User_Id int    `json:"userId,omitempty"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Product struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Store_Id int    `json:"storeId"`
}

type ProductReview struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	User_Id   int    `json:"userid,omitempty"`
	ProductId int    `json:"productid,omitempty"`
	Review    string `json:"review"`
	Rating    int    `json:"rating,omitempty"`
	Date      string `json:"Date,omitempty"`
}

type Feedback struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	User_Id  int    `json:"userid,omitempty"`
	Feedback string `json:"feedback"`
	Date     string `json:"Date,omitempty"`
}

type Cart struct {
	ID        uint `gorm:"primary_key" json:"id"`
	User_Id   int  `json:"userid"`
	ProductId int  `json:"ProductId"`
	Quantity  int  `json:"qty"`
}
