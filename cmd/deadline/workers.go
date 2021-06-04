package main

import (
	"net/url"

	// Modules
	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Workers struct {
	command
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewWorkers(client *client.Client) Command {
	return &Workers{command{Client: client}}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Workers) Matches(args []string) (fn, url.Values) {
	params := url.Values{}
	if args[0] == "workers" && len(args) == 1 {
		return this.RunWorkerNames, params
	}
	if args[0] == "workerinfo" && len(args) >= 1 {
		params["worker"] = args[1:]
		return this.RunWorkerInfo, params
	}
	if args[0] == "deleteworker" && len(args) > 1 {
		params["worker"] = args[1:]
		return this.RunDeleteWorker, params
	}
	if args[0] == "workerreport" && len(args) > 1 {
		params["worker"] = args[1:]
		return this.RunWorkerReports, params
	}
	if args[0] == "workersforjob" && len(args) == 2 {
		params.Set("job", args[1])
		return this.RunWorkersForJob, params
	}
	return nil, nil
}

func (this *Workers) RunWorkerNames(params url.Values) error {
	if workers, err := this.GetWorkerNames(); err != nil {
		return err
	} else {
		return this.output(workers)
	}
}

func (this *Workers) RunWorkerInfo(params url.Values) error {
	if workers, err := this.GetWorkerInfo(params["worker"]...); err != nil {
		return err
	} else {
		return this.output(workers)
	}
}

func (this *Workers) RunDeleteWorker(params url.Values) error {
	if err := this.DeleteWorkers(params["worker"]...); err != nil {
		return err
	} else if workers, err := this.GetWorkerNames(); err != nil {
		return err
	} else {
		return this.output(workers)
	}
}

func (this *Workers) RunWorkerReports(params url.Values) error {
	if reports, err := this.GetWorkerReports(params["worker"]...); err != nil {
		return err
	} else {
		return this.output(reports)
	}
}

func (this *Workers) RunWorkersForJob(params url.Values) error {
	if workers, err := this.WorkersForJob(params.Get("job")); err != nil {
		return err
	} else {
		return this.output(workers)
	}
}
