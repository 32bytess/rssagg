package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/32bytess/rss-agg/internal/app"
	"github.com/32bytess/rss-agg/internal/cli"
	"github.com/32bytess/rss-agg/internal/config"
	"github.com/32bytess/rss-agg/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Read err: %v\n", err)
		os.Exit(1)
	}
	dbURL := cfg.DBURL

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("db err: %v\n", err)
		os.Exit(1)
	}
	st := app.State{
		Config: &cfg,
		DB:     database.New(db),
	}
	cmds := cli.NewCommands()
	// register commands
	cmds.Register("login", cli.HandlerLogin)
	cmds.Register("register", cli.HandlerRegister)
	cmds.Register("reset", cli.HandlerReset)
	cmds.Register("users", cli.HandlerUsers)
	cmds.Register("agg", cli.HandlerAgg)
	cmds.Register("addfeed", cli.HandlerAddFeed)
	cmds.Register("feeds", cli.HandlerFeeds)
	cmds.Register("follow", cli.HandlerFollow)
	cmds.Register("following", cli.HandlerFollowing)

	// get args (user input)
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: ./rss-agg <command>")
		os.Exit(1)
	}
	// create new command
	cmd := cli.Command{
		Name: args[1],
		Args: args[2:],
	}
	err = cmds.Run(&st, cmd)
	if err != nil {
		fmt.Printf("Err %v\n", err)
		os.Exit(1)
	}
}
