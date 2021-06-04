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
	if args[0] == "complete" && len(args) > 1 {
		params.Set("job", args[1])
		params["tasks"] = args[2:]
		this.fn = this.RunCompleteTasks
		return params
	}
	if args[0] == "suspend" && len(args) > 1 {
		params.Set("job", args[1])
		params["tasks"] = args[2:]
		this.fn = this.RunSuspendTasks
		return params
	}
	if args[0] == "fail" && len(args) > 1 {
		params.Set("job", args[1])
		params["tasks"] = args[2:]
		this.fn = this.RunFailTasks
		return params
	}
	if args[0] == "resume" && len(args) > 1 {
		params.Set("job", args[1])
		params["tasks"] = args[2:]
		this.fn = this.RunResumeTasks
		return params
	}
	if args[0] == "pend" && len(args) > 1 {
		params.Set("job", args[1])
		params["tasks"] = args[2:]
		this.fn = this.RunPendTasks
		return params
	}
	if args[0] == "release" && len(args) > 1 {
		params.Set("job", args[1])
		params["tasks"] = args[2:]
		this.fn = this.RunReleaseTasks
		return params
	}
	if args[0] == "taskreports" && len(args) == 3 {
		params.Set("job", args[1])
		params.Set("task", args[2])
		this.fn = this.RunTaskReports
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

func (this *Tasks) RunCompleteTasks(params url.Values) error {
	if tasks, err := arrayStringToUint(params["tasks"]); err != nil {
		return err
	} else if err := this.CompleteTasksWithId(params.Get("job"), tasks...); err != nil {
		return err
	} else if tasks, err := this.GetTasksForJob(params.Get("job")); err != nil {
		return err
	} else {
		return this.output(tasks)
	}
}

func (this *Tasks) RunSuspendTasks(params url.Values) error {
	if tasks, err := arrayStringToUint(params["tasks"]); err != nil {
		return err
	} else if err := this.SuspendTasksWithId(params.Get("job"), tasks...); err != nil {
		return err
	} else if tasks, err := this.GetTasksForJob(params.Get("job")); err != nil {
		return err
	} else {
		return this.output(tasks)
	}
}

func (this *Tasks) RunFailTasks(params url.Values) error {
	if tasks, err := arrayStringToUint(params["tasks"]); err != nil {
		return err
	} else if err := this.FailTasksWithId(params.Get("job"), tasks...); err != nil {
		return err
	} else if tasks, err := this.GetTasksForJob(params.Get("job")); err != nil {
		return err
	} else {
		return this.output(tasks)
	}
}

func (this *Tasks) RunResumeTasks(params url.Values) error {
	if tasks, err := arrayStringToUint(params["tasks"]); err != nil {
		return err
	} else if err := this.ResumeFailedTasksWithId(params.Get("job"), tasks...); err != nil {
		return err
	} else if tasks, err := this.GetTasksForJob(params.Get("job")); err != nil {
		return err
	} else {
		return this.output(tasks)
	}
}

func (this *Tasks) RunPendTasks(params url.Values) error {
	if tasks, err := arrayStringToUint(params["tasks"]); err != nil {
		return err
	} else if err := this.PendTasksWithId(params.Get("job"), tasks...); err != nil {
		return err
	} else if tasks, err := this.GetTasksForJob(params.Get("job")); err != nil {
		return err
	} else {
		return this.output(tasks)
	}
}

func (this *Tasks) RunReleaseTasks(params url.Values) error {
	if tasks, err := arrayStringToUint(params["tasks"]); err != nil {
		return err
	} else if err := this.ReleasePendingTasksWithId(params.Get("job"), tasks...); err != nil {
		return err
	} else if tasks, err := this.GetTasksForJob(params.Get("job")); err != nil {
		return err
	} else {
		return this.output(tasks)
	}
}

func (this *Tasks) RunTaskReports(params url.Values) error {
	if task, err := strconv.ParseUint(params.Get("task"), 0, 32); err != nil {
		return err
	} else if reports, err := this.GetTaskReports(params.Get("job"), uint(task)); err != nil {
		return err
	} else {
		return this.output(reports)
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
