package main

import (
	"monkey/repl"
	"os"
    "fmt"
    "os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in command\n")
	repl.Start(os.Stdin, os.Stdout)
}
