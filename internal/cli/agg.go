package cli

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/32bytess/rss-agg/internal/app"
	"github.com/32bytess/rss-agg/internal/database"
	"github.com/32bytess/rss-agg/internal/rssapi"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
		return err
	}
	for _, item := range feed.Channel.Item {
		publishedAt, err := dateparse.ParseAny(item.PubDate)
		postParams := database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: sql.NullTime{
				Time:  publishedAt,
				Valid: err == nil,
			},
			FeedID: fetchFeed.ID,
		}
		_, err = s.DB.CreatePost(context.Background(), postParams)
		if err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) && pqErr.Code == "23505" {
				continue
			}
			log.Fatal("err : %w", err)
		}
	}
	return nil
}
