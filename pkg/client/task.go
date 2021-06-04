package client

import (
	schema "github.com/thevfxcoop/go-deadline-api/pkg/schema"
)

///////////////////////////////////////////////////////////////////////////////
// SCHEMA

type taskList struct {
	JobID    string `json:"ID"`
	PreTask  map[string]interface{}
	PostTask map[string]interface{}
	Tasks    []map[string]interface{}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

// GetTaskIdsForJob returns task id's for a job
func (this *Client) GetTaskIdsForJob(job string) ([]uint, error) {
	var tasks []uint
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &tasks, OptPath("api/tasks"), optJobId(job), optIdOnly(true)); err != nil {
		return nil, err
	} else {
		return tasks, nil
	}
}

// GetTasksForJob returns all tasks for a job
func (this *Client) GetTasksForJob(job string) ([]*schema.Task, error) {
	var obj taskList
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &obj, OptPath("api/tasks"), optJobId(job)); err != nil {
		return nil, err
	}

	result := make([]*schema.Task, 0, len(obj.Tasks))

	// Prepend the pretask
	if obj.PreTask != nil {
		if task, err := schema.NewTask(obj.PreTask); err != nil {
			return nil, err
		} else {
			task.PreTask = true
			result = append(result, task)
		}
	}

	// Add tasks
	for _, obj := range obj.Tasks {
		if task, err := schema.NewTask(obj); err != nil {
			return nil, err
		} else {
			result = append(result, task)
		}
	}

	// Append the posttask
	if obj.PostTask != nil {
		if task, err := schema.NewTask(obj.PostTask); err != nil {
			return nil, err
		} else {
			task.PostTask = true
			result = append(result, task)
		}
	}

	// Return success
	return result, nil
}

// GetTaskWithId returns a task for a job id
func (this *Client) GetTaskWithId(job string, task uint) (*schema.Task, error) {
	var obj map[string]interface{}
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &obj, OptPath("api/tasks"), optJobId(job), optTaskId(task)); err != nil {
		return nil, err
	} else {
		return schema.NewTask(obj)
	}
}

// RequeueTasksWithId requeues the Tasks that correspond to the Task IDs provided
// If no Task IDs are provided, all tasks will be requeued
func (this *Client) RequeueTasksWithId(id string, tasks ...uint) error {
	payload := NewTasksCommandPayload(id, "requeue", tasks)
	if err := this.Do(payload, nil, OptPath("api/tasks")); err != nil {
		return err
	}
	// Return success
	return nil
}
