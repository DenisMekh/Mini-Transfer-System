package repo

import (
	"context"

	"github.com/DenisMekh/mini-transfer-system/account-svc/internal/domain"
)

// AccountRepository for account implementation
type AccountRepository interface {
	Create(ctx context.Context, account *domain.Account) error
	GetByID(ctx context.Context, id string) (*domain.Account, error)
	UpdateBalance(ctx context.Context, id string, amount int64) (*domain.Account, error)
}
