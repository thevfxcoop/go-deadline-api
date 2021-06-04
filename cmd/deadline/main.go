package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	// Modules
	"github.com/thevfxcoop/go-deadline-api"
	"github.com/thevfxcoop/go-deadline-api/pkg/client"
	"github.com/thevfxcoop/go-deadline-api/pkg/config"
)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

// Command-line flags
var (
	flagEndpoint *string
	flagDebug    *bool
)

///////////////////////////////////////////////////////////////////////////////
// MAIN

func main() {
	// Create flagset
	flags := flag.NewFlagSet("deadline", flag.ContinueOnError)
	defineFlags(flags)

	// Parse flags, if no command then ping for deadline version
	if err := flags.Parse(os.Args[1:]); err == flag.ErrHelp {
		config.PrintVersion(flags.Output())
		os.Exit(0)
	} else if err != nil {
		fmt.Fprintln(flags.Output(), err)
		os.Exit(-1)
	}

	// Set client options
	opts := []client.ClientOpt{}
	if *flagDebug {
		opts = append(opts, client.OptTrace(os.Stderr, true))
	}

	// Create client, ping then run command
	if endpoint, err := getEndpoint(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else if client, err := client.NewClient(endpoint, opts...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else if version, err := client.Ping(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else if flags.NArg() == 0 {
		fmt.Fprintln(os.Stderr, version)
		os.Exit(0)
	} else if err := Run(flags.Args(), client); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func getEndpoint() (*url.URL, error) {
	if endpoint := os.Getenv("DEADLINE_ENDPOINT"); endpoint != "" {
		return url.Parse(endpoint)
	} else if *flagEndpoint != "" {
		return url.Parse(*flagEndpoint)
	} else {
		return nil, deadline.ErrBadParameter.With("-endpoint")
	}
}

func defineFlags(flags *flag.FlagSet) {
	flags.Usage = func() { Usage(flags) }
	flagEndpoint = flags.String("endpoint", "", "Endpoint URL, can be overridden with DEADLINE_ENDPOINT environment variable")
	flagDebug = flags.Bool("debug", false, "Trace request and reponse with API")
}
