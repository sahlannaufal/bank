package repositories

import (
    "go-assessment/models"
    "gorm.io/gorm"
)

type AccountRepository struct {
    DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
    return &AccountRepository{DB: db}
}

func (r *AccountRepository) CreateAccount(account *models.Account) error {
    return r.DB.Create(account).Error
}

func (r *AccountRepository) FindAccountByNoRekening(noRekening string) (*models.Account, error) {
    var account models.Account
    err := r.DB.Preload("User").Where("no_rekening = ?", noRekening).First(&account).Error
    return &account, err
}

func (r *AccountRepository) UpdateAccount(account *models.Account) error {
    return r.DB.Save(account).Error
}

func (r *AccountRepository) CreateTransaction(transaction *models.Transaction) error {
    return r.DB.Create(transaction).Error
}