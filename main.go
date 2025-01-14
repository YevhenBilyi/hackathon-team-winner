package main

import (
	"log"

	"github.com/noona-hq/noonaNordar/config"
	"github.com/noona-hq/noonaNordar/logger"
	"github.com/noona-hq/noonaNordar/server"
)
// main function

func main() {
	cfg := new(server.Config)
	err := config.Process(cfg)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := logger.FromConfig(cfg.Logger)
	if err != nil {
		log.Fatal(err)
	}

	srv, err := server.New(*cfg, *logger)
	if err != nil {
		logger.Fatal(err)
	}

	if err := srv.Serve(); err != nil {
		log.Fatal(err)
	}
}
