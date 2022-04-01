package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	model "Final-Project-JCC-Golang-2022/model"
)

// GetAllMyCart godoc
// @Summary Get all cart items.
// @Description display all cart items of users who are currently logged in.
// @Tags Cart
// @Produce json
// @Success 200 {object} model.CartsResponse
// @Router /cart [GET]
func GetAllMyCart(c *gin.Context) {

	db := connect()
	var response model.CartsResponse
	defer db.Close()
	_, userId, _, _ := validateTokenFromCookies(c)

	rows, err := db.Query("SELECT * FROM carts WHERE User_Id = ?", userId)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		c.Header("Content-Type", "application/json")
		c.JSON(400, response)
		return
	}

	var cart model.Cart
	var carts []model.Cart

	for rows.Next() {
		if err := rows.Scan(&cart.ID, &cart.UserId, &cart.ProductId, &cart.Quantity); err != nil {
			log.Println(err.Error())
		} else {
			carts = append(carts, cart)
		}
	}

	if len(carts) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = carts
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}

// DeleteMyCart godoc
// @Summary delete cart item.
// @Description delete display all cart items of users who are currently logged in.
// @Tags Cart
// @Produce json
// @Param cartId path string true "cartId"
// @Success 200 {object} model.ErrorResponse
// @Router /cart/{cartId} [delete]
func DeleteMyCart(c *gin.Context) {
	db := connect()
	defer db.Close()

	var response model.ErrorResponse
	_, userId, _, _ := validateTokenFromCookies(c)
	cartId := c.Param("cartId")

	query, errQuery := db.Exec(`DELETE FROM carts WHERE Id = ? AND User_Id = ?;`, cartId, userId)
	RowsAffected, _ := query.RowsAffected()

	if RowsAffected == 0 {
		response.Status = 400
		response.Message = "Cart item not found"
		c.Header("Content-Type", "application/json")
		c.JSON(400, response)
		return
	}

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success Delete Data"
	} else {
		response.Status = 400
		response.Message = "Error Delete Data"
		log.Println(errQuery.Error())
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}

// InsertMyCart godoc
// @Summary insert cart.
// @Description insert product to cart belongs to the user who is currently logged in.
// @Tags Cart
// @Produce json
// @Param Body body model.Cart true "cart's data"
// @Success 200 {object}  model.CartResponse
// @Router /cart [POST]
func InsertMyCart(c *gin.Context) {

	db := connect()

	var cart model.Cart
	var response model.CartResponse

	_, cart.UserId, _, _ = validateTokenFromCookies(c)
	cart.Quantity, _ = strconv.Atoi(c.PostForm("qty"))
	cart.ProductId, _ = strconv.Atoi(c.PostForm("productId"))

	if cart.Quantity == 0 {
		response.Status = 400
		response.Message = "Please Insert cart's quantity"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if cart.ProductId == 0 {
		response.Status = 400
		response.Message = "Please Insert product's id"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	res, errQuery := db.Exec("INSERT INTO carts(User_Id, Product_Id, Quantity) VALUES(?, ?, ?)", cart.UserId, cart.ProductId, cart.Quantity)

	id, _ := res.LastInsertId()

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		cart.ID = int(id)
		response.Data = cart
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		log.Println(errQuery.Error())
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)

}

// UpdateMyCart godoc
// @Summary update cart.
// @Description update cart belongs to the user who is currently logged in.
// @Tags Cart
// @Produce json
// @Param Body body model.Cart true "cart's data"
// @Success 200 {object} model.ErrorResponse
// @Router /cart/{cartId} [POST]
func UpdateMyCart(c *gin.Context) {
	db := connect()

	var cart model.Cart
	var response model.ErrorResponse

	_, cart.UserId, _, _ = validateTokenFromCookies(c)
	cart.ID, _ = strconv.Atoi(c.Param("cartId"))
	cart.Quantity, _ = strconv.Atoi(c.PostForm("qty"))

	rows, _ := db.Query("SELECT Quantity FROM carts WHERE Id = ? AND User_Id = ?", cart.ID, cart.UserId)
	var prevDatas []model.Cart
	var prevData model.Cart

	for rows.Next() {
		if err := rows.Scan(&prevData.Quantity); err != nil {
			log.Println(err.Error())
		} else {
			prevDatas = append(prevDatas, prevData)
		}
	}

	if len(prevDatas) > 0 {
		if cart.Quantity == 0 {
			cart.Quantity = prevDatas[0].Quantity
		}
		_, errQuery := db.Exec(`UPDATE carts SET Quantity = ? WHERE Id = ? AND User_Id = ?`, cart.Quantity, cart.ID, cart.UserId)

		if errQuery == nil {
			response.Status = 200
			response.Message = "Success Update Data"
		} else {
			response.Status = 400
			response.Message = "Error Update Data"

			log.Println(errQuery)
		}
	} else {
		response.Status = 400
		response.Message = "Cart Not Found"
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}
