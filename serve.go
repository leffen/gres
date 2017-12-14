package gres

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// NewRouter instanciate the router with some middleware
func NewRouter(logger logrus.FieldLogger) *chi.Mux {

	r := chi.NewRouter()
	r.Use(RequestID())
	r.Use(Header("Content-Type", "application/json"))
	r.Use(middleware.RealIP)
	r.Use(Logger(logger))

	return r
}
