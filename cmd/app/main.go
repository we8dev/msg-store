package main

import (
	"github.com/pokrovsky-io/msgstore/config"
	"github.com/pokrovsky-io/msgstore/internal/app"
)

func main() {
	cfg, _ := config.NewConfig()

	app.Run(cfg)
}
