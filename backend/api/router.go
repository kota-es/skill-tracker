package api

import (
	"backend/containers"
	"backend/controllers"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(db *sql.DB) *chi.Mux {
	serviceContainer := containers.NewServiceContainer(db)
	userController := controllers.NewUserController(serviceContainer)
	authController := controllers.NewAuthController(serviceContainer)
	skillController := controllers.NewSkillController(serviceContainer)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	r.Post("/users", userController.PostUser)
	r.Post("/login", authController.Login)
	r.Get("/users/me", authController.Me)
	r.Get("/skills", skillController.GetAllSkills)
	r.Get("/users/{user_id}/skills", skillController.GetUserSkills)
	r.Post("/users/skills", skillController.PostUserSkill)
	r.Post("/admin/skills", skillController.PostSkill)
	r.Get("/skills/categories", skillController.GetSkillCategories)

	return r
}
