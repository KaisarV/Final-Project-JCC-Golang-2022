package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	model "Final-Project-JCC-Golang-2022/model"
)

type InputUser struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Address  string `json:"address"`
}

type InputLogin struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

// GetAllUsers godoc
// @Summary Get all users.
// @Description Display all registered users.
// @Tags Users
// @Produce json
// @Success 200 {object} model.UsersResponse
// @Router /users [get]
func GetAllUsers(c *gin.Context) {

	db := connect()
	var response model.UsersResponse
	defer db.Close()

	query := "SELECT Id, Name, Phone, Email,  Address, User_Type FROM users"
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

	var user model.User
	var users []model.User

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Email, &user.Address, &user.UserType); err != nil {
			log.Println(err.Error())
		} else {
			users = append(users, user)
		}
	}

	if len(users) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = users
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, gin.H{"data": response})
}

// DeleteUser godoc
// @Summary delete user.
// @Description Delete user by id and admin only can use it.
// @Tags Users
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} model.ErrorResponse
// @Router /user/{id} [delete]
func DeleteUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	var response model.ErrorResponse
	userId := c.Param("id")

	query, errQuery := db.Exec(`DELETE FROM users WHERE Id = ?;`, userId)
	RowsAffected, _ := query.RowsAffected()

	if RowsAffected == 0 {
		response.Status = 400
		response.Message = "User not found"
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
	c.JSON(http.StatusOK, gin.H{"data": response})
}

// InsertUser godoc
// @Summary insert user.
// @Description insert user and it use for register user.
// @Tags Users
// @Produce json
// @Param Body body InputUser true "User's data"
// @Success 200 {object} model.UserResponse
// @Router /register [POST]
func InsertUser(c *gin.Context) {

	db := connect()

	var input InputUser
	var response model.UserResponse
	var user model.User
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
		input.Phone = c.PostForm("phone")
		input.Email = c.PostForm("email")
		input.Password = c.PostForm("password")
		input.Address = c.PostForm("address")
	}

	if input.Name == "" {
		response.Status = 400
		response.Message = "Please Insert User's Name"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if input.Phone == "" {
		response.Status = 400
		response.Message = "Please Insert User's Phone"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if input.Address == "" {
		response.Status = 400
		response.Message = "Please Insert User's Address"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if input.Email == "" {
		response.Status = 400
		response.Message = "Please Insert User's Email"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if input.Password == "" {
		response.Status = 400
		response.Message = "Please Insert User's Password"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	user.Name = input.Name
	user.Phone = input.Phone
	user.Email = input.Email
	user.Password = input.Password
	user.Address = input.Address

	rows, _ := db.Query("SELECT Email FROM users WHERE Email = ?", user.Email)

	i := 0
	for rows.Next() {
		i++
	}

	if i != 0 {
		response.Status = 400
		response.Message = "Email already registered"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	res, errQuery := db.Exec("INSERT INTO users(Name, Phone,  Email, Password,Address) VALUES(?, ?, ?, ?, ?)", user.Name, user.Phone, user.Email, user.Password, user.Address)

	id, _ := res.LastInsertId()

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		user.UserType = 1
		user.ID = int(id)
		response.Data = user
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		log.Println(errQuery.Error())
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

// UpdateUser godoc
// @Summary update user.
// @Description change the data of the user who is currently logged in.
// @Tags Users
// @Produce json
// @Param Body body InputUser true "User's data"
// @Success 200 {object} model.UserResponse
// @Router /user [PUT]
func UpdateUsers(c *gin.Context) {
	db := connect()

	var user model.User
	var response model.UserResponse
	var input InputUser
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
		input.Phone = c.PostForm("phone")
		input.Email = c.PostForm("email")
		input.Password = c.PostForm("password")
		input.Address = c.PostForm("address")
	}

	userId, _ := getCurrentUserId(c)
	user.Name = input.Name
	user.Phone = input.Phone
	user.Email = input.Email
	user.Password = input.Password
	user.Address = input.Address

	rows, _ := db.Query("SELECT * FROM users WHERE Id = ?", userId)
	var prevDatas []model.User
	var prevData model.User

	for rows.Next() {
		if err := rows.Scan(&prevData.ID, &prevData.Name, &prevData.Phone, &prevData.Email, &prevData.Password, &prevData.Address, &prevData.UserType); err != nil {
			log.Println(err.Error())
		} else {
			prevDatas = append(prevDatas, prevData)
		}
	}

	if len(prevDatas) > 0 {
		if user.Name == "" {
			user.Name = prevDatas[0].Name
		}
		if user.Address == "" {
			user.Address = prevDatas[0].Address
		}
		if user.Phone == "" {
			user.Phone = prevDatas[0].Phone
		}
		if user.Email == "" {
			user.Email = prevDatas[0].Email
		}
		if user.Password == "" {
			user.Password = prevDatas[0].Password
		}

		_, errQuery := db.Exec(`UPDATE users SET Name = ?,  Phone = ?, Email = ?, Password = ?,Address = ? WHERE id = ?`, user.Name, user.Phone, user.Email, user.Password, user.Address, userId)

		if errQuery == nil {
			response.Status = 200
			response.Message = "Success Update Data"
			user.ID = userId
			response.Data = user
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

	if user.Name != "" {
		generateToken(c, user.ID, user.Name, user.UserType)
	}
}

// UserLogin godoc
// @Summary login user.
// @Description login for registered users.
// @Tags Users
// @Produce json
// @Param Body body InputLogin true "User's login data"
// @Success 200 {object} model.ErrorResponse
// @Router /login [POST]
func UserLogin(c *gin.Context) {
	db := connect()
	defer db.Close()
	var response model.ErrorResponse

	var input InputLogin

	if c.Request.Header.Get("Content-Type") == "application/json" {

		if err := c.ShouldBindJSON(&input); err != nil {
			response.Status = 400
			response.Message = err.Error()
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, response)
			return
		}
	} else {
		input.Email = c.PostForm("email")
		input.Password = c.PostForm("password")
	}

	email := input.Email
	password := input.Password

	if email == "" || password == "" {
		response.Status = 400
		response.Message = "Please input name and password"

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, response)
	}

	rows, err := db.Query("SELECT * FROM users WHERE email=? AND password=?",
		email,
		password,
	)

	if err != nil {
		log.Fatal(err)
	}

	var user model.User
	var users []model.User

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Email, &user.Password, &user.Address, &user.UserType); err != nil {
			log.Print(err.Error())
		} else {
			users = append(users, user)
		}
	}

	if len(users) == 1 {
		generateToken(c, user.ID, user.Name, user.UserType)
		sendSuccessResponse(c)
	} else {
		sendErrorResponse(c)
	}

}

// Logout godoc
// @Summary logout user.
// @Description logout user.
// @Tags Users
// @Produce json
// @Router /logout [GET]
func Logout(c *gin.Context) {
	resetUserToken(c)
	sendSuccessResponse(c)
}

func sendSuccessResponse(c *gin.Context) {
	var response model.ErrorResponse
	response.Status = 200
	response.Message = "Success"
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func sendErrorResponse(c *gin.Context) {
	var response model.ErrorResponse
	response.Status = 400
	response.Message = "Failed"
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusBadRequest, response)
}

func sendUnAuthorizedResponse(c *gin.Context) {
	var response model.ErrorResponse
	response.Status = 401
	response.Message = "Unauthorized Access"
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusUnauthorized, response)
}
