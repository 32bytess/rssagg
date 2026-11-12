package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/32bytess/rssagg/internal/app"
)

func HandlerUsers(s *app.State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return errors.New("the users handler expects no argument")
	}
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return errors.New("getting users failed")
	}

	for i := range users {
		name := users[i].Name
		if name == s.Config.CurrentUserName {
			name += " (current)"
		}
		fmt.Printf("* %s\n", name)
	}
	return nil
}
