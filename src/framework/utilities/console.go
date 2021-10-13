package utilities

import (
	"fmt"
	"log"
	"sync"
)

var RED string = "\033[31m"
var GREEN string = "\033[32m"
var YELLOW string = "\033[33m"
var BLUE string = "\033[34m"
var CYAN string = "\033[36m"
var colorRESET string = "\033[0m"

var (
	once    sync.Once
	Console *console
)

type console struct{}

func init() {
	once.Do(func() {
		Console = new(console)
	})
}

func (*console) Print(s string, a ...interface{}) {
	fmt.Printf(s, a...)
}

func (*console) Printc(s string, color string, a ...interface{}) {
	if color == "" {
		color = colorRESET
	}
	fmt.Println(color, fmt.Sprintf(s, a...))
}

func (*console) Println(s string, a ...interface{}) {
	fmt.Println(fmt.Sprintf(s, a...))
}

func (*console) Warn(s string, a ...interface{}) {
	fmt.Println(YELLOW, fmt.Sprintf(s, a...), colorRESET)
}

func (*console) Debug(s string, a ...interface{}) {
	fmt.Println(BLUE, "DEBUG : ", YELLOW, fmt.Sprintf(s, a...), colorRESET)
}

func (*console) Error(s string, a ...interface{}) {
	fmt.Println(RED, "ERROR : ", YELLOW, fmt.Sprintf(s, a...), colorRESET)
}

func (*console) Success(s string, a ...interface{}) {
	fmt.Println(GREEN, fmt.Sprintf(s, a...), colorRESET)
}

func (*console) Fatal(s ...interface{}) {
	var m string
	m, ok := s[0].(string)
	if !ok {
		e := s[0].(error)
		log.Panic(RED, "ERROR : ", YELLOW, e.Error(), colorRESET)
	}
	if len(s) > 1 {
		f := s[0].(string)
		var i []interface{}
		for x := 1; x < len(s); x++ {
			i = append(i, s[x])
		}
		m = fmt.Sprintf(f, i...)
	}
	log.Panic(RED, "ERROR : ", YELLOW, m, colorRESET)
}
