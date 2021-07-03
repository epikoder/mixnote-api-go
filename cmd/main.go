package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mixnote/mixnote-api-go/cmd/helper"
	"github.com/mixnote/mixnote-api-go/cmd/migrate"
	app "github.com/mixnote/mixnote-api-go/src"
	"github.com/mixnote/mixnote-api-go/src/framework/database"

	"gorm.io/gorm"
)

var DB *gorm.DB
var argHelper helper.ArgHelper = helper.New()

var banner string = `
|\     /|  | \    /  |\      |  /------\ --------   /---
| \   / |  |   \_/   |  \    |  | |--| |     |     /____|
|  \/   |  |   / \   |    \  |  | |__| |     |    /-----|
|       |  |  /    \ |      \|  \______/     |   /_______

Mixnote is a audio platform for streaming, podcast and radio
@author Efedua Believe @epikoder Github
	`
var message string = `
usage: 
mixnote <command> [arguments]

Commands:
	serve		Start the server
	test		Run test
	migrate		Run migration
	help		Show help message
	version		Show version information
	
Use "mixnote help <command>" for more information about a command.
	`

func init() {
	DB, _ = database.DBConnection("")
	app.RegisterBindings(DB)
	fmt.Println(banner)
}

func helpServe() {
	fmt.Println(`
Usage:
mixnote serve [argument1] [argument2]

arguments: 
	--host=HOST		Server host default is localhost [127.0.0.1]
	--port=PORT		Server port default is :8080

	`)
}

func helpTest()    {}
func helpMigrate() {}
func helpVersion() {}

func main() {
	cmd := argHelper.Command()
	switch cmd {
	case "migrate":
		migrate.Migrate()
		os.Exit(0)

	case "serve":
		app.StartServer(argHelper.Option("host"), (func() int {
			_port := argHelper.Option("port")
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
		if cmd != "help" && len(os.Args) > 1 {
			fmt.Println("Unrecognized command : " + os.Args[1])
		}
		switch argHelper.Option("opt1") {
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
