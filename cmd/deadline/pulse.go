package main

import (
	"net/url"

	// Modules
	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Pulse struct {
	command
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewPulse(client *client.Client) Command {
	return &Pulse{command{Client: client}}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Pulse) Matches(args []string) (fn, url.Values) {
	params := url.Values{}
	if args[0] == "pulse" && len(args) == 1 {
		return this.RunPulseNames, params
	}
	if args[0] == "pulseinfo" && len(args) >= 1 {
		params["pulse"] = args[1:]
		return this.RunPulseInfo, params
	}
	if args[0] == "pulsesettings" && len(args) >= 1 {
		params["pulse"] = args[1:]
		return this.RunPulseSettings, params
	}

	return nil, nil
}

func (this *Pulse) RunPulseNames(params url.Values) error {
	if users, err := this.GetPulseNames(); err != nil {
		return err
	} else {
		return this.output(users)
	}
}

func (this *Pulse) RunPulseInfo(params url.Values) error {
	if users, err := this.GetPulseInfo(params["pulse"]...); err != nil {
		return err
	} else {
		return this.output(users)
	}
}

func (this *Pulse) RunPulseSettings(params url.Values) error {
	if users, err := this.GetPulseSettings(params["pulse"]...); err != nil {
		return err
	} else {
		return this.output(users)
	}
}
