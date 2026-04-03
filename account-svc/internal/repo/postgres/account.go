package postgres

import (
	"context"
	"errors"

	"github.com/DenisMekh/mini-transfer-system/account-svc/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRepo struct {
	pool *pgxpool.Pool
}

func NewAccountRepo(pool *pgxpool.Pool) *AccountRepo {
	return &AccountRepo{pool: pool}
}

func (r *AccountRepo) Create(ctx context.Context, account *domain.Account) error {
	const query = `
INSERT INTO accounts (user_id, name) VALUES ($1, $2) RETURNING account_id, created_at;`
	err := r.pool.QueryRow(ctx, query, account.UserID, account.Name).Scan(&account.AccountID, &account.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *AccountRepo) GetByID(ctx context.Context, id string) (*domain.Account, error) {
	const query = `
SELECT account_id, user_id, name, balance, created_at FROM accounts WHERE account_id = $1;`
	var account domain.Account
	err := r.pool.QueryRow(ctx, query, id).Scan(&account.AccountID, &account.UserID, &account.Name, &account.Balance, &account.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &account, nil
}

func (r *AccountRepo) UpdateBalance(ctx context.Context, id string, amount int64) (*domain.Account, error) {
	const query = `
UPDATE accounts SET balance = balance + $1 WHERE account_id = $2 RETURNING account_id, user_id, name, balance, created_at;`
	var account domain.Account
	err := r.pool.QueryRow(ctx, query, amount, id).Scan(&account.AccountID, &account.UserID, &account.Name, &account.Balance, &account.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &account, nil

}
