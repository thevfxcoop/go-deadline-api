package main

import (
	"net/url"
	"strconv"

	// Modules
	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Tasks struct {
	command
	fn func(url.Values) error
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewTasks(client *client.Client) Command {
	this := new(Tasks)
	this.Client = client
	return this
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Tasks) Matches(args []string) url.Values {
	params := url.Values{}
	if args[0] == "tasks" && len(args) == 2 {
		params.Set("job", args[1])
		this.fn = this.RunTaskId
		return params
	}
	if args[0] == "taskinfo" && len(args) == 2 {
		params.Set("job", args[1])
		this.fn = this.RunTaskInfo
		return params
	}
	if args[0] == "requeue" && len(args) > 1 {
		params.Set("job", args[1])
		params["tasks"] = args[2:]
		this.fn = this.RunRequeueTasks
		return params
	}
	return nil
}

func (this *Tasks) Run(params url.Values) error {
	return this.fn(params)
}

func (this *Tasks) RunTaskId(params url.Values) error {
	if tasks, err := this.GetTaskIdsForJob(params.Get("job")); err != nil {
		return err
	} else {
		return this.output(tasks)
	}
}

func (this *Tasks) RunTaskInfo(params url.Values) error {
	if tasks, err := this.GetTasksForJob(params.Get("job")); err != nil {
		return err
	} else {
		return this.output(tasks)
	}
}

func (this *Tasks) RunRequeueTasks(params url.Values) error {
	if tasks, err := arrayStringToUint(params["tasks"]); err != nil {
		return err
	} else if err := this.RequeueTasksWithId(params.Get("job"), tasks...); err != nil {
		return err
	} else if tasks, err := this.GetTasksForJob(params.Get("job")); err != nil {
		return err
	} else {
		return this.output(tasks)
	}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func arrayStringToUint(v []string) ([]uint, error) {
	result := make([]uint, 0, len(v))
	for _, value := range v {
		if v_, err := strconv.ParseUint(value, 0, 32); err != nil {
			return nil, err
		} else {
			result = append(result, uint(v_))
		}
	}
	return result, nil
}
