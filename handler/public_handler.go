package handler

import (
	"net/http"

	"github.com/Phi-Hoang/nc_student/db"
	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func TestDB(c echo.Context) error {
	db.Test()
	return c.String(http.StatusOK, "TestDB")
}
