package hello

import (
	"fmt"
	"net/http"
	"ztravel/pkg/api"
	"ztravel/pkg/aws"
	"ztravel/pkg/pigtail"

	"github.com/rs/zerolog/log"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func NewHelloApi(storage aws.Storage, queue aws.Queue) *chi.Mux {
	r := api.NewDefaultRouter()
	r.Get("/", GetHelloHandler(storage, queue))
	return r
}

func GetHelloHandler(storage aws.Storage, queue aws.Queue) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := storage.PersistJSON(r.Context(), fmt.Sprintf("%s.json", uuid.NewString()), "{}")
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to persis JSON")
		}

		_, err = queue.Send(r.Context(), &pigtail.MyEvent{Name: uuid.NewString()})
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to queue message")
		}

		w.Write([]byte("hello"))
	}
}
