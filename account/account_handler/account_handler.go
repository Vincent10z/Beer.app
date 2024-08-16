// accounts/handler/account_handler.go
package account_handler

import (
	"Beer.app/account/repository"
	"Beer.app/account/service"
	"Beer.app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func NewAccountHandler(service service.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

func AccountRouter(e *echo.Echo, db *gorm.DB) {
	accountRepo := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepo)
	accountHandler := NewAccountHandler(accountService)

	e.GET("/accounts/:id", accountHandler.GetAccount)
	e.POST("/accounts", accountHandler.CreateAccount)
}

func (h *AccountHandler) GetAccount(c echo.Context) error {
	account, err := h.service.GetAccount(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) CreateAccount(c echo.Context) error {
	account := new(models.Account)
	if err := c.Bind(account); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.service.CreateAccount(account); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	
	return c.JSON(http.StatusCreated, account)
}
