package main

import (
	"log"

	"github.com/Phi-Hoang/nc_student/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	route.All(e)
	log.Println(e.Start(":9090"))
}
