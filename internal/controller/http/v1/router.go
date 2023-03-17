package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// External
	"github.com/BogdanStaziyev/shop-test/pkg/logger"
	"github.com/BogdanStaziyev/shop-test/pkg/validators"

	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/controller/http/middlewares"
	"github.com/BogdanStaziyev/shop-test/internal/controller/http/responses"
	"github.com/BogdanStaziyev/shop-test/internal/service"
)

// Router initialize new CHI router
func Router(router *chi.Mux, service service.Services, l logger.Interface, user, password string) http.Handler {
	router.Use(middleware.RedirectSlashes, middleware.Logger)

	// Initialize a validator that validates data in requests using tags
	validator := validators.NewValidator()

	router.Route("/api", func(apiRouter chi.Router) {
		// Health
		apiRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())
			healthRouter.Handle("/*", NotFoundJSON())
		})
		apiRouter.Route("/v1", func(apiRouter chi.Router) {
			// Private routes
			// Only admin can send requests using basic auth (user - password)
			apiRouter.With(middlewares.CheckAuth(password, user)).Group(func(apiRouter chi.Router) {
				newCustomerHandler(apiRouter, service.Customer, validator, l)
			})
			apiRouter.Handle("/*", NotFoundJSON())
		})
		apiRouter.Handle("/*", NotFoundJSON())
	})
	return router
}

// NotFoundJSON returns a message that the page was not found and the status code 404
func NotFoundJSON() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responses.ErrorResponse(w, http.StatusNotFound, "Resource Not Found")
	}
}

// PingHandler can check the website's performance by pinging it
func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responses.Response(w, http.StatusOK, "OK")
	}
}
