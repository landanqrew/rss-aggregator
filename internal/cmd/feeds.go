package cmd

import (
	"context"
	"fmt"

	"github.com/landanqrew/rss-aggregator/internal/state"
)


func HandlerFeeds(s *state.State, cmd *Command) error {
	feeds, err := s.DBQueries.GetAllUserFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting all user feeds, %w", err)
	}

	for _, feed := range feeds {
		fmt.Println(feed)
	}
	return nil
}