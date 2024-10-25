package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, you've hit %s\n", r.URL.Path)
	if err != nil {
		return
	}
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Service RUN good!")
	log.Error().Msg("UPDATE Last one")
	log.Info().Msg("Test One two")
	log.Error().Msg("Docker Uploaded")

	http.HandleFunc("/", handler)

	go func() {
		log.Info().Msg("Starting server on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	for {
		log.Info().Msg("Service is running...")
		time.Sleep(10 * time.Second)
	}
}
