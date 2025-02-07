package services

import (
    "go-assessment/models"
    "go-assessment/repositories"
    "go-assessment/utils"
	"math/rand"
    "strconv"
    "time"
    "errors"
)

type UserService struct {
    userRepo *repositories.UserRepository
}

func generateNoRekening() string {
    rand.Seed(time.Now().UnixNano()) 
    return strconv.Itoa(rand.Intn(9000000000) + 1000000000) 
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
    return &UserService{userRepo: userRepo}
}

func (s *UserService) RegisterUser(nama, nik, noHP string) (*models.User, error) {
    existingUser, err := s.userRepo.FindUserByNIK(nik)
    if err == nil && existingUser != nil {
        utils.LogWarn("NIK already exists", map[string]interface{}{"nik": nik})
        return nil, errors.New("NIK sudah digunakan")
    }

    existingUser, err = s.userRepo.FindUserByNoHP(noHP)
    if err == nil && existingUser != nil {
        utils.LogWarn("NoHP already exists", map[string]interface{}{"no_hp": noHP})
        return nil, errors.New("NoHP sudah digunakan")
    }

    user := &models.User{
        Nama: nama,
        NIK:  nik,
        NoHP: noHP,
    }
    if err := s.userRepo.CreateUser(user); err != nil {
        utils.LogError("Failed to create user", map[string]interface{}{"error": err.Error()})
        return nil, err
    }

    if user.ID == 0 {
        utils.LogError("User ID is still empty after creation", nil)
        return nil, errors.New("gagal mendapatkan ID user")
    }

    account := &models.Account{
        NoRekening: generateNoRekening(),
        Saldo:      0,
        UserID:     user.ID,
    }
    if err := s.userRepo.CreateAccount(account); err != nil {
        utils.LogError("Failed to create account", map[string]interface{}{"error": err.Error(), "user_id": user.ID})
        return nil, err
    }

    // Ambil ulang user agar `Accounts` terisi
    updatedUser, err := s.userRepo.FindUserByID(user.ID)
    if err != nil {
        utils.LogError("Failed to retrieve user with account", map[string]interface{}{"user_id": user.ID})
        return nil, err
    }

    utils.LogInfo("User registered successfully", map[string]interface{}{"user_id": user.ID, "nama": user.Nama})
    return updatedUser, nil
}
