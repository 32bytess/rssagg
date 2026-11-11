package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/32bytess/rss-agg/internal/app"
)

func HandlerFeeds(s *app.State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return errors.New("the feeds handler expects no argument")
	}
	feeds, err := s.DB.GetFeeds(context.Background())
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
