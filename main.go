package main

import (
	"fmt"
	"log"

	"github.com/32bytess/rss-agg/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Read err: %v", err)
		return
	}
	if err := cfg.SetUser("lane"); err != nil {
		log.Fatalf("SetUser err: %v", err)
		return
	}

	c, err := config.Read()
	if err != nil {
		log.Fatalf("Read 2 err: %v", err)
		return
	}
	fmt.Printf("db_url: %s\n", c.DBURL)
	fmt.Printf("username: %s\n", c.CurrentUserName)
}
