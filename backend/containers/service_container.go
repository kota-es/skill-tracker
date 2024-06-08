package containers

import (
	"backend/services"
	"database/sql"
)

type ServiceContainer struct {
	User services.UserService
}

func NewServiceContainer(db *sql.DB) *ServiceContainer {
	return &ServiceContainer{
		User: services.NewUserService(db),
	}
}
