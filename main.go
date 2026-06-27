package main

import (
	"fmt"
	"gator/internal/config"
	"log"
	"os"
)

type state struct {
	cfgPointer *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	s := &state{
		cfgPointer: &cfg,
	}

	cmds := commands{
		commandNames: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("No command name given")
		os.Exit(1)
	}

	cmd := command{
		commandName: os.Args[1],
		arguments:   os.Args[2:],
	}
	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
