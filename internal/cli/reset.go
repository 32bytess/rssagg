package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/32bytess/rss-agg/internal/app"
)

func HandlerReset(s *app.State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return errors.New("the reset handler expects no argument")
	}
	err := s.DB.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("reseting db failed: %w", err)
	}
	fmt.Println("reset db was successful")
	return nil
}
