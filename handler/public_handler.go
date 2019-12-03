package handler

import (
	"encoding/json"
	"net/http"

	"github.com/phihdn/nc_student/db"
	"github.com/labstack/echo/v4"
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
	type Student struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Age       int    `json:"age"`
		Email     string `json:"email"`
		ClassName string `json:"class_name"`
	}
	inputJSON := `[{"first_name":"Tam", "last_name":"Nguyen","age":100,"email":"tamnguyen@abc.com"},{"first_name":"Binh","last_name":"Hoang","age":3,"email":"binh@hoang.com"}]`
	var students []Student
	json.Unmarshal([]byte(inputJSON), &students)

	return c.JSON(http.StatusOK, students)
}

// AddStudent receives a student and insert into db
func AddStudent(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
