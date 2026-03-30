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

func (r *AccountRepo) CreateAccount(ctx context.Context, account *domain.Account) error {
	const query = `
INSERT INTO accounts (name) VALUES ($1) RETURNING account_id, created_at;`
	err := r.pool.QueryRow(ctx, query, account.Name).Scan(&account.AccountID, &account.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *AccountRepo) GetByID(ctx context.Context, id string) (*domain.Account, error) {
	const query = `
SELECT user_id, name, balance, created_at FROM accounts WHERE account_id = $1;`
	var account domain.Account
	err := r.pool.QueryRow(ctx, query, id).Scan(&account.UserID, &account.Name, &account.Balance, &account.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrNotFound
	}
	return &account, nil
}

func (r *AccountRepo) UpdateBalance(ctx context.Context, id string, amount int64) error {
	const query = `
UPDATE accounts SET balance = balance + $1 WHERE account_id = $2;`
	tag, err := r.pool.Exec(ctx, query, amount, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil

}
