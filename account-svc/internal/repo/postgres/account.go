package postgres

import (
	"context"

	"github.com/DenisMekh/mini-transfer-system/account-svc/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRepo struct {
	pool *pgxpool.Pool
}

func NewAccountRepo(pool *pgxpool.Pool) *AccountRepo {
	return &AccountRepo{pool: pool}
}

func (repo *AccountRepo) CreateAccount(ctx context.Context, account *domain.Account) error {
	//
	return nil
}
