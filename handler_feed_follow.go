package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/32bytess/rss-agg/internal/database"
)

func handlerFollow(s *state, cmd command) error {

	if len(cmd.args) != 1 {
		return errors.New("the follow handler expects one argument, the url")
	}
	url := cmd.args[0]
	user, err := s.db.GetUser(context.Background(), sql.NullString{
		String: s.config.CurrentUserName,
		Valid:  true,
	})
	if err != nil {
		return fmt.Errorf("user not found")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("there is no fieed with this url")
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("failed to follow the feed (%s)", url)
	}
	fmt.Printf("user ( %s ) followed '%s' feed\n", feedFollow.UserName.String, feedFollow.FeedName)
	return nil
}
