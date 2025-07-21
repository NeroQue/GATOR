package main

import (
	"fmt"
	"github.com/NeroQue/GATOR/internal/commands"
	"github.com/NeroQue/GATOR/internal/config"
	"os"

	"log"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "error: not enough arguments\n")
		os.Exit(1)
	}

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	state := &commands.State{
		CFG: &cfg,
	}

	cmds := commands.Commands{
		Commands: make(map[string]func(*commands.State, commands.Command) error),
	}
	err = cmds.Register("login", commands.HandlerLogin)
	if err != nil {
		log.Fatalf("error registering command: %v", err)
	}

	cmd := commands.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := cmds.Run(state, cmd); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
