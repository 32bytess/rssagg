package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/32bytess/rss-agg/internal/config"
	"github.com/32bytess/rss-agg/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

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
	st := state{
		config: &cfg,
		db:     database.New(db),
	}
	cmds := commands{
		handlerFunctions: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("err: Not enough args\n")
		os.Exit(1)
	}
	cmd := command{
		name: args[1],
		args: args[2:],
	}

	err = cmds.run(&st, cmd)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)

	}
	os.Exit(0)
}
