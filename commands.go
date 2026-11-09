package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	handlerFunctions map[string](func(*state, command) error)
}

func (c *commands) run(s *state, cmd command) error {
	callback, ok := c.handlerFunctions[cmd.name]
	if !ok {
		return errors.New(cmd.name + " command not found")
	}
	return callback(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlerFunctions[name] = f
}
