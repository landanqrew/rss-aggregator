package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/landanqrew/rss-aggregator/internal/database"
	"github.com/landanqrew/rss-aggregator/internal/state"
)


func HandlerAddFeed(s *state.State, cmd *Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("incorrect number of arguments provided. Expected 2: a name and a url, got %d", len(cmd.Args))
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	// need to fetch the current user_id
	curUser, err := s.DBQueries.GetUserByName(context.Background(), s.Cfg.CurrentUser)
	if err != nil {
		return fmt.Errorf("error getting user with name (%s), %w", s.Cfg.CurrentUser, err)
	}

	// create feed in db
	dbFeed, err := s.DBQueries.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: curUser.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed, %w", err)
	}

	fmt.Printf("Feed created: %+v\n", dbFeed)
	return nil
}