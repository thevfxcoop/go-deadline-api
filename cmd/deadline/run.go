package main

import (
	"encoding/json"
	"flag"
	"fmt"
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
	*client.Client
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func Run(args []string, client *client.Client) error {
	// Register commands
	commands := []Command{
		NewJobs(client),
		NewJobReports(client),
		NewTasks(client),
		NewWorkers(client),
		NewGroups(client),
		NewPools(client),
		NewUsers(client),
	}

	// Cycle through commands
	for _, cmd := range commands {
		if params := cmd.Matches(args); params != nil {
			return cmd.Run(params)
		}
	}

	return deadline.ErrNotFound.With(args[0])
}

func Usage(flags *flag.FlagSet) {
	fmt.Fprintf(flags.Output(), "Usage of %v:\n", flags.Name())
	fmt.Fprintf(flags.Output(), "  deadline <flags> <command> (<args>)\n")
	fmt.Fprintf(flags.Output(), "\nFlags:\n")
	flags.PrintDefaults()
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *command) output(v interface{}) error {
	if data, err := json.MarshalIndent(v, "", "  "); err != nil {
		return err
	} else {
		fmt.Println(string(data))
	}
	return nil
}
