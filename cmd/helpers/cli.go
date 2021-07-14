package helpers

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type (
	cli struct {}
)

var once sync.Once

var map_ map[string]string
var cli_ *cli

func init() {
	map_ = make(map[string]string)
	map_["argv"] = "serve"
	for i := 1; i < len(os.Args); i++ {
		val := os.Args[i]
		if i == 1 {
			map_["argv"] = val
			continue
		}

		// for arguments
		if strings.Contains(val, "--") {
			key, val := stripeOptions(val)
			map_[key] = val
			continue
		}

		// for help commands
		if i == 2 {
			map_["opt"+strconv.Itoa(1)] = val
		}

		// For switches
		if len(val) == 2 && strings.Contains(val, "-") {
			map_["switches"] += val
		}
	}
}

func Cli() (*cli) {
	once.Do(func ()  {
		cli_ = new(cli)
	})
	return cli_
}

func stripeOptions(str string) (key string, val string) {
	key, val = (func() (string, string) {
		arr := strings.Split(str, "=")
		key = strings.TrimPrefix(arr[0], "--")
		val = arr[1]
		return key, val
	})()
	return
}

func (*cli) Argument() (s string) {
	return map_["argv"]
}

func (*cli) Option(option string) (s string) {
	if option == "argv" {
		panic("Access to commands not allowed from option")
	}
	return map_[option]
}

func (*cli) Switch(switch_ string) (bool) {
	return strings.Contains(map_["switches"], switch_) 
}

func (*cli) ArgsLenght() int {
	return len(os.Args)
}
