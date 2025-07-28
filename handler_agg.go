package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/NeroQue/GATOR/internal/database"
	"github.com/google/uuid"
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
		pubDate, err := time.Parse(time.RFC1123Z, feedItem.PubDate)
		pubDateSQL := sql.NullTime{Valid: false}

		if err == nil {
			pubDateSQL = sql.NullTime{
				Time:  pubDate,
				Valid: true,
			}
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			FeedID:      nextToFetch.ID,
			Title:       feedItem.Title,
			Description: sql.NullString{String: feedItem.Description},
			Url:         feedItem.Link,
			PublishedAt: pubDateSQL,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
		if err != nil {
			if err.Error() == "pq: duplicate key value violates unique constraint \"posts_url_key\"" {
				fmt.Printf("Post already exists: %s\n", feedItem.Link)
				continue
			} else {
				fmt.Printf("Error creating post %s: %v\n", feedItem.Link, err)
			}
		} else {
			fmt.Printf("Post created: %s\n", feedItem.Link)
		}
	}
	return nil
}
