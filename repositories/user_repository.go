package repositories

import (
    "go-assessment/models"
    "gorm.io/gorm"
)

type UserRepository struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
    return r.DB.Create(user).Error
}

func (r *UserRepository) FindUserByNIK(nik string) (*models.User, error) {
    var user models.User
    err := r.DB.Where("nik = ?", nik).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) FindUserByNoHP(noHP string) (*models.User, error) {
    var user models.User
    err := r.DB.Where("no_hp = ?", noHP).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) CreateAccount(account *models.Account) error {
    return r.DB.Create(account).Error
}

func (r *UserRepository) FindUserByID(userID uint) (*models.User, error) {
    var user models.User
    err := r.DB.Preload("Accounts").First(&user, userID).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}