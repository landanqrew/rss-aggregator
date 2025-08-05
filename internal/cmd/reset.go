package cmd

import (
	"context"
	"fmt"

	"github.com/landanqrew/rss-aggregator/internal/state"
)



func HandlerReset(s *state.State, cmd *Command) error {
	err := s.DBQueries.RemoveAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error removing all users: %w", err)
	}

	return nil
}