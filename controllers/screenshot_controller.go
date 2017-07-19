package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	models "github.com/thesetkehproject/jesi/model"
)

// PostScreenshot
func PostScreenshot(c echo.Context) error {
	screenshot := new(models.Screenshot)
	var err error
	return c.JSON(http.StatusCreated, screenshot)
}
