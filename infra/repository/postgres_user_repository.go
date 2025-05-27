package repository

import (
	"context"
	"database/sql"
	"errors"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type IUserRepository interface {
	GetUser(ctx context.Context, id string) (*User, error)
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) GetUser(ctx context.Context, id string) (*User, error) {
	var user User
	query := "SELECT id, first_name, last_name, email, phone FROM users WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &user, nil
}
