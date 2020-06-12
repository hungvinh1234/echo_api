package api

import (
	"example.com/m/src/api/handlers"

	"github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	e.GET("/login", handlers.Login)
	e.GET("/yallo", handlers.Yallo)
	e.GET("/cats/:data", handlers.GetCats)

	e.POST("/cats", handlers.AddCat)
	e.POST("/dogs", handlers.AddDog)
	e.POST("/hamsters", handlers.AddHamster)

	//employees management
	e.POST("/employee", handlers.AddEmployee)
	e.PUT("/update", handlers.UpdateEmployee)
	e.DELETE("/employee/:id", handlers.DeleteEmployee)

	// e.GET("/allemployee", handlers.ShowEmployee)

}
