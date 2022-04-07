package world

import (
	"net/http"
	"ztravel/pkg/api"
	"ztravel/pkg/clock"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func NewWorldApi(tokenAuth *jwtauth.JWTAuth, signer api.Signer, clock clock.Clockwork) *chi.Mux {
	rs := api.NewHttpResponseSigner(signer, clock)

	r := api.NewSecuredRouter(tokenAuth)
	r.Get("/", rs.WithSignedResponses(GetWorldHandler()))

	return r
}

func GetWorldHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("world"))
	})
}
