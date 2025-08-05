package cmd

import (
	"context"
	"fmt"

	"github.com/landanqrew/rss-aggregator/internal/state"
)


func HandlerUsers(s *state.State, cmd *Command) error {
	userList, err := s.DBQueries.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users: %w", err)
	}

	for _, user := range userList {
		appendStr := ""
		if user.Name == s.Cfg.CurrentUser {
			appendStr = " (current)"
		}
		fmt.Println(user.Name + appendStr) 
	}

	return nil
}