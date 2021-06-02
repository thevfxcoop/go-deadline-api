package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

type JobReports struct {
	command
}

func NewJobReports(client *client.Client, log *log.Logger) Command {
	this := new(JobReports)
	this.Client = client
	this.Logger = log
	return this
}

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
		fmt.Println(jobs)
	}

	// Return success
	return nil
}
