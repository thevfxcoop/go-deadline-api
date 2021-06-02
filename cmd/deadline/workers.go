package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

type Workers struct {
	command
	run string
}

func NewWorkers(client *client.Client, log *log.Logger) Command {
	this := new(Workers)
	this.Client = client
	this.Logger = log
	return this
}

func (this *Workers) Matches(args []string) url.Values {
	params := url.Values{}
	this.run = args[0]
	if args[0] == "workers" && len(args) == 1 {
		return params
	}
	if args[0] == "workerinfo" && len(args) >= 1 {
		params["worker"] = args[1:]
		return params
	}
	if args[0] == "deleteworker" && len(args) > 1 {
		params["worker"] = args[1:]
		return params
	}
	if args[0] == "workerreport" && len(args) > 1 {
		params["worker"] = args[1:]
		return params
	}
	if args[0] == "workersforjob" && len(args) == 2 {
		params.Set("job", args[1])
		return params
	}
	return nil
}

func (this *Workers) Run(params url.Values) error {
	switch this.run {
	case "workers":
		if workers, err := this.GetWorkerNames(); err != nil {
			return err
		} else if data, err := json.MarshalIndent(workers, "", "  "); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	case "workerinfo":
		if workers, err := this.GetWorkerInfo(params["worker"]...); err != nil {
			return err
		} else if data, err := json.MarshalIndent(workers, "", "  "); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	case "deleteworker":
		if err := this.DeleteWorkers(params["worker"]...); err != nil {
			return err
		} else if workers, err := this.GetWorkerNames(); err != nil {
			return err
		} else if data, err := json.MarshalIndent(workers, "", "  "); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	case "workerreport":
		if reports, err := this.GetWorkerReports(params["worker"]...); err != nil {
			return err
		} else if data, err := json.MarshalIndent(reports, "", "  "); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	case "workersforjob":
		if workers, err := this.WorkersForJob(params.Get("job")); err != nil {
			return err
		} else if data, err := json.MarshalIndent(workers, "", "  "); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	}

	// Return success
	return nil
}
