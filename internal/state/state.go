package state

import (

	"github.com/landanqrew/rss-aggregator/internal/config"
	"github.com/landanqrew/rss-aggregator/internal/database"
)


type State struct {
	Cfg *config.Config
	DBQueries *database.Queries
}

