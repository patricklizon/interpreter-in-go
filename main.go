package main

import (
	"fmt"
	"interpreter-in-go/repl"
	"os"
)

func main() {
	fmt.Printf("monkey programming language \n")
	repl.Start(os.Stdin, os.Stdout)
}
