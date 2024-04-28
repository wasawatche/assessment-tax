package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func taxHandlers(c echo.Context) error {
	t := Tax{}
	err := c.Bind(&t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Printf("requested tax: %v", t)

	return c.JSON(http.StatusCreated, t)
}
