package main

import (
	"context"
	"fmt"
	"goTempl/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}
	app.Use(withUser)

	app.GET("/user", userHandler.HandleUserShow)

	app.Start(":8080")

	fmt.Println("I'm running now")
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "jeff@jeff.com")
		c.SetRequest(c.Request().WithContext(ctx))
		fmt.Println("I'm running now" + c.Path())
		return next(c)
	}
}
