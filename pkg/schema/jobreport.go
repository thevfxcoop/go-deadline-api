package schema

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/thevfxcoop/go-deadline-api"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type JobReport struct {
	Id             string        `deadline:"_id" json:"_id"`
	Type           JobReportType `deadline:"Type" json:"type"`
	Job            string        `deadline:"Job" json:"job"`
	Task           uint          `deadline:"Task" json:"task"`
	Plugin         string        `deadline:"Plugin" json:"plugin"`
	Name           string        `deadline:"JobName" json:"name,omitempty"`
	User           string        `deadline:"JobUser" json:"user,omitempty"`
	Frames         string        `deadline:"Frames" json:"frames,omitempty"`
	Date           time.Time     `deadline:"Date" json:"date,omitempty"`
	DateStarted    time.Time     `deadline:"TaskStartTime" json:"date_start,omitempty"`
	TaskTime       time.Duration `deadline:"TaskTime" json:"duration_task,omitempty"`
	Slave          string        `deadline:"Slave" json:"slave,omitempty"`
	Data           string        `deadline:"Title" json:"data,omitempty"`
	MemMaxBytes    uint          `deadline:"PeakRam" json:"mem_max_bytes,omitempty"`
	MemMeanBytes   uint          `deadline:"AverageRam" json:"mem_mean_bytes,omitempty"`
	CPUMaxPercent  uint          `deadline:"PeakCpu" json:"cpu_max_percent,omitempty"`
	CPUMeanPercent uint          `deadline:"AverageCpu" json:"cpu_mean_percent,omitempty"`
	UsedClock      uint          `deadline:"UsedClock" json:"used_clock,omitempty"`
	TotalClock     uint          `deadline:"TotalClock" json:"total_clock,omitempty"`
}

type JobReportType uint

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	JobReportTypeLog     JobReportType = 0
	JobReportTypeError   JobReportType = 1
	JobReportTypeRequeue JobReportType = 2
)

///////////////////////////////////////////////////////////////////////////////
// CONVERT JOB

func NewJobReport(v map[string]interface{}) (*JobReport, error) {
	this := new(JobReport)

	if err := decoder.Decode(v, this); err != nil {
		return nil, err
	} else {
		return this, nil
	}
}

///////////////////////////////////////////////////////////////////////////////
// MARSHAL AND UNMARSHALL

func (s *JobReportType) Unmarshal(in interface{}) error {
	switch v := in.(type) {
	case float64:
		*s = JobReportType(v)
	default:
		return deadline.ErrBadParameter.With("Decode: ", in)
	}
	// return success
	return nil
}

func (s JobReportType) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprint(s))
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (v *JobReport) String() string {
	if data, err := json.MarshalIndent(v, "", "  "); err == nil {
		return string(data)
	} else {
		return err.Error()
	}
}

func (v JobReportType) String() string {
	switch v {
	case JobReportTypeLog:
		return "log"
	case JobReportTypeError:
		return "error"
	case JobReportTypeRequeue:
		return "requeue"
	default:
		return "[?? Invalid JobReportType value]"
	}
}
