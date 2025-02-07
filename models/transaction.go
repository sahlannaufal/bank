package models

import (
    "gorm.io/gorm" // Tambahkan import ini
)

type TransactionType string

const (
    TransactionTypeTabung TransactionType = "tabung"
    TransactionTypeTarik  TransactionType = "tarik"
)

type Transaction struct {
    gorm.Model
    Type      TransactionType `gorm:"not null" json:"type"` 
    Nominal   float64         `gorm:"not null" json:"nominal"`
    AccountID uint            `gorm:"not null" json:"account_id"` 
    Account   Account         `gorm:"foreignKey:AccountID" json:"account"` 
}