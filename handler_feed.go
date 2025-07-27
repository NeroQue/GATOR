package main

import (
	"context"
	"fmt"
	"github.com/NeroQue/GATOR/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if cmd.Args == nil || len(cmd.Args) != 2 {
		return fmt.Errorf("feed command requires a feed name and a feed URL")
	}
	arg := database.CreateFeedParams{}
	arg.ID = uuid.New()
	arg.UserID = user.ID
	arg.Name = cmd.Args[0]
	arg.Url = cmd.Args[1]
	arg.CreatedAt = time.Now().UTC()
	arg.UpdatedAt = time.Now().UTC()
	feed, err := s.db.CreateFeed(context.Background(), arg)
	if err != nil {
		return err
	}

	arg2 := database.CreateFeedFollowParams{}
	arg2.ID = uuid.New()
	arg2.UserID = user.ID
	arg2.FeedID = feed.ID
	arg2.CreatedAt = time.Now().UTC()
	arg2.UpdatedAt = time.Now().UTC()
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), arg2)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", feed)
	fmt.Printf("Added feed %s to user %s on the %s list\n", feed.Name, feedFollow.UserName, feedFollow.ID)
	return nil

}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}
	if len(feeds) == 0 {
		fmt.Errorf("no feeds found")
	}

	for _, feed := range feeds {
		createdBy, err := s.db.GetUserByUUID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		fmt.Printf("* %s\n", feed.Name)
		fmt.Printf("  URL: %s\n", feed.Url)
		fmt.Printf("  User: %s\n", feed.UserID)
		fmt.Printf("  Created at: %s\n", feed.CreatedAt)
		fmt.Printf("  Updated at: %s\n", feed.UpdatedAt)
		fmt.Printf("  ID: %s\n", feed.ID)
		fmt.Printf("  Created by: %s\n", createdBy.Name)
		fmt.Printf("\n")
	}
	return nil
}

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if cmd.Args == nil || len(cmd.Args) != 1 {
		fmt.Errorf("feed command requires a feed URL")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	arg := database.CreateFeedFollowParams{}
	arg.ID = uuid.New()
	arg.UserID = user.ID
	arg.FeedID = feed.ID
	arg.CreatedAt = time.Now().UTC()
	arg.UpdatedAt = time.Now().UTC()
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), arg)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feedFollow)
	fmt.Printf("Feed %s followed by %s\n", feedFollow.FeedName, feedFollow.UserName)
	return nil
}

func handlerFollowingFeedsByUser(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.FeedName)
	}
	return nil
}

func handlerUnfollowFeed(s *state, cmd command, user database.User) error {
	if cmd.Args == nil || len(cmd.Args) != 1 {
		fmt.Errorf("feed command requires a feed URL")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	arg := database.DeleteFeedFollowParams{}
	arg.Name = user.Name
	arg.Url = feed.Url
	err = s.db.DeleteFeedFollow(context.Background(), arg)
	if err != nil {
		return err
	}
	fmt.Printf("Feed %s unfollowed by %s\n", feed.Name, user.Name)
	return nil
}
