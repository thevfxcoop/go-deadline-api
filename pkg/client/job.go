package client

import (
	deadline "github.com/thevfxcoop/go-deadline-api"
	schema "github.com/thevfxcoop/go-deadline-api/pkg/schema"
)

///////////////////////////////////////////////////////////////////////////////
// METHODS

// GetJobs returns jobs with filters. The filters can be OptJobState(...JobState)
func (this *Client) GetJobs(opts ...RequestOpt) ([]*schema.Job, error) {
	var objs []map[string]interface{}
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, append(opts, OptPath("api/jobs"))...); err != nil {
		return nil, err
	}

	// Convert into schema
	jobs := make([]*schema.Job, 0, len(objs))
	for _, obj := range objs {
		if job, err := schema.NewJob(obj); err != nil {
			return nil, err
		} else {
			jobs = append(jobs, job)
		}
	}

	// Return success
	return jobs, nil
}

// GetJobWithId returns a job with specific id. Returns nil if
// job is not found
func (this *Client) GetJobWithId(id string) (*schema.Job, error) {
	var objs []map[string]interface{}
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/jobs"), optJobId(id)); err != nil {
		return nil, err
	} else if len(objs) == 0 {
		return nil, deadline.ErrNotFound
	}

	// Decode data
	return schema.NewJob(objs[0])
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
