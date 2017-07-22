package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/thesetkehproject/jesi/controllers"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/screenshot", controllers.PostScreenshot)
	e.GET("/test", controllers.TestRest)

	// Server
	e.Logger.Fatal(e.Start(":1323"))
}
