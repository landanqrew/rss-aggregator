package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/landanqrew/rss-aggregator/internal/cmd"
	"github.com/landanqrew/rss-aggregator/internal/config"
	"github.com/landanqrew/rss-aggregator/internal/database"
	"github.com/landanqrew/rss-aggregator/internal/state"
	_ "github.com/lib/pq"
)


func main() {
	s := state.State{}
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	s.Cfg = cfg
	// Build Command Registry
	commands := cmd.BuildCommandMap()
	commands.Register("login", cmd.HandlerLogin)
	commands.Register("register", cmd.HandlerRegister)
	commands.Register("reset", cmd.HandlerReset)
	commands.Register("users", cmd.HandlerUsers)
	commands.Register("agg", cmd.HandlerAgg)
	commands.Register("addfeed", cmd.MiddlewareHandler(cmd.HandlerAddFeed))
	commands.Register("feeds", cmd.HandlerFeeds)
	commands.Register("follow", cmd.MiddlewareHandler(cmd.HandlerFollow))
	commands.Register("following", cmd.MiddlewareHandler(cmd.HandlerFollowing))
	commands.Register("unfollow", cmd.MiddlewareHandler(cmd.HandlerUnfollow))
	commands.Register("browse", cmd.MiddlewareHandler(cmd.HandlerBrowse))

	// Open Database Connection
	db, err := sql.Open("postgres", s.Cfg.DBURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close() // dont forget to close the database connection
	s.DBQueries = database.New(db)
	
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("no command provided")
		os.Exit(1)
	}

	cmd := &cmd.Command{
		CmdName: args[0],
		Args: args[1:],
	}

	err = commands.Run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}