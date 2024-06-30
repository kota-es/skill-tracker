package containers

import (
	"backend/services"
	"database/sql"
)

type ServiceContainer struct {
	User  services.UserService
	Auth  services.AuthService
	Skill services.SkillService
}

func NewServiceContainer(db *sql.DB) *ServiceContainer {
	return &ServiceContainer{
		User:  services.NewUserService(db),
		Auth:  services.NewAuthService(),
		Skill: services.NewSkillService(db),
	}
}
