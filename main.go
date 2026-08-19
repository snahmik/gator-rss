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
		"login":    {name: "login", args: []string{"username"}, description: "Login to Gator RSS", callback: commandLogin},
		"register": {name: "register", args: []string{"username"}, description: "Registers a new user to Gator RSS", callback: commandRegister},
		"users":    {name: "users", args: []string{}, description: "Lists all registered users", callback: commandUsers},
		"agg":      {name: "agg", args: []string{"feedURL"}, description: "Displays the feed available for a given URL", callback: commandAgg},
		"reset":    {name: "reset", args: []string{}, description: "Removes all registered users", callback: commandReset},
		"help":     {name: "help", args: []string{}, description: "Displays this help menu", callback: commandHelp},
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

	err = startRepl(&systemState)
	if err != nil {
		fmt.Println("REPL error: " + err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
