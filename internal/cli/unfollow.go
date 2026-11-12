package cli

import (
	"context"
	"errors"

	"github.com/32bytess/rssagg/internal/app"
	"github.com/32bytess/rssagg/internal/database"
)

func HandlerUnfollow(s *app.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("the unfollow handler expects a single argument, the feed url")
	}
	feedURL := cmd.Args[0]
	feed, err := s.DB.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return err
	}

	err = s.DB.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return err
	}
	return nil
}
