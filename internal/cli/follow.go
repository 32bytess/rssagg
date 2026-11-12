package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/32bytess/rssagg/internal/app"
	"github.com/32bytess/rssagg/internal/database"
)

func HandlerFollow(s *app.State, cmd Command, user database.User) error {

	if len(cmd.Args) != 1 {
		return errors.New("the follow handler expects one argument, the url")
	}
	url := cmd.Args[0]
	feed, err := s.DB.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("there is no fieed with this url")
	}
	feedFollow, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("failed to follow the feed (%s)", url)
	}
	fmt.Printf("user ( %s ) followed '%s' feed\n", feedFollow.UserName, feedFollow.FeedName)
	return nil
}
