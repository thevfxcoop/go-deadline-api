package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

type Groups struct {
	command
	run string
}

func NewGroups(client *client.Client, log *log.Logger) Command {
	this := new(Groups)
	this.Client = client
	this.Logger = log
	return this
}

func (this *Groups) Matches(args []string) url.Values {
	params := url.Values{}
	this.run = args[0]
	if args[0] == "groups" && len(args) == 1 {
		return params
	}
	if args[0] == "groupworkers" && len(args) > 1 {
		params["groups"] = args[1:]
		return params
	}
	if args[0] == "addgroup" && len(args) > 1 {
		params["groups"] = args[1:]
		return params
	}
	if args[0] == "deletegroup" && len(args) > 1 {
		params["groups"] = args[1:]
		return params
	}
	if args[0] == "setgroups" && len(args) > 1 {
		params["groups"] = args[1:]
		return params
	}
	if args[0] == "addgroupworkers" && len(args) > 2 {
		params.Set("group", args[1])
		params["workers"] = args[2:]
		return params
	}
	if args[0] == "removegroupworkers" && len(args) > 2 {
		params.Set("group", args[1])
		params["workers"] = args[2:]
		return params
	}
	return nil
}

func (this *Groups) Run(params url.Values) error {
	switch this.run {
	case "groups":
		if groups, err := this.GetGroups(); err != nil {
			return err
		} else if data, err := json.Marshal(groups); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	case "groupworkers":
		if workers, err := this.GetWorkersForGroup(params["groups"]...); err != nil {
			return err
		} else if data, err := json.Marshal(workers); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	case "addgroup":
		if err := this.AddGroups(params["groups"]...); err != nil {
			return err
		} else if groups, err := this.GetGroups(); err != nil {
			return err
		} else if data, err := json.Marshal(groups); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	case "deletegroup":
		if err := this.DeleteGroups(params["groups"]...); err != nil {
			return err
		} else if groups, err := this.GetGroups(); err != nil {
			return err
		} else if data, err := json.Marshal(groups); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	case "setgroups":
		if err := this.SetGroups(params["groups"]...); err != nil {
			return err
		} else if groups, err := this.GetGroups(); err != nil {
			return err
		} else if data, err := json.Marshal(groups); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	case "addgroupworkers":
		if err := this.AddWorkersToGroup(params.Get("group"), params["workers"]...); err != nil {
			return err
		} else if workers, err := this.GetWorkersForGroup(params.Get("group")); err != nil {
			return err
		} else if data, err := json.Marshal(workers); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	case "removegroupworkers":
		if err := this.RemoveWorkersFromGroup(params.Get("group"), params["workers"]...); err != nil {
			return err
		} else if workers, err := this.GetWorkersForGroup(params.Get("group")); err != nil {
			return err
		} else if data, err := json.Marshal(workers); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	}

	// Return success
	return nil
}
