package main

import (
	"fmt"

	"github.com/snahmik/gator-rss/internal/config"
	"github.com/snahmik/gator-rss/internal/database"
)

type cliState struct {
	config   *config.Config
	db       *database.Queries
	commands map[string]cliCommand
}
type cliCommand struct {
	args        []string
	description string
	callback    func([]string, *cliState) error
}

func handleUserInput(userInput []string, state *cliState) error {
	command, args, err := extractUserInput(userInput)
	if err != nil {
		return err
	}

	targetCommand, ok := state.commands[command]
	if !ok {
		return fmt.Errorf("unknown command: %s", command)
	}

	if len(args) != len(targetCommand.args) {
		fmt.Printf("Usage: %s - %v %s\n\n", command, targetCommand.args, targetCommand.description)
		return fmt.Errorf("invalid number of arguments: expected %d, got %d", len(targetCommand.args), len(args))
	}

	err = targetCommand.callback(args, state)
	if err != nil {
		return err
	}

	return nil
}

func extractUserInput(userInput []string) (string, []string, error) {
	switch len(userInput) {
	case 0:
		return "", nil, fmt.Errorf("no command specified")
	case 1:
		return userInput[0], nil, nil
	default:
		return userInput[0], userInput[1:], nil

	}
}
