package models

import (
    "gorm.io/gorm"
)

type Account struct {
    gorm.Model
    NoRekening string       `gorm:"unique;not null" json:"no_rekening"`
    Saldo      float64      `gorm:"default:0" json:"saldo"`
    UserID     uint         `gorm:"not null" json:"user_id"`
    User       User         `gorm:"foreignKey:UserID" json:"user"` 
    Transactions []Transaction `gorm:"foreignKey:AccountID" json:"transactions"`
}