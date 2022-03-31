package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	model "Final-Project-JCC-Golang-2022/model"
)

func GetAllMyProductReviews(c *gin.Context) {

	db := connect()
	var response model.ProductReviewsResponse
	defer db.Close()

	_, userId, _, _ := validateTokenFromCookies(c)

	rows, err := db.Query("SELECT * FROM product_reviews WHERE User_Id = ?", userId)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		c.Header("Content-Type", "application/json")
		c.JSON(400, response)
		return
	}

	var review model.ProductReview
	var reviews []model.ProductReview

	for rows.Next() {
		if err := rows.Scan(&review.ID, &review.UserId, &review.ProductId, &review.Review, &review.Rating, &review.Date); err != nil {
			log.Println(err.Error())
		} else {
			reviews = append(reviews, review)
		}
	}

	if len(reviews) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = reviews
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}

func DeleteMyProductReview(c *gin.Context) {
	db := connect()
	defer db.Close()

	var response model.ErrorResponse
	_, userId, _, _ := validateTokenFromCookies(c)
	productId := c.Param("productid")

	query, errQuery := db.Exec(`DELETE FROM product_reviews WHERE User_Id = ? AND Product_Id = ?;`, userId, productId)
	RowsAffected, _ := query.RowsAffected()

	if RowsAffected == 0 {
		response.Status = 400
		response.Message = "store not found"
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

func InsertMyProductReview(c *gin.Context) {

	db := connect()

	var review model.ProductReview
	var response model.ProductReviewResponse
	_, userId, _, _ := validateTokenFromCookies(c)
	productId := c.Param("productid")

	rows, err := db.Query("SELECT * FROM transactions WHERE User_Id = ? AND Product_Id = ?", userId, productId)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	i := 0
	for rows.Next() {
		i++
	}

	if i == 0 {
		response.Status = 400
		response.Message = "you haven't bought this product"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	rows, err = db.Query("SELECT * FROM product_reviews WHERE User_Id = ? AND Product_Id = ?", userId, productId)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	i = 0
	for rows.Next() {
		i++
	}

	if i != 0 {
		response.Status = 400
		response.Message = "you already review this product"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	review.Review = c.PostForm("review")
	review.Rating, _ = strconv.Atoi(c.PostForm("rating"))

	if review.Rating > 5 {
		response.Status = 400
		response.Message = "rating can't be more than 5"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if review.Review == "" {
		response.Status = 400
		response.Message = "Please Insert your review"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	if review.Rating == 0 {
		response.Status = 400
		response.Message = "Please insert product's rating"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	res, errQuery := db.Exec("INSERT INTO product_reviews(User_Id, Product_Id,  Review, Rating) VALUES(?, ?, ?, ?)", userId, productId, review.Review, review.Rating)

	id, _ := res.LastInsertId()

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		review.ID = int(id)
		response.Data = review
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		log.Println(errQuery.Error())
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)

}

func UpdateMyProductReview(c *gin.Context) {
	db := connect()

	var review model.ProductReview
	var response model.ProductReviewResponse
	_, userId, _, _ := validateTokenFromCookies(c)
	productId := c.Param("productid")

	review.Review = c.PostForm("review")
	review.Rating, _ = strconv.Atoi(c.PostForm("rating"))

	rows, _ := db.Query("SELECT Review, Rating FROM product_reviews WHERE User_Id = ? AND Product_Id = ?", userId, productId)
	var prevDatas []model.ProductReview
	var prevData model.ProductReview

	for rows.Next() {
		if err := rows.Scan(&prevData.Review, &prevData.Rating); err != nil {
			log.Println(err.Error())
		} else {
			prevDatas = append(prevDatas, prevData)
		}
	}

	if len(prevDatas) > 0 {
		if review.Review == "" {
			review.Review = prevDatas[0].Review
		}
		if review.Rating == 0 {
			review.Rating = prevDatas[0].Rating
		}

		_, errQuery := db.Exec(`UPDATE product_reviews SET Review = ?, Rating = ? WHERE User_Id = ? AND Product_Id = ?`, review.Review, review.Rating, userId, productId)

		if errQuery == nil {
			response.Status = 200
			response.Message = "Success Update Data"
			review.ID = prevData.ID
			response.Data = review
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
