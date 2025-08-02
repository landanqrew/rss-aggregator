package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/landanqrew/rss-aggregator/internal/config"
)


func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	config.CurrentUser = "landan"
	config.DBURL = "postgres://example.com:5432/rss-aggregator"
	err = config.SaveConfig()
	if err != nil {
		log.Fatalf("error saving config: %v", err)
	}
	
	jsonData,err := json.MarshalIndent(config, "", " ")
	if err != nil {
		log.Fatalf("error marshalling config: %v", err)
	}
	fmt.Println(string(jsonData))
}