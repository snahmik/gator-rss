package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/snahmik/gator-rss/internal/config"
	"github.com/snahmik/gator-rss/internal/database"
)

type cliState struct {
	config   *config.Config
	db       *database.Queries
	commands map[string]cliCommand
}
type cliCommand struct {
	name        string
	args        []string
	description string
	callback    func([]string, *cliState) error
}

func startRepl(state *cliState) error {
	//scanner := bufio.NewScanner(os.Stdin)
	//fmt.Println("Starting Gator RSS")
	//for {
	//	fmt.Print("\nGator RSS > ")
	//
	//	scanner.Scan()
	//	userInput := strings.TrimSpace(scanner.Text())
	//
	//	if userInput == "" {
	//		continue
	//	}
	//
	//	err := handleUserInput(userInput, state)
	//	if err != nil {
	//		fmt.Printf("Error: %v \n", err)
	//	}
	//}

	userInput := os.Args
	if len(userInput) == 1 {
		return fmt.Errorf("no command specified")
	}

	//Remove file path from os.Args
	userInput = userInput[1:]
	var userCommand strings.Builder
	for _, arg := range userInput {
		userCommand.WriteString(arg + " ")
	}

	err := handleUserInput(userCommand.String(), state)
	if err != nil {
		return err
	}

	return nil
}

func handleUserInput(userInput string, state *cliState) error {
	command, args, err := extractUserInput(userInput)
	if err != nil {
		return err
	}

	acceptedCommands := state.commands
	targetCommand, ok := acceptedCommands[command]
	if !ok {
		return fmt.Errorf("unknown command: %v", command)
	}

	err = targetCommand.callback(args, state)
	if err != nil {
		return err
	}

	return nil
}

func extractUserInput(userInput string) (string, []string, error) {
	inputWords := strings.Fields(strings.ToLower(userInput))

	if len(inputWords) == 0 {
		return "", []string{}, fmt.Errorf("cannot extract empty user input")
	}

	return inputWords[0], inputWords[1:], nil
}
