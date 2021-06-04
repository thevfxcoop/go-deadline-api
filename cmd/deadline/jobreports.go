package main

import (
	"net/url"

	// Modules
	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type JobReports struct {
	command
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewJobReports(client *client.Client) Command {
	this := new(JobReports)
	this.Client = client
	return this
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *JobReports) Matches(args []string) url.Values {
	params := url.Values{}
	if args[0] == "jobreports" && len(args) == 2 {
		params.Set("id", args[1])
		return params
	}
	return nil
}

func (this *JobReports) Run(params url.Values) error {
	if jobs, err := this.GetJobReports(params.Get("id")); err != nil {
		return err
	} else {
		return this.output(jobs)
	}
}
