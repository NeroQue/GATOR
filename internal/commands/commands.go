package commands

import "fmt"

type Command struct {
	Name string
	Args []string
}
type Commands struct {
	Commands map[string]func(*State, Command) error
}

func HandlerLogin(s *State, cmd Command) error {
	if cmd.Args == nil || len(cmd.Args) == 0 {
		return fmt.Errorf("login command requires a username")
	}
	err := s.CFG.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error setting user: %w", err)
	}
	fmt.Printf("Set current user to %s\n", s.CFG.CurrentUserName)
	return nil
}

func (c *Commands) Run(s *State, cmd Command) error {
	if c.Commands == nil {
		return fmt.Errorf("commands not initialized")
	}
	if handler, exists := c.Commands[cmd.Name]; !exists {
		return fmt.Errorf("command %s not found", cmd.Name)
	} else {
		err := handler(s, cmd)
		if err != nil {
			return fmt.Errorf("error running command %s: %w", cmd.Name, err)
		}
		return nil
	}
}

func (c *Commands) Register(name string, f func(*State, Command) error) error {
	if c.Commands == nil {
		return fmt.Errorf("commands not initialized")
	}

	if name == "" {
		return fmt.Errorf("command name cannot be empty")
	}

	if f == nil {
		return fmt.Errorf("command handler cannot be nil")
	}

	if c.Commands[name] != nil {
		return fmt.Errorf("command %s already registered", name)
	}

	c.Commands[name] = f
	return nil
}
