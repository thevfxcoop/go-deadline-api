package client

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/thevfxcoop/go-deadline-api"
	"github.com/thevfxcoop/go-deadline-api/pkg/schema"
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
func OptJobState(v ...schema.JobStatus) RequestOpt {
	return func(r *http.Request) error {
		if len(v) == 0 {
			return deadline.ErrBadParameter.With("OptJobState")
		}
		values_ := make([]string, len(v))
		for i, value := range v {
			if value < schema.JobStatusUnknown || value > schema.JobStatusPending {
				return deadline.ErrBadParameter.With("OptJobState: ", v)
			}
			v_ := value.String()
			values_[i] = strings.ToUpper(v_[0:0]) + v_[1:]
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
func optTaskId(values ...uint) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		if len(values) > 0 {
			tasks := make([]string, 0, len(values))
			for _, v := range values {
				tasks = append(tasks, fmt.Sprint(v))
			}
			params.Set("TaskID", strings.Join(tasks, ","))
		}
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// optGroups for GetWorkersForGroup request sets comma-separated groups
func optGroups(value []string) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("Group", strings.Join(value, ","))
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// optPools for GetWorkersForPool request sets comma-separated groups
func optPools(value []string) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("Pool", strings.Join(value, ","))
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// optSlaves for RemoveWorkersFromGroup request sets comma-separated workers
func optSlaves(value []string) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("Slaves", strings.Join(value, ","))
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// optNameOnly for GetWorkerNames
func optNameOnly() RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("NamesOnly", "true")
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// optInfo for GetPulseNames
func optInfo() RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("Info", "true")
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// optSettings for GetPulseSettings
func optSettings() RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("Settings", "true")
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// optName for GetWorkerXX
func optName(value []string) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		if len(value) > 0 {
			params.Set("Name", strings.Join(value, ","))
		}
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// optNames for GetPulseInfo
func optNames(value []string) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		if len(value) > 0 {
			params.Set("Names", strings.Join(value, ","))
		}
		r.URL.RawQuery = params.Encode()
		return nil
	}
}

// optData for GetWorkerXX
func optData(value string) RequestOpt {
	return func(r *http.Request) error {
		params := r.URL.Query()
		params.Set("Data", value)
		r.URL.RawQuery = params.Encode()
		return nil
	}
}
