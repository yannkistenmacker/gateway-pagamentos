package service

import "github.com/devfullcycle/imersao22/go-gateway/internal/domain"

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}
