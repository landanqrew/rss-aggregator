package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/landanqrew/rss-aggregator/internal/database"
	"github.com/landanqrew/rss-aggregator/internal/state"
)



func HandlerBrowse(s *state.State, cmd *Command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("expected 1 argument (max_results), got %d", len(cmd.Args))
	}
	if len(cmd.Args) == 0 {
		cmd.Args = append(cmd.Args, "2")
	}

	maxResults, err := strconv.Atoi(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error parsing max results (%s), %w", cmd.Args[0], err)
	}
	posts, err := s.DBQueries.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		MaxResults: int32(maxResults),
	})
	if err != nil {
		return fmt.Errorf("error getting posts for user, %w", err)
	}
	if len(posts) == 0 {
		fmt.Printf("no posts found for user %s\n", user.Name)
		return nil
	}

	for _, post := range posts {
		fmt.Printf("%+v\n", post)
	}
	return nil
}