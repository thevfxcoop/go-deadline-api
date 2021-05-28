package client

// Ping returns the API version or an error if the API could
// not be reached
func (this *Client) Ping() (string, error) {
	var version string
	payload := NewGetPayload(ContentTypeText)
	if err := this.Do(payload, &version); err != nil {
		return "", err
	} else {
		return version, nil
	}
}
