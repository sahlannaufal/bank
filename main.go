package main

import (
    "go-assessment/config"
    "go-assessment/models"
    "go-assessment/routes"
    "github.com/labstack/echo/v4"
)

func main() {
    config.LoadEnv()
    models.InitDB()

    e := echo.New()
    routes.InitRoutes(e)

    e.Logger.Fatal(e.Start(":8080"))
}