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
    // Load environment variables
    config.LoadEnv()

    // Initialize logger
    utils.InitLogger()

    // Initialize database
    models.InitDB()

    // Setup repositories
    userRepo := repositories.NewUserRepository(models.DB)
    accountRepo := repositories.NewAccountRepository(models.DB)

    // Setup services
    userService := services.NewUserService(userRepo)
    accountService := services.NewAccountService(accountRepo)

    // Setup handlers
    userHandler := handlers.NewUserHandler(userService)
    accountHandler := handlers.NewAccountHandler(accountService)

    // Setup Echo
    e := echo.New()

    // Initialize routes
    routes.InitRoutes(e, userHandler, accountHandler)

    // Start server
    utils.LogInfo("Server started on :8080", nil)
    e.Logger.Fatal(e.Start(":8080"))
}