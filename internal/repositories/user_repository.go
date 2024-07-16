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
		`INSERT INTO users (first_name, last_name, cpf, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id, first_name, last_name, cpf, email`,
		user.FirstName, user.LastName, user.CPF, user.Email, hashedPassword).Scan(&user.Id, &user.FirstName, &user.LastName, &user.CPF, &user.Email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *RepositoryPostgres) ValidateUserCredentials(ctx context.Context, user models.User) error {
	var retriviedPassword string

	err := r.Conn.QueryRow(
		ctx,
		`SELECT id, first_name, last_name, cpf FROM users WHERE email=$1`, user.Email).Scan(&user.Id, &user.FirstName, &user.LastName, &user.CPF)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPassword(user.Password, retriviedPassword)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}
	return nil
}
