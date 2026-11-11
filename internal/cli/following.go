package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/32bytess/rss-agg/internal/app"
)

func HandlerFollowing(s *app.State, cmd Command) error {

	if len(cmd.Args) != 0 {
		return errors.New("the following handler expects no argument")
	}
	followedFeeds, err := s.DB.GetFeedFollowsForUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}
	for i := range followedFeeds {
		fmt.Printf("- %s\n", followedFeeds[i].Feedname)
	}
	return nil
}
