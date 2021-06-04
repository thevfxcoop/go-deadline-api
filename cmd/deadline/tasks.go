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
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewTasks(client *client.Client) Command {
	return &Tasks{command{Client: client}}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Tasks) Matches(args []string) (fn, url.Values) {
	params := url.Values{}
	if args[0] == "tasks" && len(args) == 2 {
		params.Set("job", args[1])
		return this.RunTaskId, params
	}
	if args[0] == "taskinfo" && len(args) == 2 {
		params.Set("job", args[1])
		return this.RunTaskInfo, params
	}
	if args[0] == "requeue" && len(args) > 1 {
		params.Set("job", args[1])
		params["tasks"] = args[2:]
		return this.RunRequeueTasks, params
	}
	if args[0] == "complete" && len(args) > 1 {
		params.Set("job", args[1])
		params["tasks"] = args[2:]
		return this.RunCompleteTasks, params
	}
	if args[0] == "suspend" && len(args) > 1 {
		params.Set("job", args[1])
		params["tasks"] = args[2:]
		return this.RunSuspendTasks, params
	}
	if args[0] == "fail" && len(args) > 1 {
		params.Set("job", args[1])
		params["tasks"] = args[2:]
		return this.RunFailTasks, params
	}
	if args[0] == "resume" && len(args) > 1 {
		params.Set("job", args[1])
		params["tasks"] = args[2:]
		return this.RunResumeTasks, params
	}
	if args[0] == "pend" && len(args) > 1 {
		params.Set("job", args[1])
		params["tasks"] = args[2:]
		return this.RunPendTasks, params
	}
	if args[0] == "release" && len(args) > 1 {
		params.Set("job", args[1])
		params["tasks"] = args[2:]
		return this.RunReleaseTasks, params
	}
	if args[0] == "taskreports" && len(args) == 3 {
		params.Set("job", args[1])
		params.Set("task", args[2])
		return this.RunTaskReports, params
	}
	if args[0] == "taskreportcontents" && len(args) == 3 {
		params.Set("job", args[1])
		params.Set("task", args[2])
		return this.RunTaskReports, params
	}
	return nil, nil
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

func (this *Tasks) RunTaskReportContentas(params url.Values) error {
	if task, err := strconv.ParseUint(params.Get("task"), 0, 32); err != nil {
		return err
	} else if reports, err := this.GetTaskReportContents(params.Get("job"), uint(task)); err != nil {
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
