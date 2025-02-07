package services

import (
    "go-assessment/models"
    "go-assessment/repositories"
    "go-assessment/utils"
    "errors"
)

type AccountService struct {
    accountRepo *repositories.AccountRepository
}

func NewAccountService(accountRepo *repositories.AccountRepository) *AccountService {
    return &AccountService{accountRepo: accountRepo}
}

func (s *AccountService) Tabung(noRekening string, nominal float64) (*models.Account, error) {
    account, err := s.accountRepo.FindAccountByNoRekening(noRekening)
    if err != nil {
        utils.LogError("No rekening not found for Tabung", map[string]interface{}{
            "error":       err.Error(),
            "no_rekening": noRekening,
        })
        return nil, errors.New("no rekening tidak ditemukan")
    }

    account.Saldo += nominal
    if err := s.accountRepo.UpdateAccount(account); err != nil {
        utils.LogError("Failed to update account for Tabung", map[string]interface{}{
            "error":       err.Error(),
            "no_rekening": noRekening,
            "nominal":     nominal,
        })
        return nil, err
    }

    transaction := &models.Transaction{
        Type:      models.TransactionTypeTabung,
        Nominal:   nominal,
        AccountID: account.ID,
    }
    if err := s.accountRepo.CreateTransaction(transaction); err != nil {
        utils.LogError("Failed to create transaction for Tabung", map[string]interface{}{
            "error":       err.Error(),
            "no_rekening": noRekening,
            "nominal":     nominal,
        })
        return nil, err
    }

    utils.LogInfo("Tabung successful", map[string]interface{}{
        "no_rekening": noRekening,
        "saldo":       account.Saldo,
    })
    return account, nil
}

func (s *AccountService) Tarik(noRekening string, nominal float64) (*models.Account, error) {
    account, err := s.accountRepo.FindAccountByNoRekening(noRekening)
    if err != nil {
        utils.LogError("No rekening not found for Tarik", map[string]interface{}{
            "error":       err.Error(),
            "no_rekening": noRekening,
        })
        return nil, errors.New("no rekening tidak ditemukan")
    }

    if account.Saldo < nominal {
        utils.LogWarn("Insufficient balance for Tarik", map[string]interface{}{
            "no_rekening": noRekening,
            "saldo":       account.Saldo,
            "nominal":     nominal,
        })
        return nil, errors.New("saldo tidak cukup")
    }

    account.Saldo -= nominal
    if err := s.accountRepo.UpdateAccount(account); err != nil {
        utils.LogError("Failed to update account for Tarik", map[string]interface{}{
            "error":       err.Error(),
            "no_rekening": noRekening,
            "nominal":     nominal,
        })
        return nil, err
    }

    transaction := &models.Transaction{
        Type:      models.TransactionTypeTarik,
        Nominal:   nominal,
        AccountID: account.ID,
    }
    if err := s.accountRepo.CreateTransaction(transaction); err != nil {
        utils.LogError("Failed to create transaction for Tarik", map[string]interface{}{
            "error":       err.Error(),
            "no_rekening": noRekening,
            "nominal":     nominal,
        })
        return nil, err
    }

    utils.LogInfo("Tarik successful", map[string]interface{}{
        "no_rekening": noRekening,
        "saldo":       account.Saldo,
    })
    return account, nil
}

func (s *AccountService) CekSaldo(noRekening string) (*models.Account, error) {
    account, err := s.accountRepo.FindAccountByNoRekening(noRekening)
    if err != nil {
        utils.LogError("No rekening not found for CekSaldo", map[string]interface{}{
            "error":       err.Error(),
            "no_rekening": noRekening,
        })
        return nil, errors.New("no rekening tidak ditemukan")
    }

    utils.LogInfo("CekSaldo successful", map[string]interface{}{
        "no_rekening": noRekening,
        "saldo":       account.Saldo,
    })
    return account, nil
}