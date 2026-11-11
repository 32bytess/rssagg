package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/32bytess/rss-agg/internal/app"
	"github.com/32bytess/rss-agg/internal/database"
)

func HandlerRegister(s *app.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("the register handler expects a single argument, the username")
	}
	name := cmd.Args[0]
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user, err := s.DB.CreateUser(context.Background(), userParams)
	if err != nil {
		return errors.New("error creating user")
	}
	if err := s.Config.SetUser(name); err != nil {
		return err
	}
	fmt.Printf("created user: %s\n", user.Name)
	fmt.Printf("id: %v\n", user.ID)
	fmt.Printf("created at: %v\n", user.CreatedAt)
	return nil
}
