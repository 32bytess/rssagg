package app

import (
	"github.com/32bytess/rss-agg/internal/config"
	"github.com/32bytess/rss-agg/internal/database"
)

const TestURL = "https://www.wagslane.dev/index.xml"

type State struct {
	Config *config.Config
	DB     *database.Queries
}
