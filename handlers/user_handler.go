package handlers

import (
    "net/http"
    "go-assessment/services"
    "go-assessment/utils"
    "github.com/labstack/echo/v4"
)

type UserHandler struct {
    userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
    var request struct {
        Nama  string `json:"nama"`
        NIK   string `json:"nik"`
        NoHP  string `json:"no_hp"`
    }
    if err := c.Bind(&request); err != nil {
        utils.LogError("Invalid payload", map[string]interface{}{
            "error": err.Error(),
        })
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid payload"})
    }

    user, err := h.userService.RegisterUser(request.Nama, request.NIK, request.NoHP)
    if err != nil {
        utils.LogError("Failed to register user", map[string]interface{}{
            "error": err.Error(),
            "nik":   request.NIK,
            "no_hp": request.NoHP,
        })
        return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
    }

    utils.LogInfo("User registered successfully", map[string]interface{}{
        "user_id": user.ID,
        "nama":    user.Nama,
    })
    return c.JSON(http.StatusOK, user)
}