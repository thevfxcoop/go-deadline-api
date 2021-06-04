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
	return &JobReports{command{Client: client}}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *JobReports) Matches(args []string) (fn, url.Values) {
	params := url.Values{}
	if args[0] == "jobreports" && len(args) == 2 {
		params.Set("job", args[1])
		return this.RunJobReports, params
	}
	return nil, nil
}

func (this *JobReports) RunJobReports(params url.Values) error {
	if jobs, err := this.GetJobReports(params.Get("job")); err != nil {
		return err
	} else {
		return this.output(jobs)
	}
}
