package services

import (
	"context"

	"github.com/amarantec/picpay/internal/models"
)

func (s Service) SaveUser(ctx context.Context, user models.User) (models.User, error) {
	if user.FirstName == "" {
		return models.User{}, ErrUserFirstNameEmpty
	}
	if user.LastName == "" {
		return models.User{}, ErrUserLastNameEmpty
	}
	if user.Document == "" {
		return models.User{}, ErrUserDocumentEmpty
	}
	if user.Email == "" {
		return models.User{}, ErrUserEmailEmpty
	}
	if user.Password == "" {
		return models.User{}, ErrUserPasswordEmpty
	}

	return s.Repository.SaveUser(ctx, user)
}

func (s Service) ValidateUserCredentials(ctx context.Context, user models.User) error {
	if user.Email == "" {
		return ErrUserEmailEmpty
	}
	if user.Password == "" {
		return ErrUserPasswordEmpty
	}
	return s.Repository.ValidateUserCredentials(ctx, user)
}

func (s Service) GetTotalBalanceAccount(ctx context.Context, id int64) (float64, error) {
	return s.Repository.GetTotalBalanceAccount(ctx, id)
}

func (s Service) Transfer(ctx context.Context, senderId int64, receiverId int64, value float64) error {
	if senderId == 0 {
		return ErrSenderIdEmpty
	}
	if receiverId == 0 {
		return ErrReceiverIdEmpty
	}
	if value == 0 {
		return ErrValueEmpty
	}
	return s.Repository.Transfer(ctx, senderId, receiverId, value)
}
