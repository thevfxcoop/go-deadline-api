package main

import (
	"log"
	"net/url"

	"github.com/thevfxcoop/go-deadline-api"
	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Command interface {
	Matches([]string) url.Values
	Run(url.Values) error
}

type command struct {
	*log.Logger
	*client.Client
}

///////////////////////////////////////////////////////////////////////////////
// RUN

func Run(args []string, client *client.Client, log *log.Logger) error {
	commands := []Command{}

	// Return silently if no command
	if len(args) == 0 {
		return nil
	}

	// Register commands
	commands = append(commands, NewJobs(client, log))
	commands = append(commands, NewJobReports(client, log))
	commands = append(commands, NewGroups(client, log))
	commands = append(commands, NewPools(client, log))
	commands = append(commands, NewWorkers(client, log))
	commands = append(commands, NewUsers(client, log))

	// Cycle through commands
	for _, cmd := range commands {
		if params := cmd.Matches(args); params != nil {
			return cmd.Run(params)
		}
	}

	// Command not found
	return deadline.ErrNotFound.With(args[0])
}
