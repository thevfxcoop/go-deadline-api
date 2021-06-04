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
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewUsers(client *client.Client) Command {
	return &Users{command{Client: client}}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Users) Matches(args []string) (fn, url.Values) {
	params := url.Values{}
	if args[0] == "users" && len(args) == 1 {
		return this.RunUsers, params
	}
	if args[0] == "userinfo" && len(args) >= 1 {
		params["user"] = args[1:]
		return this.RunUserInfo, params
	}

	return nil, nil
}

func (this *Users) RunUsers(params url.Values) error {
	if users, err := this.GetUserNames(); err != nil {
		return err
	} else {
		return this.output(users)
	}
}

func (this *Users) RunUserInfo(params url.Values) error {
	if users, err := this.GetUserInfo(params["user"]...); err != nil {
		return err
	} else {
		return this.output(users)
	}
}
