package cli

import (
	"context"
	"errors"

	"github.com/32bytess/rss-agg/internal/app"
	"github.com/32bytess/rss-agg/internal/rssapi"
)

func HandlerAgg(s *app.State, cmd Command) error {

	if len(cmd.Args) != 0 {
		return errors.New("the agg handler expects no argument")
	}
	feed, err := rssapi.FetchFeed(context.Background(), app.TestURL)
	if err != nil {
		return nil
	}
	feed.PrintDetails()
	return nil
}
