package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

type Jobs struct {
	command
}

func NewJobs(client *client.Client, log *log.Logger) Command {
	this := new(Jobs)
	this.Client = client
	this.Logger = log
	return this
}

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
		fmt.Println(jobs)
	}

	// Return success
	return nil
}
