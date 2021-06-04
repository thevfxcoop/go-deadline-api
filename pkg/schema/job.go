package schema

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/thevfxcoop/go-deadline-api"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Job struct {
	Id            string           `deadline:"_id" json:"_id,omitempty" `
	Plugin        string           `deadline:"Plug" json:"plugin,omitempty"`
	Name          string           `deadline:"Props/Name" json:"name,omitempty"`
	Comment       string           `deadline:"Props/Cmmt" json:"comment,omitempty"`
	Department    string           `deadline:"Props/Dept" json:"department,omitempty"`
	User          string           `deadline:"Props/User" json:"user,omitempty"`
	Workstation   string           `deadline:"Mach" json:"workstation,omitempty"`
	Batch         string           `deadline:"Props/Batch" json:"batch,omitempty"`
	Region        string           `deadline:"Props/Region" json:"region,omitempty"`
	Group         string           `deadline:"Props/Grp" json:"group,omitempty"`
	Pool          string           `deadline:"Props/Pool" json:"pool,omitempty"`
	SecondaryPool string           `deadline:"Props/SecPool" json:"secondary_pool,omitempty"`
	Priority      uint             `deadline:"Props/Pri" json:"pri"`
	ScheduledType JobScheduledType `deadline:"Props/Schd" json:"scheduled_type"`
	OnCompletion  JobOnCompletion  `deadline:"Props/OnComp" json:"on_completion"`
	TaskTimeout   JobTaskTimeout   `deadline:"Props/Timeout" json:"task_timeout"`
	Frames        string           `deadline:"Props/Frames" json:"frames,omitempty"`
	Progress      string           `deadline:"SnglTskPrg" json:"progress,omitempty"`
	Status        JobStatus        `deadline:"Stat" json:"status"`
	DateScheduled time.Time        `deadline:"Props/SchdDate" json:"date_scheduled,omitempty"`
	DateSubmitted time.Time        `deadline:"Date" json:"date_submitted,omitempty"`
	DateStarted   time.Time        `deadline:"DateStart" json:"date_started,omitempty"`
	DateCompleted time.Time        `deadline:"DateComp" json:"date_completed,omitempty"`
	Purged        bool             `deadline:"Purged" json:"purged"`
	Chunks        uint             `deadline:"Chunk" json:"chunks"`
	Tasks         uint             `deadline:"Props/Tasks" json:"tasks"`
	Errors        uint             `deadline:"Errs" json:"errors"`
	Pending       uint             `deadline:"PendingChunks" json:"pending"`
	Queued        uint             `deadline:"QueuedChunks" json:"queued"`
	Suspended     uint             `deadline:"SuspendedChunks" json:"suspended"`
	Rendering     uint             `deadline:"RenderingChunks" json:"rendering"`
	Completed     uint             `deadline:"CompletedChunks" json:"completed"`
	Failed        uint             `deadline:"FailedChunks" json:"failed"`
}

type JobStatus uint
type JobTaskTimeout uint
type JobOnCompletion uint
type JobScheduledType uint

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	StringNone = "none"
)

const (
	// https://docs.thinkboxsoftware.com/products/deadline/10.1/1_User%20Manual/manual/rest-jobs.html#job-property-values
	JobStatusUnknown   JobStatus = 0
	JobStatusActive    JobStatus = 1
	JobStatusSuspended JobStatus = 2
	JobStatusCompleted JobStatus = 3
	JobStatusFailed    JobStatus = 4
	JobStatusPending   JobStatus = 6
)

const (
	// https://docs.thinkboxsoftware.com/products/deadline/10.1/1_User%20Manual/manual/rest-jobs.html#job-property-values
	JobTaskTimeoutBoth   JobTaskTimeout = 0
	JobTaskTimeoutError  JobTaskTimeout = 1
	JobTaskTimeoutNotify JobTaskTimeout = 2
)

const (
	// https://docs.thinkboxsoftware.com/products/deadline/10.1/1_User%20Manual/manual/rest-jobs.html#job-property-values
	JobOnCompletionArchive JobOnCompletion = 0
	JobOnCompletionDelete  JobOnCompletion = 1
	JobOnCompletionNone    JobOnCompletion = 2
)

const (
	// https://docs.thinkboxsoftware.com/products/deadline/10.1/1_User%20Manual/manual/rest-jobs.html#job-property-values
	JobScheduledTypeNone  JobScheduledType = 0
	JobScheduledTypeOnce  JobScheduledType = 1
	JobScheduledTypeDaily JobScheduledType = 2
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewJob(v map[string]interface{}) (*Job, error) {
	this := new(Job)

	if err := decoder.Decode(v, this); err != nil {
		return nil, err
	} else {
		return this, nil
	}
}

///////////////////////////////////////////////////////////////////////////////
// MARSHAL AND UNMARSHALL

func (s *JobStatus) Unmarshal(in interface{}) error {
	switch v := in.(type) {
	case float64:
		*s = JobStatus(v)
	default:
		return deadline.ErrBadParameter.With("Decode: ", in)
	}
	// return success
	return nil
}

func (s JobStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprint(s))
}

func (s *JobOnCompletion) Unmarshal(in interface{}) error {
	switch v := in.(type) {
	case float64:
		*s = JobOnCompletion(v)
	default:
		return deadline.ErrBadParameter.With("Decode: ", in)
	}
	// return success
	return nil
}

func (s JobOnCompletion) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprint(s))
}

func (s *JobScheduledType) Unmarshal(in interface{}) error {
	switch v := in.(type) {
	case float64:
		*s = JobScheduledType(v)
	default:
		return deadline.ErrBadParameter.With("Decode: ", in)
	}
	// return success
	return nil
}

func (s JobScheduledType) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprint(s))
}

func (s *JobTaskTimeout) Unmarshal(in interface{}) error {
	switch v := in.(type) {
	case float64:
		*s = JobTaskTimeout(v)
	default:
		return deadline.ErrBadParameter.With("Decode: ", in)
	}
	// return success
	return nil
}

func (s JobTaskTimeout) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprint(s))
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (v *Job) String() string {
	if data, err := json.MarshalIndent(v, "", "  "); err == nil {
		return string(data)
	} else {
		return err.Error()
	}
}

func (v JobStatus) String() string {
	switch v {
	case JobStatusUnknown:
		return "unknown"
	case JobStatusActive:
		return "active"
	case JobStatusSuspended:
		return "suspended"
	case JobStatusCompleted:
		return "completed"
	case JobStatusFailed:
		return "failed"
	case JobStatusPending:
		return "pending"
	default:
		return "[?? Invalid JobStatus value]"
	}
}

func (v JobTaskTimeout) String() string {
	switch v {
	case JobTaskTimeoutBoth:
		return "both"
	case JobTaskTimeoutError:
		return "error"
	case JobTaskTimeoutNotify:
		return "notify"
	default:
		return "[?? Invalid JobTaskTimeout value]"
	}
}

func (v JobOnCompletion) String() string {
	switch v {
	case JobOnCompletionArchive:
		return "archive"
	case JobOnCompletionDelete:
		return "delete"
	case JobOnCompletionNone:
		return "none"
	default:
		return "[?? Invalid JobOnCompletion value]"
	}
}

func (v JobScheduledType) String() string {
	switch v {
	case JobScheduledTypeNone:
		return "none"
	case JobScheduledTypeOnce:
		return "once"
	case JobScheduledTypeDaily:
		return "daily"
	default:
		return "[?? Invalid JobScheduledType value]"
	}
}
