package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/landanqrew/rss-aggregator/internal/database"
	"github.com/landanqrew/rss-aggregator/internal/state"
)


func HandlerRegister(s *state.State, cmd *Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("an incorrect number of arguments were provided")
	}

	username := cmd.Args[0]

	s.Cfg.SetUser(username)

	err := s.Cfg.SaveConfig()
	if err != nil {
		return err
	}
	err = s.Cfg.SaveConfigBoots()
	if err != nil {
		return err
	}

	_, err = s.DBQueries.GetUserByName(context.Background(), username)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Printf("error getting user by name: %v", err)
			os.Exit(1)
		}
		fmt.Println("user not found, creating user")
		_, err = s.DBQueries.CreateUser(context.Background(), database.CreateUserParams{
			ID:        uuid.New().String(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      username,
		})

		if err != nil {
			return fmt.Errorf("error creating user: %w", err)
		}
		return nil
	}

	fmt.Println("user has already been registered")
	os.Exit(1)
	return nil
}