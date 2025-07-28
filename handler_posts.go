package main

import (
	"context"
	"fmt"
	"github.com/NeroQue/GATOR/internal/database"
	"strconv"
)

func HandlerBrowse(s *state, cmd command) error {
	limitStr := "2"

	if cmd.Args != nil && len(cmd.Args) >= 1 {
		limitStr = cmd.Args[0]
	}

	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		return fmt.Errorf("limit must be a number")
	}
	if limit < 1 {
		return fmt.Errorf("limit must be greater than 0")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		feed, err2 := s.db.GetFeedByUUID(context.Background(), post.FeedID)
		if err2 != nil {
			return err2
		}
		fmt.Printf("%s\n", post.Title)
		fmt.Printf("  URL: %s\n", post.Url)
		fmt.Printf("  Published at: %s\n", post.PublishedAt)
		fmt.Printf("  Created at: %s\n", post.CreatedAt)
		fmt.Printf("  Updated at: %s\n", post.UpdatedAt)
		fmt.Printf("  ID: %s\n", post.ID)
		fmt.Printf("  Feed: %s\n", feed.Name)
		fmt.Printf("\n")
	}
	return nil
}
