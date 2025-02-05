package models

import "gorm.io/gorm"

type Account struct {
    gorm.Model
    NoRekening string  `gorm:"unique"`
    Nama       string
    NIK        string  `gorm:"unique"`
    NoHP       string  `gorm:"unique"`
    Saldo      float64
}