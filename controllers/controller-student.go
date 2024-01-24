package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josebruno2020/go-api-gin/database"
	"github.com/josebruno2020/go-api-gin/models"
)

type StudentController struct {
}

// @BasePath /api/v1/students

// FindAll godoc
// @Summary Listar alunos
// @Schemes
// @Description Rota para listar todos os alunos
// @Tags students
// @Accept json
// @Produce json
// @Success 200 {object} []models.StudentView
// @Router /api/v1/students [get]
func (c *StudentController) FindAll(ctx *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)

	studentsView := []models.StudentView{}

	for _, s := range students {
		studentsView = append(studentsView, s.ToStudentView())
	}

	ctx.JSON(http.StatusOK, studentsView)
}

// FindByCPF godoc
// @Summary Buscar aluno por cpf
// @Schemes
// @Description Rota para buscar aluno por CPF
// @Tags students
// @Param cpf path string true "CPF para buscar aluno"
// @Accept json
// @Produce json
// @Success 200 {object} models.StudentView
// @Failure 404 {object} HttpError
// @Router /api/v1/students/cpf/{cpf} [get]
func (c *StudentController) FindByCPF(ctx *gin.Context) {
	cpf := ctx.Param("cpf")
	var student models.Student
	database.DB.Where("cpf = ?", cpf).First(&student)

	if student.ID == 0 {
		sendJsonError(ctx, HttpError{http.StatusNotFound, "Student not Found"})
		return
	}

	ctx.JSON(http.StatusOK, student.ToStudentView())
}

// Find godoc
// @Summary Buscar aluno por ID
// @Schemes
// @Description Rota para buscar aluno por ID
// @Tags students
// @Param id path string true "ID para buscar aluno"
// @Accept json
// @Produce json
// @Success 200 {object} models.StudentView
// @Failure 404 {object} HttpError
// @Router /api/v1/students/{id} [get]
func (c *StudentController) Find(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		sendJsonError(ctx, HttpError{http.StatusNotFound, "Student not Found"})
		return
	}

	ctx.JSON(http.StatusOK, student.ToStudentView())
}

// Create godoc
// @Summary Cadastrar aluno
// @Schemes
// @Description Rota para cadastrar aluno
// @Tags students
// @Param aluno body models.Student true "estrutura de aluno"
// @Accept json
// @Produce json
// @Success 201 {object} models.StudentView
// @Failure 400 {object} HttpError
// @Router /api/v1/students [post]
func (c *StudentController) Create(ctx *gin.Context) {
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

	ctx.JSON(http.StatusCreated, student.ToStudentView())
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
// @Success 200 {object} models.StudentView
// @Failure 400 {object} HttpError
// @Failure 404 {object} HttpError
// @Router /api/v1/students/{id} [patch]
func (c *StudentController) Update(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		sendJsonError(ctx, HttpError{http.StatusNotFound, "Student not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&student); err != nil {
		sendJsonError(ctx, HttpError{http.StatusBadRequest, err.Error()})
		return
	}

	if err := student.ValidateData(); err != nil {
		sendJsonError(ctx, HttpError{http.StatusBadRequest, err.Error()})
		return
	}

	database.DB.Save(&student)

	ctx.JSON(http.StatusOK, student.ToStudentView())
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
// @Router /api/v1/students/{id} [delete]
func (c *StudentController) Delete(ctx *gin.Context) {
	var student models.Student
	id := ctx.Params.ByName("id")
	database.DB.Delete(&student, id)

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *StudentController) IndexPage(ctx *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	ctx.HTML(200, "index.html", gin.H{
		"students": students,
	})
}
