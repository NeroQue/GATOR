package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/NeroQue/GATOR/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerLogin(s *state, cmd command) error {
	if cmd.Args == nil || len(cmd.Args) == 0 {
		return fmt.Errorf("login command requires a username")
	}
	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("user %s not found", cmd.Args[0])
	}
	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error setting user: %w", err)
	}
	fmt.Printf("Set current user to %s\n", s.cfg.CurrentUserName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if cmd.Args == nil || len(cmd.Args) == 0 {
		return fmt.Errorf("register command requires a username")
	}
	var arg database.CreateUserParams
	arg.ID = uuid.New()
	arg.Name = cmd.Args[0]
	arg.CreatedAt = time.Now()
	arg.UpdatedAt = time.Now()

	_, err := s.db.GetUser(context.Background(), arg.Name)
	if err == nil {
		return fmt.Errorf("user %s already exists", arg.Name)
	} else if !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("error checking if user exists: %w", err)
	}

	err = s.cfg.SetUser(arg.Name)

	_, err = s.db.CreateUser(context.Background(), arg)
	if err != nil {
		return err
	}

	fmt.Printf("User %s created\n", arg.Name)
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	if len(users) == 0 {
		fmt.Errorf("no users found")
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
