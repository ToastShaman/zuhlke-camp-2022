package env

import (
	"os"

	"github.com/rs/zerolog/log"
)

func Required(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		log.Fatal().Msgf("%s is not set", name)
	}
	return value
}
