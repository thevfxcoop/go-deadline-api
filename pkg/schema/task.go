package schema

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/thevfxcoop/go-deadline-api"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Task struct {
	Id            string     `deadline:"_id" json:"_id,omitempty"`
	Job           string     `deadline:"JobID" json:"job"`
	Task          uint       `deadline:"TaskID" json:"task"`
	Worker        string     `deadline:"Slave" json:"worker,omitempty"`
	Frames        string     `deadline:"Frames" json:"frames,omitempty"`
	Status        TaskStatus `deadline:"Stat" json:"status"`
	Progress      string     `deadline:"Prog" json:"progress,omitempty"`
	DateStarted   time.Time  `deadline:"Start" json:"date_started,omitempty"`
	DateCompleted time.Time  `deadline:"Comp" json:"date_completed,omitempty"`
	Errors        uint       `deadline:"Errs" json:"errors"`
	UsedClock     uint       `deadline:"UsedClock" json:"used_clock,omitempty"`
	TotalClock    uint       `deadline:"TotalClock" json:"total_clock,omitempty"`
	MemMaxBytes   uint       `deadline:"RamPeak" json:"mem_max_bytes,omitempty"`
	MemMeanBytes  uint       `deadline:"RamAvg" json:"mem_mean_bytes,omitempty"`
	CPUPercent    uint       `deadline:"Cpu" json:"cpu_percent,omitempty"`
	PreTask       bool       `json:"pre_task,omitempty"`
	PostTask      bool       `json:"post_task,omitempty"`
}

type TaskStatus uint

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	// https://docs.thinkboxsoftware.com/products/deadline/10.1/1_User%20Manual/manual/rest-tasks.html#task-property-values
	TaskStatusUnknown   TaskStatus = 1
	TaskStatusQueued    TaskStatus = 2
	TaskStatusSuspended TaskStatus = 3
	TaskStatusRendering TaskStatus = 4
	TaskStatusCompleted TaskStatus = 5
	TaskStatusFailed    TaskStatus = 6
	TaskStatusPending   TaskStatus = 8
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewTask(v map[string]interface{}) (*Task, error) {
	this := new(Task)

	if err := decoder.Decode(v, this); err != nil {
		return nil, err
	} else {
		return this, nil
	}
}

///////////////////////////////////////////////////////////////////////////////
// MARSHAL AND UNMARSHALL

func (s *TaskStatus) Unmarshal(in interface{}) error {
	switch v := in.(type) {
	case float64:
		*s = TaskStatus(v)
	default:
		return deadline.ErrBadParameter.With("Decode: ", in)
	}
	// return success
	return nil
}

func (s TaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprint(s))
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (v *Task) String() string {
	if data, err := json.MarshalIndent(v, "", "  "); err == nil {
		return string(data)
	} else {
		return err.Error()
	}
}

func (v TaskStatus) String() string {
	switch v {
	case TaskStatusUnknown:
		return "unknown"
	case TaskStatusQueued:
		return "queued"
	case TaskStatusSuspended:
		return "suspended"
	case TaskStatusRendering:
		return "rendering"
	case TaskStatusCompleted:
		return "completed"
	case TaskStatusFailed:
		return "failed"
	case TaskStatusPending:
		return "pending"
	default:
		return "[?? Invalid TaskStatus value]"
	}
}
