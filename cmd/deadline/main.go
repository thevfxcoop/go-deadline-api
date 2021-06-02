package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/thevfxcoop/go-deadline-api"
	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

var (
	flagEndpoint = flag.String("endpoint", "", "Endpoint URL")
	flagDebug    = flag.Bool("debug", false, "Debug")
)

///////////////////////////////////////////////////////////////////////////////
// MAIN

func main() {
	flag.Parse()

	// Create logging instance
	log := log.New(os.Stderr, "deadline ", log.Ltime)

	// Set options
	opts := []client.ClientOpt{}
	if *flagDebug {
		opts = append(opts, client.OptTrace(os.Stderr, true))
	}

	// Create client
	if endpoint, err := getEndpoint(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else if client, err := client.NewClient(endpoint, opts...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else if version, err := client.Ping(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else {
		log.Println(version)

		// Call method
		if err := Run(flag.Args(), client, log); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
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
