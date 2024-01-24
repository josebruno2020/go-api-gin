package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/josebruno2020/go-api-gin/database"
	"github.com/josebruno2020/go-api-gin/models"
	"github.com/josebruno2020/go-api-gin/routes"
	"github.com/stretchr/testify/assert"
)

var ID int

func CreateMockStudent() {
	student := models.Student{
		Name: "Aluno teste",
		CPF:  "09123456789",
		RG:   "11111111111",
	}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteMockStudent() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func SetupDatabase() {
	database.ConnectDB()
	CreateMockStudent()
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func TestHealthCheckRoute(t *testing.T) {
	r := routes.HandleRequest()

	req, _ := http.NewRequest("GET", "/", nil)
	res := MakeRequest(r, req)

	assert.Equal(t, http.StatusOK, res.Code)

	mockResponse := `{"ping":"pong"}`

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		t.Fatalf("Erro: %s", err.Error())
	}

	assert.Equal(t, mockResponse, string(resBody))
}

func TestFindAllStudents(t *testing.T) {
	SetupDatabase()
	defer DeleteMockStudent()
	r := routes.HandleRequest()

	req, _ := http.NewRequest("GET", "/students", nil)
	res := MakeRequest(r, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var s []models.Student
	var studentsFromDB []models.Student
	database.DB.Find(&studentsFromDB)

	json.NewDecoder(res.Body).Decode(&s)

	assert.Equal(t, studentsFromDB, s)
}

func TestFindByCPF(t *testing.T) {
	SetupDatabase()
	defer DeleteMockStudent()
	r := routes.HandleRequest()

	cpf := "09123456789"

	req, _ := http.NewRequest("GET", "/students/cpf/"+cpf, nil)
	res := MakeRequest(r, req)

	var studentResponse models.Student

	json.NewDecoder(res.Body).Decode(&studentResponse)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, cpf, studentResponse.CPF)
}

func TestFindById(t *testing.T) {
	SetupDatabase()
	defer DeleteMockStudent()

	r := routes.HandleRequest()
	pathSearch := fmt.Sprintf("/students/%d", ID)
	req, _ := http.NewRequest("GET", pathSearch, nil)

	res := MakeRequest(r, req)

	var studentResponse models.Student
	json.NewDecoder(res.Body).Decode(&studentResponse)

	assert.Equal(t, ID, int(studentResponse.ID))
	assert.Equal(t, "Aluno teste", studentResponse.Name)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestDeleteById(t *testing.T) {
	SetupDatabase()

	r := routes.HandleRequest()
	pathDelete := fmt.Sprintf("/students/%d", ID)
	req, _ := http.NewRequest("DELETE", pathDelete, nil)

	res := MakeRequest(r, req)

	assert.Equal(t, http.StatusNoContent, res.Code)
}

func TestUpdateById(t *testing.T) {
	SetupDatabase()
	defer DeleteMockStudent()

	r := routes.HandleRequest()
	pathUpdate := fmt.Sprintf("/students/%d", ID)
	studentToEdit := models.Student{
		Name: "Aluno teste Atualizado",
		CPF:  "09123456780",
		RG:   "987654321",
	}
	jsonValue, _ := json.Marshal(studentToEdit)

	req, _ := http.NewRequest("PATCH", pathUpdate, bytes.NewBuffer(jsonValue))

	res := MakeRequest(r, req)

	var responseStudent models.Student

	if err := json.Unmarshal(res.Body.Bytes(), &responseStudent); err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, studentToEdit.Name, responseStudent.Name)
	assert.Equal(t, studentToEdit.CPF, responseStudent.CPF)
	assert.Equal(t, studentToEdit.RG, responseStudent.RG)
}

func MakeRequest(r *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	return res
}
