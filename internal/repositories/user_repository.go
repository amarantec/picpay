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
		`INSERT INTO users (first_name, last_name, document, email, password, user_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, first_name, last_name, document, email, user_id`,
		user.FirstName, user.LastName, user.Document, user.Email, hashedPassword, user.UserId).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Document, &user.Email, &user.UserId)
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
		`
		SELECT 
		COALESCE (balance, 0.0) AS balance 
		FROM users WHERE id=$1
		`, id).Scan(&user.Balance)
	if err != nil {
		return 0, err
	}

	return user.Balance, nil
}

func (r *RepositoryPostgres) Transfer(ctx context.Context, senderId int64, receiverId int64, value float64) error {
	tx, err := r.Conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	var user = models.User{Id: senderId}
	err = tx.QueryRow(
		ctx,
		`SELECT balance FROM users WHERE id=$1`, senderId).Scan(&user.Balance)
	if err != nil {
		return err
	}

	if user.Balance < value {
		return err
	}

	_, err = tx.Exec(
		ctx,
		`UPDATE user SET balance = balance - $1 WHERE id = $2`, value, senderId)
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		ctx,
		`UPDATE user SET balance = balance + $1 WHERE id = $2`, value, receiverId)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}
