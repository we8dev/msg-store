package app

import (
	"github.com/pokrovsky-io/msg-store/config"
	"github.com/pokrovsky-io/msg-store/internal/transport"
)

// TODO: Закрыть STAN Connection
func Run(cfg *config.Config) {
	stan := transport.NewSTAN(cfg.NATS)
	stan.Subscribe()
}
