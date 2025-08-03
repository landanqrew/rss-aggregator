package main

import (
	"fmt"
	"log"
	"os"

	"github.com/landanqrew/rss-aggregator/internal/cmd"
	"github.com/landanqrew/rss-aggregator/internal/config"
	"github.com/landanqrew/rss-aggregator/internal/state"
)


func main() {
	s := state.State{}
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	s.Cfg = cfg
	commands := cmd.BuildCommandMap()
	commands.Register("login", cmd.HandlerLogin)
	
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