package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/32bytess/rss-agg/internal/database"
	"github.com/32bytess/rss-agg/internal/rssapi"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.args) != 0 {
		return errors.New("the agg handler expects no argument")
	}
	feed, err := rssapi.FetchFeed(context.Background(), testURL)
	if err != nil {
		return nil
	}
	fmt.Println(*feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return errors.New("the addfeed handler expects 2 argument, feedname and url")
	}
	name := cmd.args[0]
	url := cmd.args[1]
	user, err := s.db.GetUser(context.Background(), sql.NullString{
		String: s.config.CurrentUserName, Valid: true,
	})
	if err != nil {
		return err
	}
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("the feeds handler expects no argument")
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("- %s\n", feed.Name)
		fmt.Printf("\turl: %s\n", feed.Url)
		fmt.Printf("\tuser: %s\n", feed.Username.String)
	}
	return nil
}
