package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josebruno2020/go-api-gin/database"
	"github.com/josebruno2020/go-api-gin/models"
)

type StudentController struct {
}

// @BasePath /students

// FindAll godoc
// @Summary Listar alunos
// @Schemes
// @Description Rota para listar todos os alunos
// @Tags students
// @Accept json
// @Produce json
// @Success 200 {object} []models.Student
// @Router /students [get]
func FindAll(ctx *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	ctx.JSON(http.StatusOK, students)
}

// FindByCPF godoc
// @Summary Buscar aluno por cpf
// @Schemes
// @Description Rota para buscar aluno por CPF
// @Tags students
// @Param cpf path string true "CPF para buscar aluno"
// @Accept json
// @Produce json
// @Success 200 {object} models.Student
// @Failure 404 {object} HttpError
// @Router /students/cpf/{cpf} [get]
func FindByCPF(ctx *gin.Context) {
	cpf := ctx.Param("cpf")
	var student models.Student
	database.DB.Where("cpf = ?", cpf).First(&student)

	if student.ID == 0 {
		sendJsonError(ctx, HttpError{http.StatusNotFound, "Student not Found"})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

// Find godoc
// @Summary Buscar aluno por ID
// @Schemes
// @Description Rota para buscar aluno por ID
// @Tags students
// @Param id path string true "ID para buscar aluno"
// @Accept json
// @Produce json
// @Success 200 {object} models.Student
// @Failure 404 {object} HttpError
// @Router /students/{id} [get]
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

// Create godoc
// @Summary Cadastrar aluno
// @Schemes
// @Description Rota para cadastrar aluno
// @Tags students
// @Param aluno body models.Student true "estrutura de aluno"
// @Accept json
// @Produce json
// @Success 201 {object} models.Student
// @Failure 400 {object} HttpError
// @Router /students [post]
func Create(ctx *gin.Context) {
	var student models.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		sendJsonError(ctx, HttpError{http.StatusBadRequest, err.Error()})
		return
	}

	if err := student.ValidateData(); err != nil {
		sendJsonError(ctx, HttpError{http.StatusBadRequest, err.Error()})
		return
	}

	database.DB.Create(&student)

	ctx.JSON(http.StatusCreated, student)
}

// Update godoc
// @Summary Atualizar aluno
// @Schemes
// @Description Rota para atualizar aluno
// @Tags students
// @Param aluno body models.Student true "estrutura de aluno"
// @Param id    path string true "ID do aluno"
// @Accept json
// @Produce json
// @Success 200 {object} models.Student
// @Failure 400 {object} HttpError
// @Router /students/{id} [patch]
func Update(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")
	database.DB.First(&student, id)

	if err := ctx.ShouldBindJSON(&student); err != nil {
		sendJsonError(ctx, HttpError{http.StatusBadRequest, err.Error()})
		return
	}

	if err := student.ValidateData(); err != nil {
		sendJsonError(ctx, HttpError{http.StatusBadRequest, err.Error()})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)

	ctx.JSON(http.StatusOK, student)
}

// Delete godoc
// @Summary Deletar aluno
// @Schemes
// @Description Rota para deletar aluno
// @Tags students
// @Param id path string true "ID do aluno"
// @Accept json
// @Produce json
// @Success 204
// @Router /students/{id} [delete]
func Delete(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")
	database.DB.Delete(&student, id)

	ctx.JSON(http.StatusNoContent, nil)
}

func IndexPage(ctx *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	ctx.HTML(200, "index.html", gin.H{
		"students": students,
	})
}
