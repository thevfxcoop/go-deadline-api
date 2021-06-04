package client

import schema "github.com/thevfxcoop/go-deadline-api/pkg/schema"

///////////////////////////////////////////////////////////////////////////////
// METHODS

// GetPulseNames returns all the pulse names
func (this *Client) GetPulseNames() ([]string, error) {
	var objs []string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/pulse"), optNameOnly()); err != nil {
		return nil, err
	}

	// Return success
	return objs, nil
}

// GetPulseInfo returns pulse information
func (this *Client) GetPulseInfo(names ...string) ([]*schema.Pulse, error) {
	var objs []map[string]interface{}
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/pulse"), optNames(names), optInfo()); err != nil {
		return nil, err
	}

	// Convert into schema
	result := make([]*schema.Pulse, 0, len(objs))
	for _, obj := range objs {
		if info, err := schema.NewPulse(obj); err != nil {
			return nil, err
		} else {
			result = append(result, info)
		}
	}

	// Return success
	return result, nil
}

// GetPulseSettings returns pulse settings
func (this *Client) GetPulseSettings(names ...string) ([]*schema.PulseSettings, error) {
	var objs []map[string]interface{}
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/pulse"), optNames(names), optSettings()); err != nil {
		return nil, err
	}

	// Convert into schema
	result := make([]*schema.PulseSettings, 0, len(objs))
	for _, obj := range objs {
		if info, err := schema.NewPulseSettings(obj); err != nil {
			return nil, err
		} else {
			result = append(result, info)
		}
	}

	// Return success
	return result, nil
}
