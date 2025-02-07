package handlers

import (
    "net/http"
    "go-assessment/services"
    "go-assessment/utils"
    "github.com/labstack/echo/v4"
)

type AccountHandler struct {
    accountService *services.AccountService
}

func NewAccountHandler(accountService *services.AccountService) *AccountHandler {
    return &AccountHandler{accountService: accountService}
}

func (h *AccountHandler) Tabung(c echo.Context) error {
    var request struct {
        NoRekening string  `json:"no_rekening"`
        Nominal    float64 `json:"nominal"`
    }
    if err := c.Bind(&request); err != nil {
        utils.LogError("Invalid payload for Tabung", map[string]interface{}{
            "error": err.Error(),
        })
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid payload"})
    }

    account, err := h.accountService.Tabung(request.NoRekening, request.Nominal)
    if err != nil {
        utils.LogError("Failed to process Tabung", map[string]interface{}{
            "error":       err.Error(),
            "no_rekening": request.NoRekening,
            "nominal":     request.Nominal,
        })
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
    }

    utils.LogInfo("Tabung successful", map[string]interface{}{
        "no_rekening": request.NoRekening,
        "saldo":       account.Saldo,
    })
    return c.JSON(http.StatusOK, map[string]float64{"saldo": account.Saldo})
}

func (h *AccountHandler) Tarik(c echo.Context) error {
    var request struct {
        NoRekening string  `json:"no_rekening"`
        Nominal    float64 `json:"nominal"`
    }
    if err := c.Bind(&request); err != nil {
        utils.LogError("Invalid payload for Tarik", map[string]interface{}{
            "error": err.Error(),
        })
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid payload"})
    }

    account, err := h.accountService.Tarik(request.NoRekening, request.Nominal)
    if err != nil {
        utils.LogError("Failed to process Tarik", map[string]interface{}{
            "error":       err.Error(),
            "no_rekening": request.NoRekening,
            "nominal":     request.Nominal,
        })
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
    }

    utils.LogInfo("Tarik successful", map[string]interface{}{
        "no_rekening": request.NoRekening,
        "saldo":       account.Saldo,
    })
    return c.JSON(http.StatusOK, map[string]float64{"saldo": account.Saldo})
}

func (h *AccountHandler) CekSaldo(c echo.Context) error {
    noRekening := c.Param("no_rekening")

    account, err := h.accountService.CekSaldo(noRekening)
    if err != nil {
        utils.LogError("Failed to process CekSaldo", map[string]interface{}{
            "error":       err.Error(),
            "no_rekening": noRekening,
        })
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
    }

    utils.LogInfo("CekSaldo successful", map[string]interface{}{
        "no_rekening": noRekening,
        "saldo":       account.Saldo,
    })
    return c.JSON(http.StatusOK, map[string]float64{"saldo": account.Saldo})
}