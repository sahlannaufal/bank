package routes

import (
    "go-assessment/handlers"
    "github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, userHandler *handlers.UserHandler, accountHandler *handlers.AccountHandler) {
    // User routes
    e.POST("/daftar", userHandler.RegisterUser)

    // Account routes
    e.POST("/tabung", accountHandler.Tabung)
    e.POST("/tarik", accountHandler.Tarik)
    e.GET("/saldo/:no_rekening", accountHandler.CekSaldo)
}