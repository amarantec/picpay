package services

import (
	"errors"

	"github.com/amarantec/picpay/internal/repositories"
)

type Service struct {
	Repository repositories.Repository
}

var ErrUserFirstNameEmpty = errors.New("user first name is empty")
var ErrUserLastNameEmpty = errors.New("user last name is empty")
var ErrUserDocumentEmpty = errors.New("user cpf is empty")
var ErrUserEmailEmpty = errors.New("user email is empty")
var ErrUserPasswordEmpty = errors.New("user password is empty")
var ErrSenderIdEmpty = errors.New("senderId is empty")
var ErrReceiverIdEmpty = errors.New("receiveId is empty")
var ErrValueEmpty = errors.New("value is empty")
