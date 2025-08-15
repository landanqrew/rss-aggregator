package cmd

import (
	"context"
	"fmt"

	"github.com/landanqrew/rss-aggregator/internal/database"
	"github.com/landanqrew/rss-aggregator/internal/state"
)



func HandlerUnfollow(s *state.State, cmd *Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(cmd.Args))
	}

	curUserFollows, err := s.DBQueries.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("error getting feed follows for user %s, %w", user.Name, err)
	}

	for _, feedFollow := range curUserFollows {
		if feedFollow.FeedUrl == cmd.Args[0] {
			err = s.DBQueries.RemoveFeedFollow(context.Background(), feedFollow.ID)
			if err != nil {
				return fmt.Errorf("error removing feed follow %s, %w", feedFollow.ID, err)
			}
			fmt.Printf("Feed follow removed: %+v\n", feedFollow)
			return nil
		}
	}
	return fmt.Errorf("feed follow not found for user %s and feed url %s", user.Name, cmd.Args[0])
}