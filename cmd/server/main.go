package main

import (
	"net/http"
	"uni-events-backend/config"

	// "uni-events-backend/internal/api"

	"github.com/labstack/echo/v4"
)

func main() {
    config.InitDB()
    e := echo.New()
	e.GET("/",func(c echo.Context)error{
		return c.String(http.StatusOK,"hel")
	})
    // api.RegisterRoutes(e)
    e.Logger.Fatal(e.Start(":8080"))
}