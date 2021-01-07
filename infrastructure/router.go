package infrastructure

import (
	"github.com/go-chi/chi"
	//v1User "go-api/domain/user/application/v1"
	"go-api/infrastructure/database"
	//"go-api/infrastructure/middleware"
	"net/http"
)

// Routes returns the API V1 Handler with configuration.
func Routes(conn *database.Data) http.Handler {
	router := chi.NewRouter()
	//ur := v1User.NewUserHandler(conn)
	//router.Mount("/users", routesUser(ur))

	return router
}
/*// routesUser returns user router with each endpoint.
func routesUser(handler *v1User.UserRouter) http.Handler {
	router := chi.NewRouter()
	router.With(middleware.AuthMiddleware).Get("/", handler.GetAllUserHandler)
	router.With(middleware.AuthMiddleware).Get("/{id}", handler.GetOneHandler)
	router.Post("/", handler.CreateHandler)
	router.With(middleware.AuthMiddleware).Put("/{id}", handler.UpdateHandler)
	return router
}*/