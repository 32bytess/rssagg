package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/32bytess/rss-agg/internal/app"
	"github.com/32bytess/rss-agg/internal/database"
)

func HandlerFollowing(s *app.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 0 {
		return errors.New("the following handler expects no argument")
	}
	followedFeeds, err := s.DB.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return err
	}
	for i := range followedFeeds {
		fmt.Printf("- %s\n", followedFeeds[i].Feedname)
	}
	return nil
}
