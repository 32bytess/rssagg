package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/32bytess/rssagg/internal/app"
)

func HandlerLogin(s *app.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("the login handler expects a single argument, the username")
	}
	name := cmd.Args[0]
	user, err := s.DB.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("there is no user with a %s name", name)
	}
	if err := s.Config.SetUser(user.Name); err != nil {
		return err
	}
	fmt.Println("the new username has been set")
	return nil
}
