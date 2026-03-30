package service

import (
	"context"

	"github.com/DenisMekh/mini-transfer-system/account-svc/internal/domain"
	"github.com/DenisMekh/mini-transfer-system/account-svc/internal/repo"
)

type AccountServ struct {
	repo repo.AccountRepository
}

func NewAccountService(repo repo.AccountRepository) *AccountServ {
	return &AccountServ{repo: repo}
}

func (s *AccountServ) Create(ctx context.Context, account *domain.Account) error {
	if account.Name == "" {
		return domain.ErrRequiredAccountName
	}
	return s.repo.Create(ctx, account)
}
func (s *AccountServ) GetByID(ctx context.Context, id string) (*domain.Account, error) {

	return s.repo.GetByID(ctx, id)
}
func (s *AccountServ) UpdateBalance(ctx context.Context, id string, amount int64) error {
	return s.repo.UpdateBalance(ctx, id, amount)
}
