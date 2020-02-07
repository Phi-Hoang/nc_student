package handler

import (
	"github.com/phihdn/nc_student/models"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/phihdn/nc_student/db"
)

// HealthCheck function to test server is live or not
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// TestDB test the DB is good or not
func TestDB(c echo.Context) error {
	db.Test()
	return c.String(http.StatusOK, "TestDB")
}

// GetAllStudents returns all students
func GetAllStudents(c echo.Context) error {
	/* type Student struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Age       int    `json:"age"`
		Email     string `json:"email"`
		ClassName string `json:"class_name"`
	}
	inputJSON := `[{"first_name":"Tam", "last_name":"Nguyen","age":100,"email":"tamnguyen@abc.com"},{"first_name":"Binh","last_name":"Hoang","age":3,"email":"binh@hoang.com"}]`
	var students []Student
	json.Unmarshal([]byte(inputJSON), &students)*/

	students, err := db.GetAllStudent()
	if err != nil {
		log.Printf("get All student error :%v", err)
		return c.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, students)
}

func SearchStudentSimple(c echo.Context) error {

	var req models.StudentSearchRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	students, err := db.SearchStudentSimple(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, students)
}

// GetAllStudentGroupByLastName returns all students and group by last name
func GetAllStudentGroupByLastName(c echo.Context) error {
	students, err := db.GetAllStudentGroupByLastName()
	if err != nil {
		log.Printf("group last name student error :%v", err)
		return c.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, students)
}

func GetStudentById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	student, err := db.GetStudentById(id)
	if err != nil {
		log.Printf("student by id error :%v", err)
		return c.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	return c.JSON(http.StatusOK, student)
}