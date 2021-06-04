package main

import (
	"net/url"

	// Modules
	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Pools struct {
	command
	run string
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewPools(client *client.Client) Command {
	this := new(Pools)
	this.Client = client
	return this
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Pools) Matches(args []string) url.Values {
	params := url.Values{}
	this.run = args[0]
	if args[0] == "pools" && len(args) == 1 {
		return params
	}
	if args[0] == "poolworkers" && len(args) > 1 {
		params["pools"] = args[1:]
		return params
	}
	if args[0] == "addpool" && len(args) > 1 {
		params["pools"] = args[1:]
		return params
	}
	if args[0] == "deletepool" && len(args) > 1 {
		params["pools"] = args[1:]
		return params
	}
	return nil
}

func (this *Pools) Run(params url.Values) error {
	switch this.run {
	case "pools":
		if pools, err := this.GetPools(); err != nil {
			return err
		} else {
			return this.output(pools)
		}
	case "poolworkers":
		if workers, err := this.GetWorkersForPool(params["pools"]...); err != nil {
			return err
		} else {
			return this.output(workers)
		}
	case "addpool":
		if err := this.AddPools(params["pools"]...); err != nil {
			return err
		} else if pools, err := this.GetPools(); err != nil {
			return err
		} else {
			return this.output(pools)
		}
	case "deletepool":
		if err := this.DeletePools(params["pools"]...); err != nil {
			return err
		} else if pools, err := this.GetPools(); err != nil {
			return err
		} else {
			return this.output(pools)
		}
	}

	// Return success
	return nil
}
