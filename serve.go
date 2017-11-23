package gres

import (
	"github.com/Sirupsen/logrus"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// NewRouter instanciate the router with some middleware
func NewRouter(logger *logrus.Logger) *chi.Mux {

	r := chi.NewRouter()
	r.Use(RequestID())
	r.Use(Header("Content-Type", "application/json"))
	r.Use(middleware.RealIP)
	r.Use(Logger(logger))

	return r
}
