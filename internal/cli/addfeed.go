package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/32bytess/rss-agg/internal/app"
	"github.com/32bytess/rss-agg/internal/database"
)

func HandlerAddFeed(s *app.State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return errors.New("the addfeed handler expects 2 argument, feedname and url")
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
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
	feed, err := s.DB.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
