package client

import "time"

///////////////////////////////////////////////////////////////////////////////
// SCHEMA

type Task struct {
	JobId         string
	TaskId        uint
	Frames        string
	Slave         string
	Status        TaskStatus `json:"Stat"`
	Progress      string     `json:"Prog"`
	Errors        uint       `json:"Errs"`
	SubmittedDate time.Time  `json:"StartDate"`
	StartDate     time.Time  `json:"StartRen"`
	CompletedDate time.Time  `json:"Comp"`
	PreTask       bool
	PostTask      bool
}

type TaskStatus uint

type taskList struct {
	JobID    string `json:"ID"`
	PreTask  *Task
	PostTask *Task
	Tasks    []*Task
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	TaskStatusNone      TaskStatus = 0
	TaskStatusUnknown   TaskStatus = 1
	TaskStatusQueued    TaskStatus = 2
	TaskStatusSuspended TaskStatus = 3
	TaskStatusRendering TaskStatus = 4
	TaskStatusCompleted TaskStatus = 5
	TaskStatusFailed    TaskStatus = 6
	TaskStatusPending   TaskStatus = 8
)

///////////////////////////////////////////////////////////////////////////////
// METHODS

// GetTaskIdsForJobId returns task id's for a job id
func (this *Client) GetTaskIdsForJobId(id string) ([]uint, error) {
	var tasks []uint
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &tasks, OptPath("api/tasks"), optJobId(id), optIdOnly(true)); err != nil {
		return nil, err
	} else {
		return tasks, nil
	}
}

// GetTasksForJobId returns all tasks for a job id
func (this *Client) GetTasksForJobId(id string) ([]*Task, error) {
	var tasks taskList
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &tasks, OptPath("api/tasks"), optJobId(id)); err != nil {
		return nil, err
	}

	// Prepend the pretask and append the posttask
	if tasks.PreTask != nil {
		tasks.PreTask.PreTask = true
		tasks.Tasks = append([]*Task{tasks.PreTask}, tasks.Tasks...)
	}
	if tasks.PostTask != nil {
		tasks.PostTask.PostTask = true
		tasks.Tasks = append(tasks.Tasks, tasks.PostTask)
	}

	// Return tasks
	return tasks.Tasks, nil
}

// GetTaskWithId returns a task for a job id
func (this *Client) GetTaskWithId(id string, task uint) (*Task, error) {
	var obj Task
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &obj, OptPath("api/tasks"), optJobId(id), optTaskId(task)); err != nil {
		return nil, err
	} else {
		return &obj, nil
	}
}

// RequeueTasksWithId requeues the Tasks that correspond to the Task IDs provided
// for the Job that corresponds to the Job ID provided. If no Task IDs are provided,
// all Job tasks will be requeued
func (this *Client) RequeueTasksWithId(id string, tasks ...uint) error {
	payload := NewTasksCommandPayload(id, "requeue", tasks)
	if err := this.Do(payload, nil, OptPath("api/tasks"), optJobId(id)); err != nil {
		return err
	}
	// Return success
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (v TaskStatus) String() string {
	switch v {
	case TaskStatusNone:
		return "TaskStatusNone"
	case TaskStatusUnknown:
		return "TaskStatusUnknown"
	case TaskStatusQueued:
		return "TaskStatusQueued"
	case TaskStatusSuspended:
		return "TaskStatusSuspended"
	case TaskStatusRendering:
		return "TaskStatusRendering"
	case TaskStatusCompleted:
		return "TaskStatusCompleted"
	case TaskStatusFailed:
		return "TaskStatusFailed"
	case TaskStatusPending:
		return "TaskStatusPending"
	default:
		return "[?? Invalid TaskStatus value]"
	}
}
