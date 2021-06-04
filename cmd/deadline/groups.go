package main

import (
	"net/url"

	// Modules
	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Groups struct {
	command
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewGroups(client *client.Client) Command {
	return &Groups{command{Client: client}}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Groups) Matches(args []string) (fn, url.Values) {
	params := url.Values{}
	if args[0] == "groups" && len(args) == 1 {
		return this.RunGroups, params
	}
	if args[0] == "groupworkers" && len(args) > 1 {
		params["groups"] = args[1:]
		return this.RunWorkersForGroup, params
	}
	if args[0] == "addgroup" && len(args) > 1 {
		params["groups"] = args[1:]
		return this.RunAddGroups, params
	}
	if args[0] == "deletegroup" && len(args) > 1 {
		params["groups"] = args[1:]
		return this.RunDeleteGroups, params
	}
	if args[0] == "setgroups" && len(args) > 1 {
		params["groups"] = args[1:]
		return this.RunSetGroups, params
	}
	if args[0] == "addgroupworkers" && len(args) > 2 {
		params.Set("group", args[1])
		params["workers"] = args[2:]
		return this.RunAddGroupWorkers, params
	}
	if args[0] == "removegroupworkers" && len(args) > 2 {
		params.Set("group", args[1])
		params["workers"] = args[2:]
		return this.RunRemoveGroupWorkers, params
	}
	return nil, nil
}

func (this *Groups) RunGroups(params url.Values) error {
	if groups, err := this.GetGroups(); err != nil {
		return err
	} else {
		return this.output(groups)
	}
}

func (this *Groups) RunWorkersForGroup(params url.Values) error {
	if workers, err := this.GetWorkersForGroup(params["groups"]...); err != nil {
		return err
	} else {
		return this.output(workers)
	}
}

func (this *Groups) RunAddGroups(params url.Values) error {
	if err := this.AddGroups(params["groups"]...); err != nil {
		return err
	} else if groups, err := this.GetGroups(); err != nil {
		return err
	} else {
		return this.output(groups)
	}
}

func (this *Groups) RunDeleteGroups(params url.Values) error {
	if err := this.DeleteGroups(params["groups"]...); err != nil {
		return err
	} else if groups, err := this.GetGroups(); err != nil {
		return err
	} else {
		return this.output(groups)
	}
}

func (this *Groups) RunSetGroups(params url.Values) error {
	if err := this.SetGroups(params["groups"]...); err != nil {
		return err
	} else if groups, err := this.GetGroups(); err != nil {
		return err
	} else {
		return this.output(groups)
	}
}
func (this *Groups) RunAddGroupWorkers(params url.Values) error {
	if err := this.AddWorkersToGroup(params.Get("group"), params["workers"]...); err != nil {
		return err
	} else if workers, err := this.GetWorkersForGroup(params.Get("group")); err != nil {
		return err
	} else {
		return this.output(workers)
	}
}

func (this *Groups) RunRemoveGroupWorkers(params url.Values) error {
	if err := this.RemoveWorkersFromGroup(params.Get("group"), params["workers"]...); err != nil {
		return err
	} else if workers, err := this.GetWorkersForGroup(params.Get("group")); err != nil {
		return err
	} else {
		return this.output(workers)
	}
}
