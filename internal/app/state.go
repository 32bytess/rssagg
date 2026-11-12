package app

import (
	"github.com/32bytess/rssagg/internal/config"
	"github.com/32bytess/rssagg/internal/database"
)

const TestURL = "https://www.wagslane.dev/index.xml"

type State struct {
	Config *config.Config
	DB     *database.Queries
}
