package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"github.com/pokrovsky-io/msg-store/config"
	"github.com/pokrovsky-io/msg-store/internal/repo/cache"
	"github.com/pokrovsky-io/msg-store/internal/repo/psql"
	"github.com/pokrovsky-io/msg-store/internal/transport/nats"
	"github.com/pokrovsky-io/msg-store/internal/transport/rest"
	"github.com/pokrovsky-io/msg-store/internal/usecase"
	"github.com/pokrovsky-io/msg-store/pkg/server"
	"log"
	"sync"
)

func Run(cfg *config.Config) {
	// Repository
	ch := cache.New()
	pg := psql.New()

	// Use case
	uc := usecase.New(ch, pg)

	// STAN
	sc, err := stan.Connect(cfg.NATS.ClusterID, cfg.NATS.ClientID)
	if err != nil {
		// TODO: Обработать ошибку
		log.Fatal(err)
	}
	defer sc.Close()
	stn := nats.New(sc, uc)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go stn.Subscribe(wg, cfg.NATS.Subject)

	// HTTP Server
	handler := gin.New()
	rest.NewRouter(handler, uc)
	httpServer := server.New(handler, server.Port(cfg.HTTP.Port))

	// TODO: Обработать ошибку
	httpServer.Run()

	wg.Wait()

	fmt.Println("DONE")
}
