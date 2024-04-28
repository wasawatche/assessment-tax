package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// func AuthMiddleware(username, password, string, c echo.Context) (bool, error) {
// 	if username == admin
// }

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})
	e.GET("tax/calculations", calTaxHandlers)
	e.Logger.Fatal(e.Start(":1323"))
}
