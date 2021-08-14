package utilities

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"

	"github.com/mixnote/mixnote-api-go/configs"
)

type (
	logger struct{}
)

var Logger *logger

func init() {
	once.Do(func() {
		Logger = new(logger)
	})
}

func (*logger) HandleError(err error) (ok bool) {
	if configs.App.Local || configs.App.LOG_STACK == "file" {
		f, err := os.OpenFile(configs.App.Path+"/logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		wrt := io.MultiWriter(os.Stdout, f)
		log.SetOutput(wrt)
	}

	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		log.Println(RED, "[ERROR]", YELLOW, fmt.Sprintf(" in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err), colorRESET)
		return false
	}
	return true
}
