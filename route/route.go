package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phihdn/nc_student/config"
	"github.com/phihdn/nc_student/handler"
	"github.com/phihdn/nc_student/models"
)

func All(e *echo.Echo) {
	Staff(e)
	Public(e)
}

func Staff(e *echo.Echo) {
	g := e.Group("/api/student/v1/staff")
	JWTConfig := middleware.JWTConfig{
		SigningKey: []byte(config.Config.JWTSecret.JWTKey),
		Claims:     &models.UserClaims{},
	}
	g.Use(middleware.JWTWithConfig(JWTConfig))
	g.POST("/student", handler.AddStudent)
	g.PUT("/student", handler.UpdateStudent)
	g.DELETE("/student", handler.DeleteStudent)
	g.DELETE("/student/id/:id", handler.DeleteStudentById)
}

func Public(e *echo.Echo) {
	g := e.Group("/api/student/v1/public")
	g.GET("/health", handler.HealthCheck)
	g.GET("/test", handler.TestDB)
	g.GET("/student", handler.GetAllStudents)
	g.GET("/student/id/:id", handler.GetStudentById)
	g.GET("/student/group/last_name", handler.GetAllStudentGroupByLastName)
	g.PATCH("/student/simple", handler.SearchStudentSimple)
	g.PATCH("/student", handler.SearchStudent)
}
