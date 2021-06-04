package client

import (
	"github.com/thevfxcoop/go-deadline-api/pkg/schema"
)

///////////////////////////////////////////////////////////////////////////////
// METHODS

// GetTaskReports returns task reports
func (this *Client) GetTaskReports(job string, task uint) ([]*schema.TaskReport, error) {
	var objs []map[string]interface{}
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/taskreports"), optJobId(job), optTaskId(task), optData("all")); err != nil {
		return nil, err
	}

	// Convert into schema
	result := make([]*schema.TaskReport, 0, len(objs))
	for _, obj := range objs {
		if taskreport, err := schema.NewTaskReport(obj); err != nil {
			return nil, err
		} else {
			result = append(result, taskreport)
		}
	}

	// Return success
	return result, nil
}
