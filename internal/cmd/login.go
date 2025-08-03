package cmd

import (
	"fmt"

	"github.com/landanqrew/rss-aggregator/internal/state"
)


func HandlerLogin(s *state.State, cmd *Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("an incorrect number of arguments were provided")
	}

	username := cmd.Args[0]

	s.Cfg.SetUser(username)

	err := s.Cfg.SaveConfig()
	if err != nil {
		return err
	}

	/*
	// ONLY FOR boot.dev submission (I wanted a different directory for the config file)
	err = s.Cfg.SaveConfigBoots()
	if err != nil {
		return err
	}
	*/

	fmt.Println("user has been set")
	return nil
}