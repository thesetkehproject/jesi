package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/thesetkehproject/jesi/functions"
	models "github.com/thesetkehproject/jesi/model"
)

// PostScreenshot Post the Screenshot
func PostScreenshot(c echo.Context) (err error) {
	screenshot := new(models.Screenshot)
	if err = c.Bind(screenshot); err != nil {
		return
	}

	functions.InsertScreenshot(screenshot)

	return c.JSON(http.StatusCreated, screenshot)
}
