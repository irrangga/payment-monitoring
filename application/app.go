package application

import (
	"payment-monitoring/config/env"
	"payment-monitoring/router"

	g "github.com/incubus8/go/pkg/gin"
	"github.com/rs/zerolog/log"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		return
	}
}

func StartApp() {
	addr := env.Config.Environment + ":" + env.Config.AppPort
	conf := g.Config{
		ListenAddr: addr,
		Handler:    router.NewRouter(),
		OnStarting: func() {
			log.Info().Msg("Your service is up and running at " + addr)
		},
	}

	g.Run(conf)
}
