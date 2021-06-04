package main

import (
	"net/url"

	// Modules
	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Users struct {
	command
	run string
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewUsers(client *client.Client) Command {
	this := new(Users)
	this.Client = client
	return this
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Users) Matches(args []string) url.Values {
	params := url.Values{}
	this.run = args[0]
	if args[0] == "users" && len(args) == 1 {
		return params
	}
	if args[0] == "userinfo" && len(args) >= 1 {
		params["user"] = args[1:]
		return params
	}

	// Command not found
	return nil
}

func (this *Users) Run(params url.Values) error {
	switch this.run {
	case "users":
		if users, err := this.GetUserNames(); err != nil {
			return err
		} else {
			return this.output(users)
		}
	case "userinfo":
		if users, err := this.GetUserInfo(params["user"]...); err != nil {
			return err
		} else {
			return this.output(users)
		}
	}
	// Return success
	return nil
}
