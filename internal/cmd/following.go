package cmd

import (
	"context"
	"fmt"

	"github.com/landanqrew/rss-aggregator/internal/state"
)


func HandlerFollowing(s *state.State, cmd *Command) error {
	feedFollows, err := s.DBQueries.GetFeedFollowsForUser(context.Background(), s.Cfg.CurrentUser)
	if err != nil {
		return fmt.Errorf("error getting feed follows for user %s, %w", s.Cfg.CurrentUser, err)
	}

	if len(feedFollows) == 0 {
		return fmt.Errorf("no feed follows found for user %s", s.Cfg.CurrentUser)
	}

	for _, feedFollow := range feedFollows {
		fmt.Printf("%+v\n", feedFollow)
	}

	return nil
}