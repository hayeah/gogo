package gogo

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

func Main(run func() error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("dotenv load")
	}

	loglevel, found := os.LookupEnv("LOG_LEVEL")
	if found {
		level, err := zerolog.ParseLevel(loglevel)
		if err != nil {
			log.Warn().Err(err).Msg("LOG_LEVEL env")
		} else {
			zerolog.SetGlobalLevel(level)
		}
	}

	err = run()
	if err != nil {
		log.Fatal().Err(err).Msg("run")
	}
}