package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Nama      string    `gorm:"not null" json:"nama"`
    NIK       string    `gorm:"unique;not null" json:"nik"`
    NoHP      string    `gorm:"unique;not null" json:"no_hp"`
    Accounts  []Account `gorm:"foreignKey:UserID" json:"accounts"`
}