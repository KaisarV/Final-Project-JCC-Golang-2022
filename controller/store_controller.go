package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	model "Final-Project-JCC-Golang-2022/model"
)

type StoreInput struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// GetAllStores godoc
// @Summary Get all stores.
// @Description Display all stores.
// @Tags Stores
// @Produce json
// @Success 200 {object} model.StoresResponse
// @Router /stores [get]
func GetAllStores(c *gin.Context) {

	db := connect()
	var response model.StoresResponse
	defer db.Close()

	query := "SELECT * FROM stores"
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

	var store model.Store
	var stores []model.Store

	for rows.Next() {
		if err := rows.Scan(&store.ID, &store.UserId, &store.Name, &store.Address); err != nil {
			log.Println(err.Error())
		} else {
			stores = append(stores, store)
		}
	}

	if len(stores) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = stores
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}

// DeleteMyStore godoc
// @Summary delete store.
// @Description delete the store belonging to the user who is currently logged in.
// @Tags Stores
// @Produce json
// @Success 200 {object} model.ErrorResponse
// @Router /store [delete]
func DeleteMyStore(c *gin.Context) {
	db := connect()
	defer db.Close()

	var response model.ErrorResponse
	_, userId, _, _ := validateTokenFromCookies(c)

	query, errQuery := db.Exec(`DELETE FROM stores WHERE User_Id = ?;`, userId)
	RowsAffected, _ := query.RowsAffected()

	if RowsAffected == 0 {

		response.Status = 400
		response.Message = "store not found"
		c.Header("Content-Type", "application/json")
		c.JSON(400, response)
		return
	}

	if errQuery == nil {
		query, errQuery = db.Exec("UPDATE users SET User_Type = ? WHERE Id = ?", 1, userId)
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

// InsertMyStore godoc
// @Summary insert store.
// @Description added the store of the currently logged in user.
// @Tags Stores
// @Produce json
// @Param Body body StoreInput true "store's data"
// @Success 200 {object} model.StoreResponse
// @Router /store [POST]
func InsertMyStore(c *gin.Context) {

	db := connect()

	var store model.Store
	var response model.StoreResponse
	var userType int
	var userName string
	var input StoreInput
	_, store.UserId, userName, userType = validateTokenFromCookies(c)

	if userType == 2 {
		response.Status = 400
		response.Message = "You already have a store"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if c.Request.Header.Get("Content-Type") == "application/json" {
		if err := c.ShouldBindJSON(&input); err != nil {
			response.Status = 400
			response.Message = err.Error()
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, response)
			return
		}
	} else {
		input.Name = c.PostForm("name")
		input.Address = c.PostForm("address")
	}

	store.Name = input.Name
	store.Address = input.Address

	if store.Name == "" {
		response.Status = 400
		response.Message = "Please Insert store's Name"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if store.Address == "" {
		response.Status = 400
		response.Message = "Please Insert store's Address"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	res, errQuery := db.Exec("INSERT INTO stores(User_Id, Name,  Address) VALUES(?, ?, ?)", store.UserId, store.Name, store.Address)

	id, _ := res.LastInsertId()

	if errQuery == nil {
		res, errQuery = db.Exec("UPDATE users SET User_Type = ? WHERE Id = ?", 2, store.UserId)
		response.Status = 200
		response.Message = "Success"
		store.ID = int(id)
		response.Data = store
		generateToken(c, store.UserId, userName, 2)
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		log.Println(errQuery.Error())
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)

}

// UpdateMyStore godoc
// @Summary update store.
// @Description update the store of the currently logged in user.
// @Tags Stores
// @Produce json
// @Param Body body StoreInput true "store's data"
// @Success 200 {object} model.StoreResponse
// @Router /store [PUT]
func UpdateMyStore(c *gin.Context) {
	db := connect()

	var store model.Store
	var response model.StoreResponse
	var input StoreInput

	if c.Request.Header.Get("Content-Type") == "application/json" {
		if err := c.ShouldBindJSON(&input); err != nil {
			response.Status = 400
			response.Message = err.Error()
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, response)
			return
		}
	} else {
		input.Name = c.PostForm("name")
		input.Address = c.PostForm("address")
	}
	_, userId, _, _ := validateTokenFromCookies(c)
	store.Name = input.Name
	store.Address = input.Address

	rows, _ := db.Query("SELECT * FROM stores WHERE User_Id = ?", userId)
	var prevDatas []model.Store
	var prevData model.Store

	for rows.Next() {
		if err := rows.Scan(&prevData.ID, &prevData.UserId, &prevData.Name, &prevData.Address); err != nil {
			log.Println(err.Error())
		} else {
			prevDatas = append(prevDatas, prevData)
		}
	}

	if len(prevDatas) > 0 {
		if store.Name == "" {
			store.Name = prevDatas[0].Name
		}
		if store.Address == "" {
			store.Address = prevDatas[0].Address
		}

		_, errQuery := db.Exec(`UPDATE stores SET Name = ?,  Address = ? WHERE User_id = ?`, store.Name, store.Address, userId)

		if errQuery == nil {
			response.Status = 200
			response.Message = "Success Update Data"
			store.ID = prevData.ID
			response.Data = store
		} else {
			response.Status = 400
			response.Message = "Error Update Data"

			log.Println(errQuery)
		}
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}

func getStoreId(c *gin.Context) int {
	db := connect()
	_, userId, _, _ := validateTokenFromCookies(c)
	rows, _ := db.Query("SELECT Id FROM stores WHERE User_Id = ?", userId)

	var store model.Store
	var stores []model.Store

	for rows.Next() {
		if err := rows.Scan(&store.UserId); err != nil {
			log.Println(err.Error())
		} else {
			stores = append(stores, store)
		}
	}

	return stores[0].UserId
}
