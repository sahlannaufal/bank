package utils

import (
    "os"
    "github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
    Logger = logrus.New()

    // Set output ke stdout (bisa diganti ke file jika diperlukan)
    Logger.SetOutput(os.Stdout)

    // Set format log ke JSON (lebih mudah dibaca oleh log management tools)
    Logger.SetFormatter(&logrus.JSONFormatter{
        TimestampFormat: "2006-01-02 15:04:05", // Format timestamp
    })

    // Set level log ke Info (bisa diubah ke Debug, Warn, Error, dll.)
    Logger.SetLevel(logrus.InfoLevel)
}

// LogInfo untuk logging level Info
func LogInfo(message string, fields map[string]interface{}) {
    Logger.WithFields(fields).Info(message)
}

// LogWarn untuk logging level Warn
func LogWarn(message string, fields map[string]interface{}) {
    Logger.WithFields(fields).Warn(message)
}

// LogError untuk logging level Error
func LogError(message string, fields map[string]interface{}) {
    Logger.WithFields(fields).Error(message)
}

// LogFatal untuk logging level Fatal (akan menghentikan aplikasi)
func LogFatal(message string, fields map[string]interface{}) {
    Logger.WithFields(fields).Fatal(message)
}