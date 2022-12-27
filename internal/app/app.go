package app

import (
	"github.com/nats-io/stan.go"
	"github.com/pokrovsky-io/msg-store/config"
	"github.com/pokrovsky-io/msg-store/internal/repo/cache"
	"github.com/pokrovsky-io/msg-store/internal/repo/psql"
	"github.com/pokrovsky-io/msg-store/internal/transport/nats"
	"github.com/pokrovsky-io/msg-store/internal/usecase"
	"log"
)

func Run(cfg *config.Config) {
	sc, err := stan.Connect(cfg.NATS.ClusterID, cfg.NATS.ClientID)
	if err != nil {
		// TODO: Обработать ошибку
		log.Fatal(err)
	}
	defer sc.Close()

	ch := cache.New()
	pg := psql.New()

	uc := usecase.New(ch, pg)

	stn := nats.New(sc, uc)
	stn.Subscribe(cfg.NATS.Subject)
}
