package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/snahmik/gator-rss/internal/database"
)

func commandRegister(args []string, state *cliState) error {
	username := args[0]
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	user, err := state.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}

	fmt.Println("User successfully created!")
	fmt.Println(user)

	return setConfigCurrentUser(user.Name, state)
}

func commandLogin(args []string, state *cliState) error {
	username := args[0]

	user, err := state.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}

	return setConfigCurrentUser(user.Name, state)
}

func commandUsers(args []string, state *cliState) error {
	users, err := state.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}

	for _, user := range users {
		if state.config.CurrentUserName == user.Name {
			fmt.Printf("%s (current)\n", user.Name)
			continue
		}
		fmt.Println(user.Name)
	}

	return nil
}

func setConfigCurrentUser(name string, state *cliState) error {
	_, err := state.config.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("Current user set to: %s \n", name)
	return nil
}
