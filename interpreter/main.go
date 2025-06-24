package main

import (
	"fmt"
	repl "github.com/steviepreston/stelang/interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!, This is the SteLang programming language!\n", user.Username)
	fmt.Printf("Feel free to type in some commands!\n")
	repl.Start(os.Stdin, os.Stdout)
}
