package middlewares

import (
	"context"
	"errors"

	"github.com/32bytess/rss-agg/internal/app"
	"github.com/32bytess/rss-agg/internal/cli"
	"github.com/32bytess/rss-agg/internal/database"
)

func MiddlewareLoggedIn(handler func(s *app.State, cmd cli.Command, user database.User) error) func(*app.State, cli.Command) error {
	return func(s *app.State, cmd cli.Command) error {
		user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return errors.New("you need to be logged in")
		}
		return handler(s, cmd, user)
	}

}
