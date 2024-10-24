package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Service RUN good!")
	log.Error().Msg("UPDATE Last one")
	log.Info().Msg("Test One two")
	log.Error().Msg("Docker Uploaded")

	fmt.Println()

	for {
		log.Info().Msg("Service is running...")
		time.Sleep(10 * time.Second)
	}

}
