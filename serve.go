package gres

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// New instanciate the http server and return a channel
func NewRouter(host string, httpPort int, logger *logrus.Logger) *chi.Mux {

	r := chi.NewRouter()
	r.Use(RequestID())
	r.Use(Header("Content-Type", "application/json"))
	r.Use(middleware.RealIP)
	r.Use(Logger(logger))

	return r
}