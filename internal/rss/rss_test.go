package rss

import (
	"context"
	"testing"
)



func TestFetchFeed(t *testing.T) {
	expectedTitle := "Lane's Blog"
	
	feed, err := FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		t.Fatalf("error fetching feed: %v", err)
	}

	if feed.Channel.Title != expectedTitle {
		t.Errorf("expected title '%s', got '%s'", expectedTitle, feed.Channel.Title)
	}

}