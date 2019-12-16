package handler

import (
	"net/http"

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
	// var students []db.Student
	// inputJSON := `[{"first_name":"Tam", "last_name":"Nguyen","age":100,"email":"tamnguyen@abc.com"},{"first_name":"Binh","last_name":"Hoang","age":3,"email":"binh@hoang.com"}]`
	// json.Unmarshal([]byte(inputJSON), &students)

	students, err := db.GetAllStudent()
	if err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, students)
}
