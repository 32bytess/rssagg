package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/32bytess/rssagg/internal/app"
	"github.com/32bytess/rssagg/internal/database"
)

func HandlerAddFeed(s *app.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return errors.New("the addfeed handler expects 2 argument, feedname and url")
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	feed, err := s.DB.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}
	_, err = s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
