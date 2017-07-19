package main

import (
	"github.com/thesetkehproject/jesi/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", controllers.PostScreenshot(e))

	// Server
	e.Run(standard.New(":1323"))
}