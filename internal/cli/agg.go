package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/32bytess/rss-agg/internal/app"
	"github.com/32bytess/rss-agg/internal/rssapi"
)

func HandlerAgg(s *app.State, cmd Command) error {

	if len(cmd.Args) != 1 {
		return errors.New("the agg handler expects no argument, the time_between_reqs")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %s \n", timeBetweenRequests.String())
	ticker := time.NewTicker(timeBetweenRequests)
	for range ticker.C {
		fmt.Println("fetching feeds ...")
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func scrapeFeeds(s *app.State) error {
	fetchFeed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.DB.MarkFeedFetched(context.Background(), fetchFeed.ID)
	if err != nil {
		return err
	}

	feed, err := rssapi.FetchFeed(context.Background(), app.TestURL)
	if err != nil {
		return nil
	}
	for _, item := range feed.Channel.Item {
		fmt.Printf("- title: %s\n", item.Title)
	}
	return nil
}
