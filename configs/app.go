package configs

import (
	"os"
	"sync"
)

type app struct {
	Name         string
	Path         string
	DatabasePath string
	Env          string
	Local        bool
	DEBUG        bool
	LOG_STACK     string
}

var once sync.Once
var App *app

func init() {
	once.Do(func() {
		var name, env string
		var debug bool
		if name = os.Getenv("APP_NAME"); name == "" {
			name = "Soro"
		}
		if env = os.Getenv("APP_ENV"); env == "" {
			env = "production"
		}
		if debug_ := os.Getenv("APP_DEBUG"); debug_ == "" {
			debug = false
		}

		App = &app{
			Name:         name,
			Path:         __PATH__,
			DatabasePath: __DB_PATH__,
			Env:          env,
			Local:        env != "production",
			DEBUG:        debug,
			LOG_STACK:  os.Getenv("LOG_STACK"),
		}
	})
}
