package app

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

var __PATH__ , _ = os.Getwd()
var __DB_PATH__ = __PATH__ + "/bin/database.sqlite"

func init() {
	loadEnv()
}

func loadEnv() {
	envPath := __PATH__ + "/.env"
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Printf("Could not load .env file in %s", envPath)
	}
}

func RegisterBindings() {
	
}

func AppPath() string {
	return __PATH__
}

func DatabasePath() string {
	return __DB_PATH__
}

