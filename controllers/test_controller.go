package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/thesetkehproject/jesi/model"
)

//TestRest Rest Endpoint test
func TestRest(c echo.Context) error {
	TestOuput := &model.Test{Test1: "Hello", Test2: "World"}

	fmt.Println(TestOuput)
	return c.JSONPretty(http.StatusOK, TestOuput, " ")
}
