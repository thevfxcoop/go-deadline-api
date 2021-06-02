package client

import "net/http"

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Payload interface {
	Method() string
	Accept() string
}

type payload struct {
	method string
	accept string
}

type commandpayload struct {
	payload
	Command  string
	JobID    string
	TaskList []uint `json:"TaskList,omitempty"`
}

type grouppayload struct {
	payload
	Group     []string `json:"Group,omitempty"`
	Slaves    []string `json:"Slave,omitempty"`
	Overwrite bool     `json:"OverWrite"`
}

type poolpayload struct {
	payload
	Pool      []string `json:"Pool,omitempty"`
	Slaves    []string `json:"Slave,omitempty"`
	Overwrite bool     `json:"OverWrite"`
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewGetPayload(accept string) Payload {
	this := new(payload)
	this.method = http.MethodGet
	this.accept = accept
	return this
}

func NewDeletePayload(accept string) Payload {
	this := new(payload)
	this.method = http.MethodDelete
	this.accept = accept
	return this
}

func NewJobCommandPayload(id, command string) Payload {
	this := new(commandpayload)
	this.method = http.MethodPut
	this.accept = ContentTypeText
	this.Command = command
	this.JobID = id
	return this
}

func NewTasksCommandPayload(id, command string, tasks []uint) Payload {
	this := new(commandpayload)
	this.method = http.MethodPut
	this.accept = ContentTypeText
	this.Command = command
	this.JobID = id
	if len(tasks) > 0 {
		this.TaskList = tasks
	}
	return this
}

func NewGroupPayload(groups []string, overwrite bool, workers ...string) Payload {
	this := new(grouppayload)
	this.method = http.MethodPost
	this.accept = ContentTypeText
	if len(groups) > 0 {
		this.Group = groups
	}
	if len(workers) > 0 {
		this.Slaves = workers
		this.method = http.MethodPut
	}
	this.Overwrite = overwrite
	return this
}

func NewPoolPayload(pools []string, overwrite bool, workers ...string) Payload {
	this := new(poolpayload)
	this.method = http.MethodPost
	this.accept = ContentTypeText
	if len(pools) > 0 {
		this.Pool = pools
	}
	if len(workers) > 0 {
		this.Slaves = workers
		this.method = http.MethodPut
	}
	this.Overwrite = overwrite
	return this
}

///////////////////////////////////////////////////////////////////////////////
// PAYLOAD METHODS

func (this *payload) Method() string {
	return this.method
}

func (this *payload) Accept() string {
	return this.accept
}
