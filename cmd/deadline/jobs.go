package main

import (
	"net/url"

	// Modules
	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Jobs struct {
	command
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewJobs(client *client.Client) Command {
	return &Jobs{command{Client: client}}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Jobs) Matches(args []string) (fn, url.Values) {
	params := url.Values{}
	if args[0] == "jobs" && len(args) == 1 {
		return this.RunJobs, params
	}
	return nil, nil
}

func (this *Jobs) RunJobs(params url.Values) error {
	if jobs, err := this.Client.GetJobs(); err != nil {
		return err
	} else {
		return this.output(jobs)
	}
}
