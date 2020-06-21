package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/axbarsan/doggo/internal/repl"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the doggo programming language!\n", u.Name)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
