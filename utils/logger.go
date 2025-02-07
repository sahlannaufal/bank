package utils

import (
    "os"
    "github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() {
    Logger = logrus.New()

    Logger.SetOutput(os.Stdout)

    Logger.SetFormatter(&logrus.JSONFormatter{
        TimestampFormat: "2006-01-02 15:04:05",
    })

    Logger.SetLevel(logrus.InfoLevel)
}

func LogInfo(message string, fields map[string]interface{}) {
    Logger.WithFields(fields).Info(message)
}

func LogWarn(message string, fields map[string]interface{}) {
    Logger.WithFields(fields).Warn(message)
}

func LogError(message string, fields map[string]interface{}) {
    Logger.WithFields(fields).Error(message)
}

func LogFatal(message string, fields map[string]interface{}) {
    Logger.WithFields(fields).Fatal(message)
}