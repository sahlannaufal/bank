package routes

import (
    "go-assessment/handlers"
    "github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
    e.POST("/daftar", handlers.Daftar)
    e.POST("/tabung", handlers.Tabung)
    e.POST("/tarik", handlers.Tarik)
    e.GET("/saldo/:no_rekening", handlers.Saldo)
}