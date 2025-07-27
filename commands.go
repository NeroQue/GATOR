package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}
type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if c.registeredCommands == nil {
		return fmt.Errorf("commands not initialized")
	}
	if handler, exists := c.registeredCommands[cmd.Name]; !exists {
		return fmt.Errorf("command %s not found", cmd.Name)
	} else {
		err := handler(s, cmd)
		if err != nil {
			return fmt.Errorf("error running command %s: %w", cmd.Name, err)
		}
		return nil
	}
}

func (c *commands) register(name string, f func(*state, command) error) error {
	if c.registeredCommands == nil {
		return fmt.Errorf("commands not initialized")
	}
	if name == "" {
		return fmt.Errorf("command name cannot be empty")
	}
	if f == nil {
		return fmt.Errorf("command handler cannot be nil")
	}
	if c.registeredCommands[name] != nil {
		return fmt.Errorf("command %s already registered", name)
	}
	c.registeredCommands[name] = f
	return nil
}
