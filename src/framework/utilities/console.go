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

var once sync.Once
var Console *console

type console struct {}


func init() {
	once.Do(func() {
		Console = new(console)
	})
}

func (*console) Print(s string, a ...interface{}) {
	fmt.Print(fmt.Sprintf(s, a...))
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

func (*console) Error(s string, a ...interface{}) {
	fmt.Println(RED, fmt.Sprintf(s, a...), colorRESET)
}

func (*console) Success(s string, a ...interface{}) {
	fmt.Println(GREEN, fmt.Sprintf(s, a...), colorRESET)
}

func (*console) Fatal(s interface{})  {
	log.Fatal(RED, s, colorRESET)
}

func (*console) Log(s string, a ...interface{}) {
	fmt.Println(BLUE, fmt.Sprintf(s, a...), colorRESET)
} 