package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type logtransport struct {
	http.RoundTripper
	w io.Writer
	v bool
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewLogTransport creates middleware into the request/response so you can log
// the transmission on the wire. Setting verbose to true also displays the
// body of each response
func NewLogTransport(w io.Writer, parent http.RoundTripper, verbose bool) http.RoundTripper {
	this := new(logtransport)
	this.w = w
	this.v = verbose
	this.RoundTripper = parent
	return this
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS http.RoundTripper

func (this *logtransport) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Fprintln(this.w, "req:", req.Method, redactedUrl(req.URL))
	then := time.Now()
	defer func() {
		fmt.Fprintln(this.w, "  Took", time.Since(then).Milliseconds(), "ms")
	}()
	resp, err := this.RoundTripper.RoundTrip(req)
	if err != nil {
		fmt.Fprintln(this.w, "  => Error:", err)
	} else {
		fmt.Fprintln(this.w, "  =>", resp.Status)
	}
	// If verbose is switched on, read the body
	if this.v {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Fprintln(this.w, "    ", string(body))
		}
		resp.Body = ioutil.NopCloser(bytes.NewReader(body))
	}
	return resp, err
}
