package main

import (
	"context"
	"fmt"
	"strings"
)

func commandReset(args []string, state *cliState) error {
	err := state.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}

	fmt.Println("Reset successful")
	return nil
}

func commandHelp(args []string, state *cliState) error {
	acceptedCommands := state.commands

	println("Welcome to Gator RSS!")
	println("Usages:")
	println("")

	fmt.Printf("%-10s %-20s %s \n", "Command", "Usage", "Description")
	for _, acceptedCommand := range acceptedCommands {
		fmt.Printf("%-10s %-20s %s \n", acceptedCommand.name, getCommandArgString(acceptedCommand), acceptedCommand.description)
	}

	return nil
}

func getCommandArgString(command cliCommand) string {
	if len(command.args) <= 0 {
		return "<>"
	}

	var commandArgString strings.Builder
	commandArgString.WriteString("<")
	for i, arg := range command.args {
		commandArgString.WriteString(fmt.Sprintf("%s", arg))
		if i+1 < len(command.args) {
			commandArgString.WriteString(", ")
			continue
		}
		commandArgString.WriteString(">")
	}

	return commandArgString.String()
}
