package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josebruno2020/go-api-gin/database"
	"github.com/josebruno2020/go-api-gin/models"
)

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func sendJsonError(ctx *gin.Context, err HttpError) {
	ctx.JSON(err.Code, err)
}

func FindAll(ctx *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	ctx.JSON(http.StatusOK, students)
}

func Find(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		sendJsonError(ctx, HttpError{http.StatusNotFound, "Student not Found"})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func Create(ctx *gin.Context) {
	var student models.Student

	err := ctx.ShouldBindJSON(&student)

	if err != nil {
		sendJsonError(ctx, HttpError{http.StatusBadRequest, err.Error()})
		return
	}
	database.DB.Create(&student)

	ctx.JSON(http.StatusCreated, student)
}
