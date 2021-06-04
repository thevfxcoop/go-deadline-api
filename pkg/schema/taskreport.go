package schema

import (
	"encoding/json"
	"fmt"
	"time"

	// Modules
	"github.com/thevfxcoop/go-deadline-api"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type TaskReport struct {
	Id             string         `deadline:"_id" json:"_id"`
	Type           TaskReportType `deadline:"Type" json:"type"`
	Job            string         `deadline:"Job" json:"job"`
	Task           uint           `deadline:"Task" json:"task"`
	Plugin         string         `deadline:"Plugin" json:"plugin"`
	Name           string         `deadline:"JobName" json:"name,omitempty"`
	User           string         `deadline:"JobUser" json:"user,omitempty"`
	Frames         string         `deadline:"Frames" json:"frames,omitempty"`
	Date           time.Time      `deadline:"Date" json:"date,omitempty"`
	DateStarted    time.Time      `deadline:"TaskStartTime" json:"date_start,omitempty"`
	TaskTime       time.Duration  `deadline:"TaskTime" json:"duration_task,omitempty"`
	Worker         string         `deadline:"Slave" json:"slave,omitempty"`
	Data           string         `deadline:"Title" json:"data,omitempty"`
	MemMaxBytes    uint           `deadline:"PeakRam" json:"mem_max_bytes,omitempty"`
	MemMeanBytes   uint           `deadline:"AverageRam" json:"mem_mean_bytes,omitempty"`
	CPUMaxPercent  uint           `deadline:"PeakCpu" json:"cpu_max_percent,omitempty"`
	CPUMeanPercent uint           `deadline:"AverageCpu" json:"cpu_mean_percent,omitempty"`
	UsedClock      uint           `deadline:"UsedClock" json:"used_clock,omitempty"`
	TotalClock     uint           `deadline:"TotalClock" json:"total_clock,omitempty"`
}

type TaskReportType uint

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	// https://docs.thinkboxsoftware.com/products/deadline/10.1/1_User%20Manual/manual/rest-task-reports.html#task-report-property-values
	TaskReportTypeLog     TaskReportType = 0
	TaskReportTypeError   TaskReportType = 1
	TaskReportTypeRequeue TaskReportType = 2
)

///////////////////////////////////////////////////////////////////////////////
// CONVERT JOB

func NewTaskReport(v map[string]interface{}) (*TaskReport, error) {
	this := new(TaskReport)

	if err := decoder.Decode(v, this); err != nil {
		return nil, err
	} else {
		return this, nil
	}
}

///////////////////////////////////////////////////////////////////////////////
// MARSHAL AND UNMARSHALL

func (s *TaskReportType) Unmarshal(in interface{}) error {
	switch v := in.(type) {
	case float64:
		*s = TaskReportType(v)
	default:
		return deadline.ErrBadParameter.With("Decode: ", in)
	}
	// return success
	return nil
}

func (s TaskReportType) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprint(s))
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (v *TaskReport) String() string {
	if data, err := json.MarshalIndent(v, "", "  "); err == nil {
		return string(data)
	} else {
		return err.Error()
	}
}

func (v TaskReportType) String() string {
	switch v {
	case TaskReportTypeLog:
		return "log"
	case TaskReportTypeError:
		return "error"
	case TaskReportTypeRequeue:
		return "requeue"
	default:
		return "[?? Invalid TaskReportType value]"
	}
}
