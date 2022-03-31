package controllers

import (
	"log"

	"github.com/gin-gonic/gin"

	model "Final-Project-JCC-Golang-2022/model"
)

func GetAllFeedbacks(c *gin.Context) {

	db := connect()
	var response model.FeedbacksResponse
	defer db.Close()

	rows, err := db.Query("SELECT * FROM feedbacks")

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		c.Header("Content-Type", "application/json")
		c.JSON(400, response)
		return
	}

	var feedback model.Feedback
	var feedbacks []model.Feedback

	for rows.Next() {
		if err := rows.Scan(&feedback.ID, &feedback.UserId, &feedback.Feedback, &feedback.Date); err != nil {
			log.Println(err.Error())
		} else {
			feedbacks = append(feedbacks, feedback)
		}
	}

	if len(feedbacks) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = feedbacks
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}

func GetAllMyFeedbacks(c *gin.Context) {

	db := connect()
	var response model.FeedbacksResponse
	defer db.Close()

	_, userId, _, _ := validateTokenFromCookies(c)

	rows, err := db.Query("SELECT * FROM feedbacks WHERE User_Id = ?", userId)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		c.Header("Content-Type", "application/json")
		c.JSON(400, response)
		return
	}

	var feedback model.Feedback
	var feedbacks []model.Feedback

	for rows.Next() {
		if err := rows.Scan(&feedback.ID, &feedback.UserId, &feedback.Feedback, &feedback.Date); err != nil {
			log.Println(err.Error())
		} else {
			feedbacks = append(feedbacks, feedback)
		}
	}

	if len(feedbacks) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = feedbacks
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)
}

func InsertMyFeedbacks(c *gin.Context) {

	db := connect()

	var feedback model.Feedback
	var response model.FeedbackResponse
	_, userId, _, _ := validateTokenFromCookies(c)

	feedback.Feedback = c.PostForm("feedback")

	if feedback.Feedback == "" {
		response.Status = 400
		response.Message = "Please Insert your feedback"
		c.Header("Content-Type", "application/json")
		c.JSON(response.Status, response)
		return
	}

	res, errQuery := db.Exec("INSERT INTO feedbacks(User_Id, Feedback) VALUES(?, ?)", userId, feedback.Feedback)

	id, _ := res.LastInsertId()

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		feedback.ID = int(id)
		response.Data = feedback
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		log.Println(errQuery.Error())
	}
	c.Header("Content-Type", "application/json")
	c.JSON(response.Status, response)

}
