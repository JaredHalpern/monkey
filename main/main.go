package main

import (
	"fmt"
	"os"
	"os/user"
	"monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s", user.Username)
	fmt.Printf("Please type commands..")
	repl.Start(os.Stdin, os.Stdout)
}