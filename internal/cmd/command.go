package cmd

import (
	"fmt"

	"github.com/landanqrew/rss-aggregator/internal/state"
)


type Command struct {
	CmdName string
	Args []string
}

type Commands struct {
	cmds map[string]func(*state.State, *Command) error
}



func BuildCommandMap() *Commands {
	return &Commands{
		cmds: map[string]func(*state.State, *Command) error {
			// "login": handlerLogin,
		},
	}
}

func (c *Commands) Run(s *state.State, cmd *Command) error {
	f, ok := c.cmds[cmd.CmdName]
	if !ok {
		return fmt.Errorf("command %s was not found in command registry", cmd.CmdName)
	}
	return f(s, cmd)
}

func (c *Commands) Register(name string, f func(s *state.State, cmd *Command) error) {
	c.cmds[name] = f
}