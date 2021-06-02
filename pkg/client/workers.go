package client

import (
	"github.com/thevfxcoop/go-deadline-api/pkg/schema"
)

///////////////////////////////////////////////////////////////////////////////
// METHODS

// GetWorkerNames gets all the Worker names.
func (this *Client) GetWorkerNames() ([]string, error) {
	var objs []string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/slaves"), optNameOnly()); err != nil {
		return nil, err
	}

	// Return success
	return objs, nil
}

// GetWorkerInfo gets the InfoSettings for every Worker name provided. If
// no workers provided, gets settings for all workers
func (this *Client) GetWorkerInfo(worker ...string) ([]*schema.WorkerInfo, error) {
	var objs []map[string]interface{}
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/slaves"), optName(worker), optData("infosettings")); err != nil {
		return nil, err
	}

	// Convert into schema
	infos := make([]*schema.WorkerInfo, 0, len(objs))
	for _, obj := range objs {
		if info, err := schema.NewWorkerInfo(obj); err != nil {
			return nil, err
		} else {
			infos = append(infos, info)
		}
	}

	// Return success
	return infos, nil
}

// DeleteWorkers removes one or more workers
func (this *Client) DeleteWorkers(worker ...string) error {
	payload := NewDeletePayload(ContentTypeJson)
	if err := this.Do(payload, nil, OptPath("api/slaves"), optName(worker)); err != nil {
		return err
	}

	// Return success
	return nil
}

// GetWorkerReports returns worker reports
func (this *Client) GetWorkerReports(worker ...string) ([]*schema.WorkerReport, error) {
	var objs []map[string]interface{}
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/slaves"), optName(worker), optData("reports")); err != nil {
		return nil, err
	}

	result := make([]*schema.WorkerReport, 0, len(objs))
	for _, obj := range objs {
		reports, ok := obj["Reps"].([]interface{})
		if !ok {
			continue
		}
		for _, report := range reports {
			if report, err := schema.NewWorkerReport(report.(map[string]interface{})); err != nil {
				return nil, err
			} else {
				result = append(result, report)
			}
		}
	}

	// Return success
	return result, nil
}

// WorkersForJob gets all worker names rendering Job that corresponds to Job ID provided
func (this *Client) WorkersForJob(job string) ([]string, error) {
	var objs []string

	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/slavesrenderingjob"), optJobId(job)); err != nil {
		return nil, err
	}

	// Return success
	return objs, nil
}
