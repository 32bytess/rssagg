package cli

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/32bytess/rssagg/internal/app"
)

func HandlerBrowse(s *app.State, cmd Command) error {

	if len(cmd.Args) > 1 {
		return errors.New("the browse command expects at most one argument")
	}
	limit := int32(2)
	if len(cmd.Args) == 1 {
		parsedLimit, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid limit %q:%w", cmd.Args[0], err)
		}
		limit = int32(parsedLimit)
	}

	posts, err := s.DB.GetPosts(context.Background(), limit)
	if err != nil {
		return err
	}
	for i := range posts {
		fmt.Println()
		fmt.Println()
		fmt.Printf("Post %v\n\n", i+1)
		fmt.Printf("- Title: %s\n", posts[i].Title)
		fmt.Printf("- URL: %s\n", posts[i].Url)
		fmt.Printf("- Description:  %v\n", posts[i].Description)
		fmt.Printf("- Published At: %s\n", posts[i].PublishedAt.Time)
		fmt.Printf("- Published In: %v\n", posts[i].FeedName)
	}
	return nil
}
