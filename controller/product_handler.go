package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	model "Final-Project-JCC-Golang-2022/model"
)

// GetAllMyProduct godoc
// @Summary Get all product.
// @Description display all products.
// @Tags Products
// @Produce json
// @Success 200 {object} model.ProductsResponse
// @Router /products [get]
func GetAllProducts(c *gin.Context) {

	db := connect()
	var response model.ProductsResponse
	defer db.Close()

	query := "SELECT * FROM products"
	id := c.Query("id")
	if id != "" {
		query += " WHERE id = " + id
	}

	rows, err := db.Query(query)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		c.Header("Content-Type", "application/json")
		c.JSON(400, response)
		return
	}

	var product model.Product
	var products []model.Product

	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.StoreId); err != nil {
			log.Println(err.Error())
		} else {
			products = append(products, product)
		}
	}

	if len(products) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = products
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}

// DeleteMyProduct godoc
// @Summary delete prodduct.
// @Description delete products sold by logged in users.
// @Tags Products
// @Produce json
// @Param productid path string true "productid"
// @Success 200 {object} model.ErrorResponse
// @Router /product/{productid} [delete]
func DeleteMyProduct(c *gin.Context) {
	db := connect()
	defer db.Close()

	var response model.ErrorResponse
	storeId := getStoreId(c)
	productId := c.Param("productid")

	query, errQuery := db.Exec(`DELETE FROM products WHERE Store_Id = ? AND Id = ?;`, storeId, productId)
	RowsAffected, _ := query.RowsAffected()

	if RowsAffected == 0 {
		response.Status = 400
		response.Message = "Product not found"
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

// InsertMyProduct godoc
// @Summary delete prodduct.
// @Description delete products sold by logged in users.
// @Tags Products
// @Produce json
// @Param Body body model.Product true "product's data"
// @Success 200 {object}  model.ProductResponse
// @Router /product [POST]
func InsertMyProduct(c *gin.Context) {

	db := connect()

	var product model.Product
	var response model.ProductResponse

	product.Name = c.PostForm("name")
	product.Category = c.PostForm("category")
	product.Price, _ = strconv.Atoi(c.PostForm("price"))
	product.StoreId = getStoreId(c)

	if product.Name == "" {
		response.Status = 400
		response.Message = "Please Insert product's Name"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if product.Category == "" {
		response.Status = 400
		response.Message = "Please Insert product's category"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if product.Price == 0 {
		response.Status = 400
		response.Message = "Please Insert product's price"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	res, errQuery := db.Exec("INSERT INTO products(Name, Category,  Price, Store_Id) VALUES(?, ?, ?,?)", product.Name, product.Category, product.Price, product.StoreId)

	id, _ := res.LastInsertId()

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		product.ID = int(id)
		response.Data = product
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		log.Println(errQuery.Error())
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)

}

// UpdateMyProduct godoc
// @Summary update product's data.
// @Description product updates sold by logged in users.
// @Tags Products
// @Produce json
// @Param productid path string true "productid"
// @Param Body body model.Product true "product's data"
// @Success 200 {object} model.ProductResponse
// @Router /product/{productid} [PUT]
func UpdateMyProduct(c *gin.Context) {
	db := connect()

	var product model.Product
	var response model.ProductResponse

	product.ID, _ = strconv.Atoi(c.Param("productid"))
	product.Name = c.PostForm("name")
	product.Category = c.PostForm("category")
	product.Price, _ = strconv.Atoi(c.PostForm("price"))
	product.StoreId = getStoreId(c)

	rows, _ := db.Query("SELECT * FROM products WHERE Id = ? AND Store_Id = ?", product.ID, product.StoreId)
	var prevDatas []model.Product
	var prevData model.Product

	for rows.Next() {
		if err := rows.Scan(&prevData.ID, &prevData.Name, &prevData.Category, &prevData.Price, &prevData.StoreId); err != nil {
			log.Println(err.Error())
		} else {
			prevDatas = append(prevDatas, prevData)
		}
	}

	if len(prevDatas) > 0 {
		if product.Name == "" {
			product.Name = prevDatas[0].Name
		}
		if product.Category == "" {
			product.Category = prevDatas[0].Category
		}
		if product.Price == 0 {
			product.Price = prevDatas[0].Price
		}

		_, errQuery := db.Exec(`UPDATE products SET Name = ?, Category = ?, Price = ? WHERE Id = ? AND Store_Id = ?`, product.Name, product.Category, product.Price, product.ID, product.StoreId)

		if errQuery == nil {
			response.Status = 200
			response.Message = "Success Update Data"
			product.ID = prevData.ID
			response.Data = product
		} else {
			response.Status = 400
			response.Message = "Error Update Data"

			log.Println(errQuery)
		}
	} else {
		response.Status = 400
		response.Message = "Product Not Found"
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}
