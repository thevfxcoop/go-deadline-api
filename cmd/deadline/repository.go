package main

import (
	"net/url"

	// Modules
	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Repository struct {
	command
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewRepository(client *client.Client) Command {
	return &Repository{command{Client: client}}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Repository) Matches(args []string) (fn, url.Values) {
	params := url.Values{}
	if args[0] == "repository" && len(args) == 1 {
		return this.RunRepository, params
	}

	return nil, nil
}

func (this *Repository) RunRepository(params url.Values) error {
	var repo struct {
		Root          string
		Bin           string
		Settings      string
		Events        string
		CustomEvents  string
		Plugins       string
		CustomPlugins string
		Scripts       string
		CustomScripts string
	}

	if value, err := this.GetRepositoryRoot(); err != nil {
		return err
	} else {
		repo.Root = value
	}
	/*if value, err := this.GetRepositoryBin(); err != nil {
		return err
	} else {
		repo.Bin = value
	}
	if value, err := this.GetRepositorySettings(); err != nil {
		return err
	} else {
		repo.Settings = value
	}
	*/
	if value, err := this.GetRepositoryEvents(); err != nil {
		return err
	} else {
		repo.Events = value
	}
	if value, err := this.GetRepositoryCustomEvents(); err != nil {
		return err
	} else {
		repo.CustomEvents = value
	}
	if value, err := this.GetRepositoryPlugins(); err != nil {
		return err
	} else {
		repo.Plugins = value
	}
	if value, err := this.GetRepositoryCustomPlugins(); err != nil {
		return err
	} else {
		repo.CustomPlugins = value
	}
	if value, err := this.GetRepositoryScripts(); err != nil {
		return err
	} else {
		repo.Scripts = value
	}
	if value, err := this.GetRepositoryCustomScripts(); err != nil {
		return err
	} else {
		repo.CustomScripts = value
	}

	return this.output(repo)
}
