package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/landanqrew/rss-aggregator/internal/database"
	"github.com/landanqrew/rss-aggregator/internal/rss"
	"github.com/landanqrew/rss-aggregator/internal/state"
)




func HandlerAgg(s *state.State, cmd *Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("expected 1 argument (time_between_reqs), got %d", len(cmd.Args))
	}

	d, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error parsing duration (%s), %w", cmd.Args[0], err)
	}

	fmt.Printf("Collecting feeds every %s\n", d)
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("\n\nattempting to scrape feed...\n\n")
		err := scrapeFeeds(s)
		if err != nil {
			fmt.Printf("error scraping feeds, %v\n", err)
		}
	}
	return nil
}

func scrapeFeeds(s *state.State) error {
	feed, err := s.DBQueries.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch, %w", err)
	}

	feed.LastFetchedAt = sql.NullTime{Time: time.Now(), Valid: true}
	feed.UpdatedAt = time.Now()
	err = s.DBQueries.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchedAt: feed.LastFetchedAt,
		UpdatedAt: feed.UpdatedAt,
	})
	if err != nil {
		return fmt.Errorf("error marking feed fetched, %w", err)
	}

	rssFeed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed at url (%s), %w", feed.Url, err)
	}

	for _, item := range rssFeed.Channel.Item {
		exists, err := s.DBQueries.CheckPostExists(context.Background(), item.Link)
		if err != nil {
			return fmt.Errorf("error checking if post exists, %w", err)
		}
		if exists == 0 {
			publishedAt, err := time.Parse(time.RFC1123, item.PubDate)
			if err != nil {
				return fmt.Errorf("error parsing published at (%s), %w", item.PubDate, err)
			}
			s.DBQueries.CreatePost(context.Background(), database.CreatePostParams{
				ID:          uuid.New().String(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Title:       item.Title,
				Url:         item.Link,
				Description: item.Description,
				PublishedAt: publishedAt,
				FeedID:      feed.ID,
			})
		}
	}

	return nil
	

}

/*
Test cases:
TechCrunch: https://techcrunch.com/feed/
Hacker News: https://news.ycombinator.com/rss
Boot.dev Blog: https://blog.boot.dev/index.xml

go run . addFeed "TechCrunch" "https://techcrunch.com/feed/"
go run . addFeed "Hacker News" "https://news.ycombinator.com/rss"
go run . addFeed "Boot.dev Blog" "https://blog.boot.dev/index.xml"

go run . agg 5s

go run . follow "TechCrunch"
go run . follow "Hacker News"
go run . follow "Boot.dev Blog"
*/