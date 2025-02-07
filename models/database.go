package models

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "go-assessment/config"
)

var DB *gorm.DB

func InitDB() {
    dsn := "host=" + config.GetEnv("DB_HOST") +
           " user=" + config.GetEnv("DB_USER") +
           " password=" + config.GetEnv("DB_PASSWORD") +
           " dbname=" + config.GetEnv("DB_NAME") +
           " port=" + config.GetEnv("DB_PORT") +
           " sslmode=disable TimeZone=Asia/Jakarta"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }
    DB.AutoMigrate(&User{}, &Account{}, &Transaction{})
}