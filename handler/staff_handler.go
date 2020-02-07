package handler

import (
	"github.com/phihdn/nc_student/models"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/phihdn/nc_student/db"
)

// AddStudent receives a student and insert into db
func AddStudent(c echo.Context) error {
	var student models.StudentAddRequest
	if err := c.Bind(&student); err != nil {
		log.Printf("req error :%v", err)
		return c.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.AddStudent(&student)
	if err != nil {
		log.Printf("add error :%v", err)
		return c.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}

// UpdateStudent receives a student and update new data into db
func UpdateStudent(c echo.Context) error {
	var student models.StudentUpdateRequest
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.UpdateStudent(&student)
	if err != nil {
		log.Printf("update error :%v", err)
		return c.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}

func DeleteStudent(c echo.Context) error {
	var student models.StudentDeleteRequest
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.DeleteStudent(&student)
	if err != nil {
		log.Printf("delete error :%v", err)
		return c.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}

func DeleteStudentById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	student, err := db.DeleteStudentById(id)
	if err != nil {
		log.Printf("student by id error :%v", err)
		return c.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	return c.JSON(http.StatusOK, student)
}