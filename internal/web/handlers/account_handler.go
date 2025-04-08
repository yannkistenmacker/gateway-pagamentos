package handlers

import "github.com/devfullcycle/imersao22/go-gateway/internal/service"

type AccountHandler struct {
	accountService *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}
