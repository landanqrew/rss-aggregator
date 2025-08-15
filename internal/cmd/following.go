package cmd

import (
	"context"
	"fmt"

	"github.com/landanqrew/rss-aggregator/internal/database"
	"github.com/landanqrew/rss-aggregator/internal/state"
)


func HandlerFollowing(s *state.State, cmd *Command, user database.User) error {
	feedFollows, err := s.DBQueries.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("error getting feed follows for user %s, %w", user.Name, err)
	}

	if len(feedFollows) == 0 {
		fmt.Printf("no feed follows found for user %s", user.Name)
		return nil
	}

	for _, feedFollow := range feedFollows {
		fmt.Printf("%+v\n", feedFollow)
	}

	return nil
}