package main

import (
	"blog-aggregator/internal/database"
	"context"
	"fmt"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2

	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s <limit>", cmd.Name)
	}

	if len(cmd.Args) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{UserID: user.ID, Limit: int32(limit)})
	if err != nil {
		return fmt.Errorf("Unable to retrieve posts: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("---%s---\n", post.Title)
		fmt.Printf("   %v\n   ", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("====================================")
	}
	return nil
}
