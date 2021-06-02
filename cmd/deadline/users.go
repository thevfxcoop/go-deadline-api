package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

type Users struct {
	command
	run string
}

func NewUsers(client *client.Client, log *log.Logger) Command {
	this := new(Users)
	this.Client = client
	this.Logger = log
	return this
}

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
		} else if data, err := json.MarshalIndent(users, "", "  "); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	case "userinfo":
		if users, err := this.GetUserInfo(params["user"]...); err != nil {
			return err
		} else if data, err := json.MarshalIndent(users, "", "  "); err != nil {
			return err
		} else {
			fmt.Println(string(data))
		}
	}
	// Return success
	return nil
}
