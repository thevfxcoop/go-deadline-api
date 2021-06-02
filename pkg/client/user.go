package client

import schema "github.com/thevfxcoop/go-deadline-api/pkg/schema"

///////////////////////////////////////////////////////////////////////////////
// METHODS

// GetUserNames gets all the User names.
func (this *Client) GetUserNames() ([]string, error) {
	var objs []string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/users"), optNameOnly()); err != nil {
		return nil, err
	}

	// Return success
	return objs, nil
}

// GetUserInfo gets info for users
func (this *Client) GetUserInfo(users ...string) ([]*schema.User, error) {
	var objs []map[string]interface{}
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/users"), optName(users)); err != nil {
		return nil, err
	}

	// Convert into schema
	result := make([]*schema.User, 0, len(objs))
	for _, obj := range objs {
		if info, err := schema.NewUser(obj); err != nil {
			return nil, err
		} else {
			result = append(result, info)
		}
	}

	// Return success
	return result, nil
}
