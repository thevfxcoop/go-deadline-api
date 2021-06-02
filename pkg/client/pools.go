package client

///////////////////////////////////////////////////////////////////////////////
// METHODS

// GetPools returns all pools
func (this *Client) GetPools() ([]string, error) {
	var objs []string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/pools")); err != nil {
		return nil, err
	}

	// Return success
	return objs, nil
}

// GetWorkersForPool returns all workers in one or more pools
func (this *Client) GetWorkersForPool(pools ...string) ([]string, error) {
	var objs []string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &objs, OptPath("api/pools"), optPools(pools)); err != nil {
		return nil, err
	}

	// Return success
	return objs, nil
}

// AddPools adds one or more pools
func (this *Client) AddPools(pools ...string) error {
	payload := NewPoolPayload(pools, false)
	if err := this.Do(payload, nil, OptPath("api/pools"), optPools(pools)); err != nil {
		return err
	}

	// Return success
	return nil
}

// DeletePools removes one or more pools
func (this *Client) DeletePools(pools ...string) error {
	payload := NewDeletePayload(ContentTypeJson)
	if err := this.Do(payload, nil, OptPath("api/pools"), optPools(pools)); err != nil {
		return err
	}

	// Return success
	return nil
}
