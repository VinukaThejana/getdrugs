package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	env "github.com/VinukaThejana/getdrugs/internal/config"
	"github.com/VinukaThejana/getdrugs/internal/enums"
	"github.com/VinukaThejana/getdrugs/internal/router"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	e = &env.Env{}
)

func init() {
	e.Load()

	if e.Env == string(enums.Dev) {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out: os.Stderr,
		})
	}
}

func main() {
	server := &http.Server{
		Addr:    ":" + e.Port,
		Handler: router.Init(e),
	}

	go func() {
		log.Info().Msgf("starting server on port %s", e.Port)
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Fatal().Err(err).Msg("server failed to start")
			}
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch

	log.Info().Msg("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stopped := make(chan struct{})
	go func() {
		if err := server.Shutdown(ctx); err != nil {
			log.Fatal().Err(err).Msg("failed to shutdown the server")
		}
		close(stopped)
	}()

	select {
	case <-ctx.Done():
		log.Error().Msg("server shutdown failed, forcing shutdown")
	case <-stopped:
		log.Info().Msg("server shutdown successfully")
	}
}
