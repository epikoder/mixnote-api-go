package main

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mixnote/mixnote-api-go/cmd/helpers"
	"github.com/mixnote/mixnote-api-go/cmd/migrate"
	"github.com/mixnote/mixnote-api-go/configs"
	app "github.com/mixnote/mixnote-api-go/src"
	"github.com/mixnote/mixnote-api-go/src/framework/utilities"
	"gorm.io/gorm"
)

var DB *gorm.DB
var cli = helpers.Cli()

var banner string
var message string

func init() {
	banner = `
|\     /|  | \    /  |\      |  /------\ --------   /---
| \   / |  |   \_/   |  \    |  | |--| |     |     /____|
|  \/   |  |   / \   |    \  |  | |__| |     |    /-----|
|       |  |  /    \ |      \|  \______/     |   /_______

` + configs.App.Name + ` is a audio platform for streaming, podcast and radio
	`
	message = `
usage: 
` + strings.ToLower(configs.App.Name) + ` <command> [arguments]

Commands:
	serve		Start the server
	test		Run test
	migrate		Run migration
	help		Show help message
	version		Show version information
	
Use "` + strings.ToLower(configs.App.Name) + ` help <command>" for more information about a command.
	`
	fmt.Println(banner)
}

func helpServe() {
	fmt.Println(`
Usage:
` + strings.ToLower(configs.App.Name) + ` serve [argument1] [argument2]

arguments: 
	--host=HOST		Server host default is localhost [127.0.0.1]
	--port=PORT		Server port default is :8080

	`)
}

func helpTest()    {}
func helpMigrate() {}
func helpVersion() {}

func main() {
	// Setup
	if sec := os.Getenv("JWT_SECRET"); sec == "" {
		var err error
		var secret *ecdsa.PrivateKey

		utilities.Console.Warn("JWT secret not found.. creating secret")
		if secret, err = utilities.Crypto.KeyGen(); err != nil {
			utilities.Console.Fatal(err)
		}
		privKey, _, err := utilities.Crypto.PemKeyPair(secret)
		if err != nil {
			utilities.Console.Fatal(err)
		}

		utilities.Env.Write(map[string]string{
			"JWT_SECRET": string(privKey),
		})
		utilities.Console.Success("JWT secret created successfully")
	}

	cmd := cli.Argument()
	switch cmd {
	case "migrate":
		migrate.Migrate()
		os.Exit(0)

	case "serve":
		app.StartServer(cli.Option("host"), (func() int {
			_port := cli.Option("port")
			if _port != "" {
				port, err := strconv.Atoi(_port)
				if err != nil {
					panic("Invalid port address")
				}
				return port
			}
			return 8080
		})())

	default:
		if cmd != "help" && cli.ArgsLenght() > 1 {
			fmt.Println("Unrecognized command : " + os.Args[1])
		}
		switch cli.Option("opt1") {
		case "serve":
			helpServe()
		case "test":
			helpTest()
		case "migrate":
			helpMigrate()
		case "version":
			helpVersion()
		default:
			fmt.Println(message)
		}
		os.Exit(0)
	}
}
