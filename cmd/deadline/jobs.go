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
	this := new(Jobs)
	this.Client = client
	return this
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Jobs) Matches(args []string) url.Values {
	params := url.Values{}
	if args[0] == "jobs" && len(args) == 1 {
		return params
	}
	return nil
}

func (this *Jobs) Run(params url.Values) error {
	if jobs, err := this.Client.GetJobs(); err != nil {
		return err
	} else {
		return this.output(jobs)
	}
}
