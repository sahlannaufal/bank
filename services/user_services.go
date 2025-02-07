package services

import (
    "go-assessment/models"
    "go-assessment/repositories"
    "go-assessment/utils"
    "errors"
)

type UserService struct {
    userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
    return &UserService{userRepo: userRepo}
}

func (s *UserService) RegisterUser(nama, nik, noHP string) (*models.User, error) {
    // Cek apakah NIK atau NoHP sudah digunakan
    existingUser, err := s.userRepo.FindUserByNIK(nik)
    if err == nil && existingUser != nil {
        utils.LogWarn("NIK already exists", map[string]interface{}{
            "nik": nik,
        })
        return nil, errors.New("NIK sudah digunakan")
    }

    existingUser, err = s.userRepo.FindUserByNoHP(noHP)
    if err == nil && existingUser != nil {
        utils.LogWarn("NoHP already exists", map[string]interface{}{
            "no_hp": noHP,
        })
        return nil, errors.New("NoHP sudah digunakan")
    }

    user := &models.User{
        Nama:  nama,
        NIK:   nik,
        NoHP:  noHP,
    }
    if err := s.userRepo.CreateUser(user); err != nil {
        utils.LogError("Failed to create user", map[string]interface{}{
            "error": err.Error(),
            "nama":  nama,
            "nik":   nik,
            "no_hp": noHP,
        })
        return nil, err
    }

	account := &models.Account{
        NoRekening: generateNoRekening(), 
        Saldo:      0,
        UserID:     user.ID, 
    }
    if err := s.userRepo.CreateAccount(account); err != nil {
        return nil, err
    }

    utils.LogInfo("User registered successfully", map[string]interface{}{
        "user_id": user.ID,
        "nama":    user.Nama,
    })
    return user, nil
}