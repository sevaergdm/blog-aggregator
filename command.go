package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdFunc, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("%s does not exist", cmd.name)
	}
	return cmdFunc(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) error {
	_, ok := c.cmds[name]
	if !ok {
		c.cmds[name] = f
		return nil
	}
	return fmt.Errorf("command %s is already registered", name)
}
