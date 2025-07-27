package main

import (
	"context"
	"fmt"
)

func HandlerAgg(s *state, cmd command) error {
	//if cmd.Args == nil || len(cmd.Args) == 0 {
	//	return fmt.Errorf("feed command requires a feed URL")
	//}
	//feed, err := fetchFeed(context.Background(), cmd.Args[0])
	//if err != nil {
	//	return err
	//}
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)
	return nil
}
