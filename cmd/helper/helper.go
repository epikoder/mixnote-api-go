package helper

import (
	"os"
	"strconv"
	"strings"
)

type (
	icommand interface {
		Command()
	}

	ArgHelper struct {
		icommand
	}
)

var map_ map[string]string

func init() {
	map_ = make(map[string]string)
	map_["cmd"] = "serve"
	for i := 1; i < len(os.Args); i++ {
		val := os.Args[i]
		if i == 1 {
			map_["cmd"] = val
			continue
		}

		// for arguments
		if strings.Contains(val, "--") {
			key, val := stripeOptions(val)
			map_[key] = val
			continue
		}

		// for help commands
		map_["opt"+strconv.Itoa(i-1)] = val
	}
}

func New() (arghelper ArgHelper) {
	return
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

func (argv *ArgHelper) Command() (s string) {
	return map_["cmd"]
}

func (argv ArgHelper) Option(option string) (s string) {
	if option == "cmd" {
		panic("Access to commands not allowed from option")
	}
	return map_[option]
}
