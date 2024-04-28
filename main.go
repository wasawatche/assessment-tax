package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Tax struct {
	TotalIncome float64      `json:"totalIncome"`
	Wht         float64      `json:"wht"`
	Allowances  []Allowances `json:"allowances"`
}

type Allowances struct {
	AllowanceType string  `json:"allowanceType"`
	Amount        float64 `json:"amount"`
}

func taxHandlers(c echo.Context) error {
	t := Tax{}
	err := c.Bind(&t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Printf("requested tax: %v", t)

	return c.JSON(http.StatusCreated, t)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})
	e.GET("tax/calculations", taxHandlers)
	e.Logger.Fatal(e.Start(":1323"))
}
