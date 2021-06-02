package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

type Pools struct {
	command
	run string
}

func NewPools(client *client.Client, log *log.Logger) Command {
	this := new(Pools)
	this.Client = client
	this.Logger = log
	return this
}

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
		} else if data, err := json.Marshal(pools); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	case "poolworkers":
		if workers, err := this.GetWorkersForPool(params["pools"]...); err != nil {
			return err
		} else if data, err := json.Marshal(workers); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	case "addpool":
		if err := this.AddPools(params["pools"]...); err != nil {
			return err
		} else if pools, err := this.GetPools(); err != nil {
			return err
		} else if data, err := json.Marshal(pools); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	case "deletepool":
		if err := this.DeletePools(params["pools"]...); err != nil {
			return err
		} else if pools, err := this.GetPools(); err != nil {
			return err
		} else if data, err := json.Marshal(pools); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	}

	// Return success
	return nil
}
