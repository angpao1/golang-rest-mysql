package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/users/add", createUser)
	e.GET("/users", getUser)
	e.GET("/users/:id", getUserByID)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.POST("/test", returnResquest)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
