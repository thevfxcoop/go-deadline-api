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
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewPools(client *client.Client) Command {
	return &Pools{command{Client: client}}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Pools) Matches(args []string) (fn, url.Values) {
	params := url.Values{}
	if args[0] == "pools" && len(args) == 1 {
		return this.RunGetPools, params
	}
	if args[0] == "poolworkers" && len(args) > 1 {
		params["pools"] = args[1:]
		return this.RunGetPoolWorkers, params
	}
	if args[0] == "addpool" && len(args) > 1 {
		params["pools"] = args[1:]
		return this.RunAddPool, params
	}
	if args[0] == "deletepool" && len(args) > 1 {
		params["pools"] = args[1:]
		return this.RunDeletePool, params
	}
	return nil, nil
}

func (this *Pools) RunGetPools(params url.Values) error {
	if pools, err := this.GetPools(); err != nil {
		return err
	} else {
		return this.output(pools)
	}
}
func (this *Pools) RunGetPoolWorkers(params url.Values) error {
	if workers, err := this.GetWorkersForPool(params["pools"]...); err != nil {
		return err
	} else {
		return this.output(workers)
	}
}

func (this *Pools) RunAddPool(params url.Values) error {
	if err := this.AddPools(params["pools"]...); err != nil {
		return err
	} else if pools, err := this.GetPools(); err != nil {
		return err
	} else {
		return this.output(pools)
	}
}

func (this *Pools) RunDeletePool(params url.Values) error {
	if err := this.DeletePools(params["pools"]...); err != nil {
		return err
	} else if pools, err := this.GetPools(); err != nil {
		return err
	} else {
		return this.output(pools)
	}
}
