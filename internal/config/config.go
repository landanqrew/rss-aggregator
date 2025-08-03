package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/landanqrew/rss-aggregator/internal/utils"
)


type Config struct {
	DBURL string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error getting home directory: %w", err)
	}
	// If the config file does not exist, create it with the default config
	if !utils.FileExists(homeDir + "/.rss-aggregator/config.json") {
		defaultConfig := generateDefaultConfig()
		data, err := json.MarshalIndent(defaultConfig, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("error marshalling default config: %w", err)
		}
		err = utils.CreateFile(homeDir + "/.rss-aggregator/config.json", data)
		if err != nil {
			return nil, fmt.Errorf("error creating config file: %w", err)
		}
		return defaultConfig, nil
	}
	// If the config file exists, read it and unmarshal it into the config struct
	data, err := utils.ReadFile(homeDir + "/.rss-aggregator/config.json")
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config file: %w", err)
	}
	return config, nil
}

func (c *Config) SetUser(user string) {
	c.CurrentUser = user
}

func (c *Config) SetDBURL(url string) {
	c.DBURL = url
}

func (c *Config) SaveConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %w", err)
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling config: %w", err)
	}
	err = utils.CreateFile(homeDir + "/.rss-aggregator/config.json", data)
	if err != nil {
		return fmt.Errorf("error creating config file: %w", err)
	}
	return nil
}

func (c *Config) SaveConfigBoots() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %w", err)
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling config: %w", err)
	}
	err = utils.CreateFile(homeDir + "/.gatorconfig.json", data)
	if err != nil {
		return fmt.Errorf("error creating config file: %w", err)
	}
	return nil
}


func (c *Config) PrintCfg() {
	jsonData,err := json.MarshalIndent(c, "", " ")
	if err != nil {
		log.Fatalf("error marshalling config: %v", err)
	}
	fmt.Println(string(jsonData))
}


func generateDefaultConfig() *Config {
	return &Config{
		DBURL: "localhost:5432",
		CurrentUser: "default",
	}
}