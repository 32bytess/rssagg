package cli

import (
	"errors"

	"github.com/32bytess/rss-agg/internal/app"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	HandlerFunctions map[string](func(*app.State, Command) error)
}

func (c *Commands) Run(s *app.State, cmd Command) error {
	callback, ok := c.HandlerFunctions[cmd.Name]
	if !ok {
		return errors.New(cmd.Name + " command not found")
	}
	return callback(s, cmd)
}

func (c *Commands) Register(name string, f func(*app.State, Command) error) {
	c.HandlerFunctions[name] = f
}
