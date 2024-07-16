package repositories

import (
	"context"

	"github.com/amarantec/picpay/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryPostgres struct {
	Conn *pgxpool.Pool
}

type Repository interface {
	SaveUser(ctx context.Context, user models.User) (models.User, error)
	ValidateUserCredentials(ctx context.Context, user models.User) error
}
