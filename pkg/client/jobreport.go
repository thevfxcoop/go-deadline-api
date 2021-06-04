package client

import (
	"github.com/thevfxcoop/go-deadline-api/pkg/schema"
)

///////////////////////////////////////////////////////////////////////////////
// METHODS

// GetJobReports returns job reports with filters.
func (this *Client) GetJobReports(id string, opts ...RequestOpt) ([]*schema.JobReport, error) {
	var objs []map[string]interface{}
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, append(opts, OptPath("api/jobreports"), optJobId(id))...); err != nil {
		return nil, err
	}

	// Convert into schema
	jobreports := make([]*schema.JobReport, 0, len(objs))
	for _, obj := range objs {
		if jobreport, err := schema.NewJobReport(obj); err != nil {
			return nil, err
		} else {
			jobreports = append(jobreports, jobreport)
		}
	}

	// Return success
	return jobreports, nil
}
