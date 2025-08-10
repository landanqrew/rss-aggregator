package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/landanqrew/rss-aggregator/internal/database"
	"github.com/landanqrew/rss-aggregator/internal/state"
)



func HandlerFollow(s *state.State, cmd *Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("expected 1 argument, got %d", len(cmd.Args))
	}

	// need to fetch the feed by url
	feed, err := s.DBQueries.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error getting feed for url %s, %w", cmd.Args[0], err)
	}	
	
	// need to fetch the current user_id
	curUser, err := s.DBQueries.GetUserByName(context.Background(), s.Cfg.CurrentUser)
	if err != nil {
		return fmt.Errorf("error getting user with name (%s), %w", s.Cfg.CurrentUser, err)
	}

	feedFollow, err := s.DBQueries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    curUser.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow, %w", err)
	}

	fmt.Printf("Feed follow created: %+v\n", feedFollow)
	return nil
}