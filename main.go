package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/axbarsan/doggo/internal/repl"
	"github.com/axbarsan/doggo/internal/runner"
)

func main() {
	var fileName string

	args := os.Args
	if len(args) > 1 {
		fileName = args[1]
	}

	if fileName != "" {
		code, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(fmt.Sprintf("Cannot read file: %s", err.Error()))
		}

		r := runner.New()
		result := r.Run(string(code))
		fmt.Println(result)

		return
	}

	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the doggo programming language!\n", u.Name)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
