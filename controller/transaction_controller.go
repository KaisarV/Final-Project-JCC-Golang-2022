package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	config "Final-Project-JCC-Golang-2022/config"
	model "Final-Project-JCC-Golang-2022/model"
)

type TransactionInput struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"qty"`
}

// GetAllMyTransactions godoc
// @Summary Get all transactions.
// @Description display all transactions of users who are currently logged in.
// @Tags Transactions
// @Produce json
// @Success 200 {object} model.TransactionsResponse
// @Router /transactions [get]
func GetAllMyTransactions(c *gin.Context) {
	db := config.Connect()
	var response model.TransactionsResponse
	defer db.Close()
	_, userId, _, _ := validateTokenFromCookies(c)
	rows, err := db.Query("SELECT * FROM transactions WHERE User_Id = ?", userId)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		c.Header("Content-Type", "application/json")
		c.JSON(400, response)
		return
	}

	var transaction model.Transaction
	var transactions []model.Transaction

	for rows.Next() {
		if err := rows.Scan(&transaction.ID, &transaction.UserId, &transaction.ProductId, &transaction.Date, &transaction.Quantity); err != nil {
			log.Println(err.Error())
		} else {
			transactions = append(transactions, transaction)
		}
	}

	if len(transactions) != 0 {
		response.Status = 200
		response.Message = "Success Get Transactions"
		response.Data = transactions
	} else {
		response.Status = 400
		response.Message = "Transactions Not Found"
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}

// InsertMyTransaction godoc
// @Summary insert transaction.
// @Description insert user's transaction who currently logged in.
// @Tags Transactions
// @Produce json
// @Param Body body TransactionInput true "transaction's data"
// @Success 200 {object} model.TransactionResponse
// @Router /transaction [POST]
func InsertMyTransactions(c *gin.Context) {

	db := config.Connect()

	var transaction model.Transaction
	var response model.TransactionResponse
	var input TransactionInput

	if c.Request.Header.Get("Content-Type") == "application/json" {
		if err := c.ShouldBindJSON(&input); err != nil {
			response.Status = 400
			response.Message = err.Error()
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, response)
			return
		}
	} else {
		input.Quantity, _ = strconv.Atoi(c.PostForm("qty"))
		input.ProductId, _ = strconv.Atoi(c.PostForm("productid"))
	}

	_, transaction.UserId, _, _ = validateTokenFromCookies(c)
	transaction.Quantity = input.Quantity
	transaction.ProductId = input.ProductId

	if transaction.Quantity == 0 {
		response.Status = 400
		response.Message = "Please Insert transaction's quantity"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if transaction.ProductId == 0 {
		response.Status = 400
		response.Message = "Please Insert product"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	res, errQuery := db.Exec("INSERT INTO transactions(User_Id, Product_Id, Quantity) VALUES(?, ?, ?)", transaction.UserId, transaction.ProductId, transaction.Quantity)

	id, _ := res.LastInsertId()

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		transaction.ID = int(id)
		response.Data = transaction
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		log.Println(errQuery.Error())
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)

}
