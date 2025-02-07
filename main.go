package main

import (
    "go-assessment/config"
    "go-assessment/handlers"
    "go-assessment/models"
    "go-assessment/repositories"
    "go-assessment/routes"
    "go-assessment/services"
    "go-assessment/utils"
    "github.com/labstack/echo/v4"
)

func main() {
    config.LoadEnv()

    utils.InitLogger()

    models.InitDB()

    userRepo := repositories.NewUserRepository(models.DB)
    accountRepo := repositories.NewAccountRepository(models.DB)

    userService := services.NewUserService(userRepo)
    accountService := services.NewAccountService(accountRepo)

    userHandler := handlers.NewUserHandler(userService)
    accountHandler := handlers.NewAccountHandler(accountService)

    e := echo.New()

    routes.InitRoutes(e, userHandler, accountHandler)

    utils.LogInfo("Server started on :8080", nil)
    e.Logger.Fatal(e.Start(":8080"))
}