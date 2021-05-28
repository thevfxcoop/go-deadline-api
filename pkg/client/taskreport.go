package client

import "time"

///////////////////////////////////////////////////////////////////////////////
// SCHEMA

type TaskReportType uint

type TaskReport struct {
	JobId      string    `json:"Job"`
	TaskId     uint      `json:"Task"`
	ReportDate time.Time `json:"Date"`
	StartDate  time.Time `json:"TaskStartTime"`
	Slave      string
	Frames     string
	Name       string `json:"TaskName"`
	User       string `json:"JobUser"`
	Data       string `json:"Title"`
	Plugin     string
	Type       TaskReportType
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	TaskReportLog     TaskReportType = 0
	TaskReportError   TaskReportType = 1
	TaskReportRequeue TaskReportType = 2
)

///////////////////////////////////////////////////////////////////////////////
// METHODS

// GetTaskReportsForJobId returns all task reports for a job id
func (this *Client) GetTaskReportsForJobId(id string, task uint) ([]*TaskReport, error) {
	var objs []*TaskReport
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/taskreports"), optJobId(id), optTaskId(task)); err != nil {
		return nil, err
	}
	return objs, nil
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (v TaskReportType) String() string {
	switch v {
	case TaskReportLog:
		return "TaskReportLog"
	case TaskReportError:
		return "TaskReportError"
	case TaskReportRequeue:
		return "TaskReportRequeue"
	default:
		return "[?? Invalid TaskReportType value]"
	}
}
