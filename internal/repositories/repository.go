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
	GetTotalBalanceAccount(ctx context.Context, id int64) (float64, error)
	Transfer2(ctx context.Context, senderId int64, receiverId int64, value float64) error
}
