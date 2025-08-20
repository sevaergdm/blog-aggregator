package main

import (
	"blog-aggregator/internal/config"
	"errors"
	"fmt"
)

type state struct {
	config *config.Config
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Expected a username")
	}

	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User: %s has been set\n", cmd.args[0])
	
	return nil
}
