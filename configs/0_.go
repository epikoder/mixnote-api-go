package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var __PATH__, _ = os.Getwd()
var __DB_PATH__ = __PATH__ + "/bin/database.sqlite"
func init() {
	envPath := __PATH__ + "/.env"
	if err := godotenv.Load(envPath); err != nil {
		fmt.Printf("Could not load .env file in %s", envPath)
	}
}