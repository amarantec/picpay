package handlers

import (
	"github.com/amarantec/picpay/internal/database"
	"github.com/amarantec/picpay/internal/repositories"
	"github.com/amarantec/picpay/internal/services"
)

var service services.Service

func Configure() {
	service = services.Service{
		Repository: &repositories.RepositoryPostgres{
			Conn: database.Conn,
		},
	}
}
