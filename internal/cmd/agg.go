package cmd

import (
	"context"
	"fmt"

	"github.com/landanqrew/rss-aggregator/internal/rss"
	"github.com/landanqrew/rss-aggregator/internal/state"
)




func HandlerAgg(s *state.State, cmd *Command) error {
	if len(cmd.Args) == 0 {
		cmd.Args = append(cmd.Args, "https://www.wagslane.dev/index.xml")
	}
	feed, err := rss.FetchFeed(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}