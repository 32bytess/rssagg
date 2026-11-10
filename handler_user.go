package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/32bytess/rss-agg/internal/database"
	"github.com/32bytess/rss-agg/internal/rssapi"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("the login handler expects a single argument, the username")
	}
	name := cmd.args[0]
	user, err := s.db.GetUser(context.Background(), sql.NullString{
		String: name,
		Valid:  true,
	})
	if err != nil {
		return fmt.Errorf("there is no user with a %s name", name)
	}
	if err := s.config.SetUser(user.Name.String); err != nil {
		return err
	}
	fmt.Println("the new username has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("the register handler expects a single argument, the username")
	}
	name := cmd.args[0]
	userParams := database.CreateUserParams{
		ID: uuid.New(),
		Name: sql.NullString{
			String: name,
			Valid:  true,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return errors.New("error creating user")
	}
	if err := s.config.SetUser(name); err != nil {
		return err
	}
	fmt.Printf("created user: %s\n", user.Name.String)
	fmt.Printf("id: %v\n", user.ID)
	fmt.Printf("created at: %v\n", user.CreatedAt)
	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("the reset handler expects no argument")
	}
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return errors.New("reseting db failed")
	}
	fmt.Println("reset db was successful")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("the users handler expects no argument")
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return errors.New("getting users failed")
	}

	for i := range users {
		name := users[i].Name.String
		if name == s.config.CurrentUserName {
			name += " (current)"
		}
		fmt.Printf("* %s\n", name)
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {

	if len(cmd.args) != 0 {
		return errors.New("the agg handler expects no argument")
	}
	feed, err := rssapi.FetchFeed(context.Background(), testURL)
	if err != nil {
		return nil
	}
	fmt.Println(*feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return errors.New("the addfeed handler expects 2 argument, feedname and url")
	}
	name := cmd.args[0]
	url := cmd.args[1]
	user, err := s.db.GetUser(context.Background(), sql.NullString{
		String: s.config.CurrentUserName, Valid: true,
	})
	if err != nil {
		return err
	}
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
