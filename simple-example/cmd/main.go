package main

import (
	"fmt"
	"hello"
	"os"
)

func main() {
	// os.Args is a slice of strings
	// os.Args[0] is the name of the command - mostly in many programming languages
	// os.Args[1:] is a slice of strings that contains all the arguments excluding the command name (os.Args[0])
	fmt.Println(hello.SayHello(os.Args[1:]))
} 