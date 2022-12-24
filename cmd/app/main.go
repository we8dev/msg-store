package main

import (
	"github.com/pokrovsky-io/msg-store/config"
	"github.com/pokrovsky-io/msg-store/internal/app"
)

func main() {
	natsCfg := config.NATS{
		ClusterID: "test-cluster",
		ClientID:  "test-client-1",
		Subject:   "foo",
	}

	cfg := config.Config{
		natsCfg,
	}

	app.Run(&cfg)
}
