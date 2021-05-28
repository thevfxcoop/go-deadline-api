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

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewGetPayload(accept string) Payload {
	this := new(payload)
	this.method = http.MethodGet
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

///////////////////////////////////////////////////////////////////////////////
// PAYLOAD METHODS

func (this *payload) Method() string {
	return this.method
}

func (this *payload) Accept() string {
	return this.accept
}
