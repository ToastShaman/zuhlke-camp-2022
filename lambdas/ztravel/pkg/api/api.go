package api

import (
	"ztravel/pkg/aws"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/go-chi/jwtauth/v5"
)

func NewDefaultRouter() *chi.Mux {
	logger := httplog.NewLogger("httplog", httplog.Options{JSON: true})

	r := chi.NewRouter()
	r.Use(aws.RequestIdHeaders)
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.NoCache)
	r.Use(middleware.Heartbeat("/ping"))

	return r
}

func NewSecuredRouter(tokenAuth *jwtauth.JWTAuth) *chi.Mux {
	r := NewDefaultRouter()
	r.Use(jwtauth.Verifier(tokenAuth))
	r.Use(jwtauth.Authenticator)

	return r
}

func NewTokenAuth(secretFn aws.SecretSupplier) *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte(secretFn()), nil)
}
