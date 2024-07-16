package repositories

import (
	"context"
	"errors"

	"github.com/amarantec/picpay/internal/models"
	"github.com/amarantec/picpay/internal/utils"
)

func (r *RepositoryPostgres) SaveUser(ctx context.Context, user models.User) (models.User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}

	err = r.Conn.QueryRow(
		ctx,
		`INSERT INTO users (first_name, last_name, document, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id, first_name, last_name, cpf, email`,
		user.FirstName, user.LastName, user.Document, user.Email, hashedPassword).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Document, &user.Email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *RepositoryPostgres) ValidateUserCredentials(ctx context.Context, user models.User) error {
	var retriviedPassword string

	err := r.Conn.QueryRow(
		ctx,
		`SELECT id, first_name, last_name, document, password FROM users WHERE email=$1`, user.Email).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Document, &retriviedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPassword(user.Password, retriviedPassword)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}
	return nil
}

func (r *RepositoryPostgres) GetTotalBalanceAccount(ctx context.Context, id int64) (float64, error) {
	var user = models.User{Id: id}
	err := r.Conn.QueryRow(
		ctx,
		"SELECT balance FROM users WHERE id=$1", id).Scan(
		&user.Id, &user.Balance)
	if err != nil {
		return 0, err
	}

	return user.Balance, nil
}
