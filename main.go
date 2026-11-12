package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/32bytess/rssagg/internal/app"
	"github.com/32bytess/rssagg/internal/cli"
	"github.com/32bytess/rssagg/internal/config"
	"github.com/32bytess/rssagg/internal/database"
	"github.com/32bytess/rssagg/internal/middlewares"

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
	cmds.Register("addfeed", middlewares.MiddlewareLoggedIn(cli.HandlerAddFeed))
	cmds.Register("feeds", cli.HandlerFeeds)
	cmds.Register("follow", middlewares.MiddlewareLoggedIn(cli.HandlerFollow))
	cmds.Register("following", middlewares.MiddlewareLoggedIn(cli.HandlerFollowing))
	cmds.Register("unfollow", middlewares.MiddlewareLoggedIn(cli.HandlerUnfollow))
	cmds.Register("browse", cli.HandlerBrowse)

	// get args (user input)
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: ./rssagg <command>")
		os.Exit(1)
	}
	// create new command
	cmd := cli.Command{
		Name: args[1],
		Args: args[2:],
	}
	err = cmds.Run(&st, cmd)
	if err != nil {
		fmt.Printf("Err:\t %v\n", err)
		os.Exit(1)
	}
}
