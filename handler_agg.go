package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if cmd.Args == nil || len(cmd.Args) != 1 {
		return fmt.Errorf("agg command requires a time between requests")
	}
	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error parsing time between requests: %w", err)
	}
	fmt.Printf("Collecting feeds every %s\n", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *state) error {
	nextToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(), nextToFetch.ID)
	if err != nil {
		return err
	}
	fmt.Printf("Fetching feed %s\n", nextToFetch.Url)
	feed, err := fetchFeed(context.Background(), nextToFetch.Url)
	if err != nil {
		return err
	}
	for _, feedItem := range feed.Channel.Item {
		fmt.Printf("%+v\n", feedItem.Title)
	}
	return nil
}
