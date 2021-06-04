package client

///////////////////////////////////////////////////////////////////////////////
// METHODS

// GetRepositoryRoot returns root directory for repository
func (this *Client) GetRepositoryRoot() (string, error) {
	var obj string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &obj, OptPath("api/repository"), optDirectory("root")); err != nil {
		return "", err
	} else {
		return obj, nil
	}
}

// GetRepositoryBin returns bin directory for repository
func (this *Client) GetRepositoryBin() (string, error) {
	var obj string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &obj, OptPath("api/repository"), optDirectory("bin")); err != nil {
		return "", err
	} else {
		return obj, nil
	}
}

// GetRepositorySettings returns settings directory for repository
func (this *Client) GetRepositorySettings() (string, error) {
	var obj string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &obj, OptPath("api/repository"), optDirectory("settings")); err != nil {
		return "", err
	} else {
		return obj, nil
	}
}

// GetRepositoryEvents returns events directory for repository
func (this *Client) GetRepositoryEvents() (string, error) {
	var obj string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &obj, OptPath("api/repository"), optDirectory("events")); err != nil {
		return "", err
	} else {
		return obj, nil
	}
}

// GetRepositoryEvents returns custom events directory for repository
func (this *Client) GetRepositoryCustomEvents() (string, error) {
	var obj string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &obj, OptPath("api/repository"), optDirectory("customevents")); err != nil {
		return "", err
	} else {
		return obj, nil
	}
}

// GetRepositoryPlugins returns plugins directory for repository
func (this *Client) GetRepositoryPlugins() (string, error) {
	var obj string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &obj, OptPath("api/repository"), optDirectory("plugins")); err != nil {
		return "", err
	} else {
		return obj, nil
	}
}

// GetRepositoryCustomPlugins returns custom plugins directory for repository
func (this *Client) GetRepositoryCustomPlugins() (string, error) {
	var obj string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &obj, OptPath("api/repository"), optDirectory("customplugins")); err != nil {
		return "", err
	} else {
		return obj, nil
	}
}

// GetRepositoryScripts returns scripts directory for repository
func (this *Client) GetRepositoryScripts() (string, error) {
	var obj string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &obj, OptPath("api/repository"), optDirectory("scripts")); err != nil {
		return "", err
	} else {
		return obj, nil
	}
}

// GetRepositoryCustomScripts returns custom scripts directory for repository
func (this *Client) GetRepositoryCustomScripts() (string, error) {
	var obj string
	payload := NewGetPayload(ContentTypeJson)
	if err := this.Do(payload, &obj, OptPath("api/repository"), optDirectory("customscripts")); err != nil {
		return "", err
	} else {
		return obj, nil
	}
}
