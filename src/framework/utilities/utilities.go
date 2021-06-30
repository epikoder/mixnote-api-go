package utilities

import (
	"fmt"
	"log"
)

var RED string = "\033[31m"
var GREEN string = "\033[32m"
var YELLOW string = "\033[33m"
var BLUE string = "\033[34m"
var CYAN string = "\033[36m"
var colorRESET string = "\033[0m"


type console struct {}
func Console() (c *console) {
	return
}

func (*console) Print(s interface{}) {
	fmt.Print(s, colorRESET)
}

func (*console) PrintLn(s interface{}) {
	fmt.Println(s, colorRESET)
}

func (*console) PrintF(s string, a ...interface{}) {
	fmt.Printf(s, a...)
	fmt.Print(colorRESET)
}

func (*console) Warn(s string, ) {
	fmt.Println(YELLOW, s, colorRESET)
}

func (*console) Error(s interface{}) {
	fmt.Println(RED, s, colorRESET)
}

func (*console) Success(s interface{}) {
	fmt.Println(GREEN, s, colorRESET)
}

func (*console) Fatal(s interface{})  {
	log.Fatal(RED, s, colorRESET)
}