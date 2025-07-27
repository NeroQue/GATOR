package main

import (
	"context"
)

func HandlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	return nil
}
