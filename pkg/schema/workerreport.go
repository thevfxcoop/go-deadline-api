package schema

import (
	"encoding/json"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type WorkerReport struct {
	Id             string        `deadline:"_id" json:"_id"`
	Type           JobReportType `deadline:"Type" json:"type"`
	Plugin         string        `deadline:"Plugin" json:"plugin,omitempty"`
	Job            string        `deadline:"Job" json:"job"`
	Name           string        `deadline:"JobName" json:"name,omitempty"`
	User           string        `deadline:"JobUser" json:"user,omitempty"`
	Frames         string        `deadline:"Frames" json:"frames,omitempty"`
	Task           uint          `deadline:"Task" json:"task"`
	Slave          string        `deadline:"Slave" json:"worker,omitempty"`
	Data           string        `deadline:"Title" json:"data,omitempty"`
	Date           time.Time     `deadline:"Date" json:"date,omitempty"`
	DateStarted    time.Time     `deadline:"TaskStartTime" json:"date_start,omitempty"`
	TaskTime       time.Duration `deadline:"TaskTime" json:"duration_task,omitempty"`
	MemMaxBytes    uint          `deadline:"PeakRam" json:"mem_max_bytes,omitempty"`
	MemMeanBytes   uint          `deadline:"AverageRam" json:"mem_mean_bytes,omitempty"`
	CPUMaxPercent  uint          `deadline:"PeakCpu" json:"cpu_max_percent,omitempty"`
	CPUMeanPercent uint          `deadline:"AverageCpu" json:"cpu_mean_percent,omitempty"`
	UsedClock      uint          `deadline:"UsedClock" json:"used_clock,omitempty"`
	TotalClock     uint          `deadline:"TotalClock" json:"total_clock,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewWorkerReport(v map[string]interface{}) (*WorkerReport, error) {
	this := new(WorkerReport)

	if err := decoder.Decode(v, this); err != nil {
		return nil, err
	} else {
		return this, nil
	}
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (v *WorkerReport) String() string {
	if data, err := json.MarshalIndent(v, "", "  "); err == nil {
		return string(data)
	} else {
		return err.Error()
	}
}
