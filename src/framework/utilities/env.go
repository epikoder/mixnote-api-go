package utilities

import (
	"os"

	"github.com/joho/godotenv"
)

type env struct{}

var (
	Env     *env
	enVars  map[string]string
	file    *os.File
	envPath string
	err     error
)

func init() {
	once.Do(func() {
		Env = new(env)
	})
}

func (*env) Write(map_ map[string]string) error {
	envPath = os.Getenv("ENV_PATH")
	if file, err = os.Open(envPath); err != nil {
		Console.Fatal(err)
	}

	if enVars, err = godotenv.Parse(file); err != nil {
		Console.Fatal(err)
	}

	for k, v := range map_ {
		enVars[k] = v
	}

	if err = godotenv.Write(enVars, envPath); err != nil {
		Console.Fatal(err)
	}
	for k, v := range map_ {
		os.Setenv(k, v)
	}

	return err
}
