package main

import (
	"fmt"
	"os"

	"github.com/element-box/gator-rss/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("gator-rss needs at least 2 args: a command and argument\n")
		os.Exit(1)
	}
	initState := state{}
	cfg := config.Read()
	initState.cfg = &cfg
	cmds := commands{make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	userCommand := command{os.Args[1], os.Args[2:]}

	err := cmds.run(&initState, userCommand)
	if err != nil {
		fmt.Printf("Error on command: %s, %v\n", userCommand.name, err)
		os.Exit(1)
	}
}
