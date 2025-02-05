package handlers

import (
    "net/http"
    "strconv"
    "go-assessment/models"
    "github.com/labstack/echo/v4"
)

func Daftar(c echo.Context) error {
    var account models.Account
    if err := c.Bind(&account); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid payload"})
    }

    // Cek NIK atau NoHP sudah digunakan
    var existing models.Account
    if models.DB.Where("nik = ? OR no_hp = ?", account.NIK, account.NoHP).First(&existing).RowsAffected > 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "NIK atau NoHP sudah digunakan"})
    }

    // Generate NoRekening
    account.NoRekening = strconv.Itoa(int(account.ID))
    models.DB.Create(&account)

    return c.JSON(http.StatusOK, map[string]string{"no_rekening": account.NoRekening})
}

func Tabung(c echo.Context) error {
    var payload struct {
        NoRekening string  `json:"no_rekening"`
        Nominal    float64 `json:"nominal"`
    }
    if err := c.Bind(&payload); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid payload"})
    }

    var account models.Account
    if models.DB.Where("no_rekening = ?", payload.NoRekening).First(&account).RowsAffected == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "No Rekening tidak ditemukan"})
    }

    account.Saldo += payload.Nominal
    models.DB.Save(&account)

    return c.JSON(http.StatusOK, map[string]float64{"saldo": account.Saldo})
}

func Tarik(c echo.Context) error {
    var payload struct {
        NoRekening string  `json:"no_rekening"`
        Nominal    float64 `json:"nominal"`
    }
    if err := c.Bind(&payload); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid payload"})
    }

    var account models.Account
    if models.DB.Where("no_rekening = ?", payload.NoRekening).First(&account).RowsAffected == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "No Rekening tidak ditemukan"})
    }

    if account.Saldo < payload.Nominal {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Saldo tidak cukup"})
    }

    account.Saldo -= payload.Nominal
    models.DB.Save(&account)

    return c.JSON(http.StatusOK, map[string]float64{"saldo": account.Saldo})
}

func Saldo(c echo.Context) error {
    noRekening := c.Param("no_rekening")

    var account models.Account
    if models.DB.Where("no_rekening = ?", noRekening).First(&account).RowsAffected == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "No Rekening tidak ditemukan"})
    }

    return c.JSON(http.StatusOK, map[string]float64{"saldo": account.Saldo})
}