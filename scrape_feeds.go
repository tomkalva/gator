package main

import (
	"context"
	"fmt"
	"gator/internal/rss"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Fetching feed: %s\n", feed.Url)

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	rssFeeds, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range rssFeeds.Channel.Item {
		fmt.Println(item.Title)
	}
	return nil
}
