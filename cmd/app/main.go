package main

import (
	"github.com/pokrovsky-io/msg-store/config"
	"github.com/pokrovsky-io/msg-store/internal/app"
)

func main() {
	cfg, _ := config.NewConfig()

	app.Run(cfg)
}
