package models

import "gorm.io/gorm"

type Account struct {
    gorm.Model
    NoRekening string  `gorm:"unique" json:"no_rekening"`
    Nama       string  `json:"nama"`
    NIK        string  `gorm:"unique" json:"nik"`
    NoHP       string  `gorm:"unique" json:"no_hp"`
    Saldo      float64 `json:"saldo"`
}