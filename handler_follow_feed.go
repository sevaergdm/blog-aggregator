package main

import (
	"blog-aggregator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>\n", cmd.Name)
	}

	feedID, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Unable to find feed: %w", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feedID.ID,
	}

	row, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("Unable to follow feed: %w", err)
	}

	fmt.Printf("%s successfully followed %s", row.UserName, row.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	rows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Unable to retrieve followed feeds: %w", err)
	}

	if len(rows) == 0 {
		fmt.Println("No follwed feeds found")
		return nil
	}

	fmt.Printf("Feeds followed by %s:\n", user.Name)
	for _, row := range rows {
		fmt.Printf("* %s\n", row.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	params := database.UnfollowFeedParams{
		Name: user.Name,
		Url:  url,
	}

	err := s.db.UnfollowFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Unable to unfollow feed: %s", err)
	}

	fmt.Printf("Successfully unfollowed: %s\n", url)
	return nil
}
