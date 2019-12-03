package route

import (
	"github.com/phihdn/nc_student/handler"
	"github.com/labstack/echo/v4"
)

func All(e *echo.Echo) {
	Private(e)
	Staff(e)
	Public(e)
}

func Private(e *echo.Echo) {

}

func Staff(e *echo.Echo) {

}

func Public(e *echo.Echo) {
	g := e.Group("/student/v1/public")
	g.GET("/health", handler.HealthCheck)
	g.GET("/test", handler.TestDB)
	g.GET("/student", handler.GetAllStudents)
	g.POST("/student", handler.AddStudent)
}
