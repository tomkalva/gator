package main

import "fmt"

type command struct {
	commandName string
	arguments   []string
}

type commands struct {
	commandNames map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if c.commandNames[cmd.commandName] == nil {
		return fmt.Errorf("Command doesn't exist")
	}
	return c.commandNames[cmd.commandName](s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandNames[name] = f
}
