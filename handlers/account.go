package handlers

import (
    "net/http"
    "fmt"
    "go-assessment/models"
    "github.com/labstack/echo/v4"
)

func Daftar(c echo.Context) error {
    var account models.Account
    if err := c.Bind(&account); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid payload"})
    }

    // Debug - cetak nilai yang di-bind
    fmt.Printf("Received account: %+v\n", account)

    // Validasi input tidak boleh kosong
    if account.NIK == "" || account.NoHP == "" || account.Nama == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "remark": "NIK, NoHP, dan Nama harus diisi",
            "debug": fmt.Sprintf("NIK: '%s', NoHP: '%s', Nama: '%s'", account.NIK, account.NoHP, account.Nama),
        })
    }

    // Create account terlebih dahulu untuk mendapatkan ID
    if err := models.DB.Create(&account).Error; err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "NIK atau NoHP sudah digunakan"})
    }

    // Update NoRekening berdasarkan ID yang sudah di-generate
    account.NoRekening = fmt.Sprintf("%d", account.ID)
    if err := models.DB.Save(&account).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Gagal update no rekening"})
    }

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

    // Validasi nominal
    if payload.Nominal <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Nominal harus lebih dari 0"})
    }

    var account models.Account
    if err := models.DB.Where("no_rekening = ? AND deleted_at IS NULL", payload.NoRekening).First(&account).Error; err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "No Rekening tidak ditemukan"})
    }

    account.Saldo += payload.Nominal
    if err := models.DB.Save(&account).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Gagal update saldo"})
    }

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

    // Validasi nominal
    if payload.Nominal <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Nominal harus lebih dari 0"})
    }

    var account models.Account
    if err := models.DB.Where("no_rekening = ? AND deleted_at IS NULL", payload.NoRekening).First(&account).Error; err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "No Rekening tidak ditemukan"})
    }

    if account.Saldo < payload.Nominal {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Saldo tidak cukup"})
    }

    account.Saldo -= payload.Nominal
    if err := models.DB.Save(&account).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"remark": "Gagal update saldo"})
    }

    return c.JSON(http.StatusOK, map[string]float64{"saldo": account.Saldo})
}

func Saldo(c echo.Context) error {
    noRekening := c.Param("no_rekening")

    var account models.Account
    if err := models.DB.Where("no_rekening = ? AND deleted_at IS NULL", noRekening).First(&account).Error; err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "No Rekening tidak ditemukan"})
    }

    return c.JSON(http.StatusOK, map[string]float64{"saldo": account.Saldo})
}