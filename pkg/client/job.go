package client

import (
	"time"

	deadline "github.com/thevfxcoop/go-deadline-api"
)

///////////////////////////////////////////////////////////////////////////////
// SCHEMA

type JobState string

type Job struct {
	Id     string `json:"_id"`
	Plugin string `json:"Plug"`

	Props struct {
		Name       string
		Batch      string
		User       string
		Region     string
		Comment    string `json:"Cmmt"`
		Department string `json:"Dept"`
		Frames     string
		Chunk      uint
		Tasks      uint
		Group      string `json:"Grp"`
		Pool       string
		SecPool    string
		Priority   uint `json:"Pri"`
	} `json:"Props"`

	Purged          bool
	Machine         string    `json:"Mach"`
	SubmittedDate   time.Time `json:"Date"`
	StartedDate     time.Time `json:"DateStart"`
	CompletedDate   time.Time `json:"DateComp"`
	CompletedChunks uint
	QueuedChunks    uint
	SuspendedChunks uint
	RenderingChunks uint
	FailedChunks    uint
	PendingChunks   uint
	Progress        string `json:"SnglTskPrg"`
	Errors          uint   `json:"Errs"`
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	JobStateActive    JobState = "Active"
	JobStateSuspended JobState = "Suspended"
	JobStateCompleted JobState = "Completed"
	JobStateFailed    JobState = "Failed"
	JobStatePending   JobState = "Pending"
)

///////////////////////////////////////////////////////////////////////////////
// METHODS

// GetJobs returns jobs with filters. The filters can be OptJobState(...JobState)
func (this *Client) GetJobs(opts ...RequestOpt) ([]Job, error) {
	var jobs []Job
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &jobs, append(opts, OptPath("api/jobs"))...); err != nil {
		return nil, err
	} else {
		return jobs, nil
	}
}

// GetJobWithId returns a job with specific id. Returns nil if
// job is not found
func (this *Client) GetJobWithId(id string) (*Job, error) {
	var jobs []Job
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &jobs, OptPath("api/jobs"), optJobId(id)); err != nil {
		return nil, err
	} else if len(jobs) == 0 {
		return nil, deadline.ErrNotFound
	} else {
		return &jobs[0], nil
	}
}

// GetJobIds returns job identifiers
func (this *Client) GetJobIds() ([]string, error) {
	var jobs []string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &jobs, optIdOnly(true), OptPath("api/jobs")); err != nil {
		return nil, err
	} else {
		return jobs, nil
	}
}

// SuspendJobWithId - Puts the job with the matching ID into the Suspended state.
func (this *Client) SuspendJobWithId(id string) error {
	var status string
	payload := NewJobCommandPayload(id, "suspend")
	if err := this.Do(payload, &status, OptPath("api/jobs")); err != nil {
		return err
	} else if status != "Success" {
		return deadline.ErrUnexpectedResponse.With(status)
	} else {
		return nil
	}
}

// SuspendJobTasksWithId - Puts job tasks with the matching ID into the Suspended state.
func (this *Client) SuspendJobTasksWithId(id string) error {
	var status string
	payload := NewJobCommandPayload(id, "suspendnonrendering")
	if err := this.Do(payload, &status, OptPath("api/jobs")); err != nil {
		return err
	} else if status != "Success" {
		return deadline.ErrUnexpectedResponse.With(status)
	} else {
		return nil
	}
}

// ResumeJobWithId - Resumes the job with the ID that matches the provided ID.
func (this *Client) ResumeJobWithId(id string) error {
	var status string
	payload := NewJobCommandPayload(id, "resume")
	if err := this.Do(payload, &status, OptPath("api/jobs")); err != nil {
		return err
	} else if status != "Success" {
		return deadline.ErrUnexpectedResponse.With(status)
	} else {
		return nil
	}
}

// ResumeFailedJobWithId - Resumes the failed job with the ID that matches the provided ID.
func (this *Client) ResumeFailedJobWithId(id string) error {
	var status string
	payload := NewJobCommandPayload(id, "resumefailed")
	if err := this.Do(payload, &status, OptPath("api/jobs")); err != nil {
		return err
	} else if status != "Success" {
		return deadline.ErrUnexpectedResponse.With(status)
	} else {
		return nil
	}
}

// RequeueJobWithId - Requeues the job with the ID that matches the provided ID.
func (this *Client) RequeueJobWithId(id string) error {
	var status string
	payload := NewJobCommandPayload(id, "requeue")
	if err := this.Do(payload, &status, OptPath("api/jobs")); err != nil {
		return err
	} else if status != "Success" {
		return deadline.ErrUnexpectedResponse.With(status)
	} else {
		return nil
	}
}

// ArchiveJobWithId - Archives the job with the ID that matches the provided ID.
func (this *Client) ArchiveJobWithId(id string) error {
	var status string
	payload := NewJobCommandPayload(id, "archive")
	if err := this.Do(payload, &status, OptPath("api/jobs")); err != nil {
		return err
	} else if status != "Success" {
		return deadline.ErrUnexpectedResponse.With(status)
	} else {
		return nil
	}
}

// PendJobWithId - Puts the job with the ID that matches the provided ID in the pending state.
func (this *Client) PendJobWithId(id string) error {
	var status string
	payload := NewJobCommandPayload(id, "pend")
	if err := this.Do(payload, &status, OptPath("api/jobs")); err != nil {
		return err
	} else if status != "Success" {
		return deadline.ErrUnexpectedResponse.With(status)
	} else {
		return nil
	}
}

// ReleasePendingJobWithId - Puts the job with the ID that matches the provided ID in the pending state.
func (this *Client) ReleasePendingJobWithId(id string) error {
	var status string
	payload := NewJobCommandPayload(id, "releasepending")
	if err := this.Do(payload, &status, OptPath("api/jobs")); err != nil {
		return err
	} else if status != "Success" {
		return deadline.ErrUnexpectedResponse.With(status)
	} else {
		return nil
	}
}

// CompleteJobWithId - Marks the job with the ID that matches the provided ID as complete.
func (this *Client) CompleteJobWithId(id string) error {
	var status string
	payload := NewJobCommandPayload(id, "complete")
	if err := this.Do(payload, &status, OptPath("api/jobs")); err != nil {
		return err
	} else if status != "Success" {
		return deadline.ErrUnexpectedResponse.With(status)
	} else {
		return nil
	}
}

// FailJobWithId - Marks the job with the ID that matches the provided ID as failed.
func (this *Client) FailJobWithId(id string) error {
	var status string
	payload := NewJobCommandPayload(id, "fail")
	if err := this.Do(payload, &status, OptPath("api/jobs")); err != nil {
		return err
	} else if status != "Success" {
		return deadline.ErrUnexpectedResponse.With(status)
	} else {
		return nil
	}
}
