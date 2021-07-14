package configs

import (
	"fmt"
	"os"
	"sync"
)

type app struct{
	Name string
	Path string
	DatabasePath string
	Env string
	Local bool
}

var once sync.Once
var App *app

func init() {
	once.Do(func() {
		var name string
		var env string
		if name =  os.Getenv("APP_NAME"); name == "" {name = "Soro"}
		if env = os.Getenv("APP_ENV"); env == "" {env = "production"}
		
		fmt.Println(os.Getenv("APP_NAME") + " From configs")
		App = &app{
			Name: name,
			Path: __PATH__,
			DatabasePath: __DB_PATH__,
			Env: env,
			Local: env != "production",
		}
	})
}