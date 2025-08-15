package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/landanqrew/rss-aggregator/internal/database"
	"github.com/landanqrew/rss-aggregator/internal/state"
)


func HandlerAddFeed(s *state.State, cmd *Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("incorrect number of arguments provided. Expected 2: a name and a url, got %d", len(cmd.Args))
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	// create feed in db
	dbFeed, err := s.DBQueries.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed, %w", err)
	}

	// create feed follow
	feedFollow, err := s.DBQueries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID: dbFeed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow, %w", err)
	}

	fmt.Printf("Feed created: %+v\n for user %s\n", feedFollow, user.Name)
	return nil
}