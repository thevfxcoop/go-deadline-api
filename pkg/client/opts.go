package client

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/thevfxcoop/go-deadline-api"
)

// OptTimeout sets the timeout on any request. By default, a timeout
// of 10 seconds is used if OptTimeout is not set
func OptTimeout(value time.Duration) ClientOpt {
	return func(client *Client) error {
		client.Client.Timeout = value
		return nil
	}
}

// OptUserAgent sets the user agent string on each API request
// It is set to the default if empty string is passed
func OptUserAgent(value string) ClientOpt {
	return func(client *Client) error {
		value = strings.TrimSpace(value)
		if value == "" {
			client.ua = DefaultUserAgent
		} else {
			client.ua = value
		}
		return nil
	}
}

// OptTrace allows you to be the "man in the middle" on any
// requests so you can see traffic move back and forth.
// Setting verbose to true also displays the JSON response
func OptTrace(w io.Writer, verbose bool) ClientOpt {
	return func(client *Client) error {
		client.Client.Transport = NewLogTransport(w, client.Client.Transport, verbose)
		return nil
	}
}

// OptStrict turns on strict content type checking on anything returned
// from the API
func OptStrict() ClientOpt {
	return func(client *Client) error {
		client.strict = true
		return nil
	}
}

// OptPath appends path elements onto a request
func OptPath(value string) RequestOpt {
	return func(r *http.Request) error {
		// Make a copy
		url := *r.URL
		// Clean up and append path
		url.Path = PathSeparator + filepath.Join(strings.Trim(url.Path, PathSeparator), value)
		// Set new path
		r.URL = &url
		return nil
	}
}

// OptJobState(...JobState) for GetJobs request
func OptJobState(values ...JobState) RequestOpt {
	return func(r *http.Request) error {
		if len(values) == 0 {
			return deadline.ErrBadParameter.With("OptJobState")
		}
		values_ := make([]string, len(values))
		for i, value := range values {
			values_[i] = string(value)
		}
		params := r.URL.Query()
		params.Set("States", strings.Join(values_, ","))
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// optIdOnly(bool) for GetJobIds and GetTaskIds request returns only the Id's
func optIdOnly(value bool) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("IdOnly", fmt.Sprint(value))
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// optJobId for GetJobWithId request returns a single job
func optJobId(value string) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("JobID", value)
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// optTaskId for GetTaskWithId request returns a single task
func optTaskId(value uint) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("TaskID", fmt.Sprint(value))
		r.URL.RawQuery = params.Encode()
		return nil
	}
}
