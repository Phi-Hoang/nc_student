package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/phihdn/nc_student/db"
)

// AddStudent receives a student and insert into db
func AddStudent(c echo.Context) error {
	var student db.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.AddStudent(&student)
	if err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}

// UpdateStudent receives a student and update new data into db
func UpdateStudent(c echo.Context) error {
	var student db.StudentUpdateRequest
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.UpdateStudent(&student)
	if err != nil {
		log.Printf("update error :%v", err)
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}
