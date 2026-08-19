package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/snahmik/gator-rss/internal/config"
	"github.com/snahmik/gator-rss/internal/database"
)

func main() {
	systemCommands := map[string]cliCommand{
		"login":    {args: []string{"username"}, description: "Login to Gator RSS", callback: commandLogin},
		"register": {args: []string{"username"}, description: "Registers a new user to Gator RSS", callback: commandRegister},
		"users":    {args: []string{}, description: "Lists all registered users", callback: commandUsers},
		"agg":      {args: []string{"feedURL"}, description: "Displays the feed available for a given URL", callback: commandAgg},
		"addfeed":  {args: []string{"feedName", "feedURL"}, description: "Adds a new feed to Gator RSS", callback: commandAddFeed},
		"feeds":    {args: []string{}, description: "Lists all registered feeds", callback: commandFeeds},
		"reset":    {args: []string{}, description: "Removes all registered users", callback: commandReset},
		"help":     {args: []string{}, description: "Displays this help menu", callback: commandHelp},
	}

	systemConfig, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config: " + err.Error())
		os.Exit(1)
	}

	db, err := sql.Open("postgres", systemConfig.DbURL)
	if err != nil {
		fmt.Println("Error connecting to database: " + err.Error())
		os.Exit(1)
	}

	systemState := cliState{
		config:   systemConfig,
		db:       database.New(db),
		commands: systemCommands,
	}

	// Remove file path from os.Args by [1:]
	err = handleUserInput(os.Args[1:], &systemState)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
