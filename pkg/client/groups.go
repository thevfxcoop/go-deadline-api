package client

///////////////////////////////////////////////////////////////////////////////
// METHODS

// GetGroups returns all groups
func (this *Client) GetGroups() ([]string, error) {
	var objs []string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/groups")); err != nil {
		return nil, err
	}

	// Return success
	return objs, nil
}

// GetWorkersForGroup returns all workers in one or more groups
func (this *Client) GetWorkersForGroup(groups ...string) ([]string, error) {
	var objs []string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/groups"), optGroups(groups)); err != nil {
		return nil, err
	}

	// Return success
	return objs, nil
}

// AddGroups adds one or more groups
func (this *Client) AddGroups(groups ...string) error {
	payload := NewGroupPayload(groups, false)
	if err := this.Do(payload, nil, OptPath("api/groups"), optGroups(groups)); err != nil {
		return err
	}

	// Return success
	return nil
}

// DeleteGroups removes one or more groups
func (this *Client) DeleteGroups(groups ...string) error {
	payload := NewDeletePayload(ContentTypeJson)
	if err := this.Do(payload, nil, OptPath("api/groups"), optGroups(groups)); err != nil {
		return err
	}

	// Return success
	return nil
}

// SetGroups removes all groups not provided and creates any provided groups that did not exist
func (this *Client) SetGroups(groups ...string) error {
	payload := NewGroupPayload(groups, true)
	if err := this.Do(payload, nil, OptPath("api/groups"), optGroups(groups)); err != nil {
		return err
	}

	// Return success
	return nil
}

// AddWorkersToGroup adds the provided Groups to the assigned groups for each provided Worker
func (this *Client) AddWorkersToGroup(group string, workers ...string) error {
	payload := NewGroupPayload([]string{group}, false, workers...)
	if err := this.Do(payload, nil, OptPath("api/groups")); err != nil {
		return err
	}

	// Return success
	return nil
}

// RemoveWorkersFromGroup removes workers from a group
func (this *Client) RemoveWorkersFromGroup(group string, workers ...string) error {
	payload := NewDeletePayload(ContentTypeJson)
	if err := this.Do(payload, nil, OptPath("api/groups"), optGroups([]string{group}), optSlaves(workers)); err != nil {
		return err
	}

	// Return success
	return nil
}

// SetGroupsForWorker sets the provided Groups to the assigned groups for a Worker
func (this *Client) SetGroupsForWorker(worker string, groups ...string) error {
	payload := NewGroupPayload(groups, true, worker)
	if err := this.Do(payload, nil, OptPath("api/groups")); err != nil {
		return err
	}

	// Return success
	return nil
}
