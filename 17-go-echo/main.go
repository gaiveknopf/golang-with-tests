package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := NewEchoServer()
	e.Logger.Fatal(e.Start(":3334"))
}

func HelloWorld() string {
	return fmt.Sprintf("Hello World!")
}

func NewEchoServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, HelloWorld())
	})
	return e
}
