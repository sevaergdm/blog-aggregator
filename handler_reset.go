package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.TruncateUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't truncate the users: %w", err)
	}
	fmt.Println("Successfully truncated the users table!")
	return nil
}
