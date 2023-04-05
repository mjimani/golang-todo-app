package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	var tm = NewTodoManager() // 👈 instance of todo manager

	e := echo.New()

	// 👇 new GET route
	e.GET("/", func(c echo.Context) error {
		todos := tm.GetAll()

		return c.JSON(http.StatusOK, todos) // 👈 sending json data back
	})

	e.Start(":8888")
}