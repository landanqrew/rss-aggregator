package cmd

import (
	"context"
	"fmt"

	"github.com/landanqrew/rss-aggregator/internal/database"
	"github.com/landanqrew/rss-aggregator/internal/state"
)



func MiddlewareHandler(handler func(s *state.State, cmd *Command, user database.User) error) func(s *state.State, cmd *Command) error {
	wrapped := func(s *state.State, cmd *Command) error {
		if s.Cfg.CurrentUser == "" {
			return fmt.Errorf("user not logged in")
		}
		usr, err := s.DBQueries.GetUserByName(context.Background(), s.Cfg.CurrentUser)
		if err != nil {
			return fmt.Errorf("error getting user by name, %w", err)
		}
		return handler(s, cmd, usr)
	}
	return wrapped
}
