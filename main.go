package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tongson/gl"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", file)
	e.Logger.Fatal(e.Start(":1323"))
}

func file(c echo.Context) error {
	return c.String(http.StatusOK, gl.FileRead(os.Args[1]))
}
