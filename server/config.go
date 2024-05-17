package server

import (
	"github.com/noona-hq/noonaNordar/db"
	"github.com/noona-hq/noonaNordar/logger"
	"github.com/noona-hq/noonaNordar/services/noona"
)

type Config struct {
	Noona  noona.Config
	Logger logger.Config
	DB     db.Config
	// Store can either be memory or mongodb
	Store string `default:"mongodb"`
}
