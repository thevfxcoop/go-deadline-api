# go-deadline-api

Deadline API. Please see outline of the REST API [here](https://docs.thinkboxsoftware.com/products/deadline/10.1/1_User%20Manual/manual/index-rest-api.html). There is a lack of informaton on the schema, but it may just come straight out of MongoDB, which is still to be determined.

The pre-requisites for this api are:

 * A working Deadline Web Service
 * Any recent version of go

 You can view the API reference for this module here: https://pkg.go.dev/github.com/thevfxcoop/go-deadline-api
 
## Getting Started

To use this API in your own code, import it and create a client:

```go
package main

import (
	"github.com/thevfxcoop/go-deadline-api/pkg/client"
)

func main() {
    var endpoint *url.URL
    client, err := client.NewClient(endpoint)
    if err != nil {
        // ...
    }
    version, err := client.Ping()
    if err != nil {
        // ...
    }    
}
```

You can call the `Ping` method in order to get the Deadline version,
or it will return an error if the repository cannot be accessed.

## Schema

The native deadline schema is somewhat undocumented, so the schema for the
webservice is translated into a documented schema and available
in `pkg/schema` for jobs, tasks, and so forth. The schema documentation is:

  * [Job](https://github.com/thevfxcoop/go-deadline-api/blob/main/pkg/schema/job.go)
  * [JobReport](https://github.com/thevfxcoop/go-deadline-api/blob/main/pkg/schema/jobreport.go)
  * [Task](https://github.com/thevfxcoop/go-deadline-api/blob/main/pkg/schema/task.go)
  * [TaskReport](https://github.com/thevfxcoop/go-deadline-api/blob/main/pkg/schema/taskreport.go)
  * [Worker](https://github.com/thevfxcoop/go-deadline-api/blob/main/pkg/schema/worker.go)
  * [WorkerReport](https://github.com/thevfxcoop/go-deadline-api/blob/main/pkg/schema/workerreport.go)
  * [User & UserGroup](https://github.com/thevfxcoop/go-deadline-api/blob/main/pkg/schema/user.go)
  * [Pulse](https://github.com/thevfxcoop/go-deadline-api/blob/main/pkg/schema/pulse.go)

## Command-Line Tool

There is a command-line tool to demonstrate some of the features of the API. In order to
build this, use the following command:

```bash
[bash] git clone https://github.com/thevfxcoop/go-deadline-api
[bash] cd go-deadline-api && make
```

The tool is placed in the "build" folder within the repository. To use the tool, you can
set the environment variable DEADLINE_ENDPOINT to the endpoint for your web service or
include the endpoint URL on the command line with the `-endpoint` flag. For example,


```bash
[bash] export DEADLINE_ENDPOINT=DEADLINE_ENDPOINT=http://localhost:8082/
[bash] build/deadline 
Deadline Web Service 10.1 [v10.1.15.2 Release (313fe6482)]
```

Without any arguments the version of deadline is returned, or an error if
the deadline repository cannot be found. You can add a `-debug` flag to trace
traffic to and from your web service:

```bash
[bash] export DEADLINE_ENDPOINT=DEADLINE_ENDPOINT=http://localhost:8082/
[bash] build/deadline -debug jobs
payload: {}
req: GET http://localhost:8082/
  => 200 OK
     Deadline Web Service 10.1 [v10.1.15.2 Release (313fe6482)]
  Took 84 ms
payload: {}
req: GET http://localhost:8082/api/jobs
  => 200 OK
     [{"Props":{"Name":"PDG_generate","Batch":"PDG deadline_testing 2021-06-03 14:48:34.032908",
     ...]
```

Responses from the tool are currently displayed in JSON format. The commands you can use
are listed if you use the `-help` flag:

```bash
[bash] build/deadline -help
Usage of deadline:
  deadline <flags> <command> (<args>)

Flags:
  -debug
    	Trace request and reponse with API
  -endpoint string
    	Endpoint URL, can be overridden with DEADLINE_ENDPOINT environment variable

Build:
  Build: v0.0.1-4-g483f386 (branch: "main" hash:"483f3869a2a15970a88bf013adf1e6b52899903d")
   Time: 2021-06-04T10:43:57Z
     Go: go1.16.3 (darwin/amd64)
```

## License

Copyright 2021 The VFX Cooperative

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

>http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

## Appendix: Port forwarding

If you want to access a Deadline repository which is in a secure environment
accessible using SSH, use the following command to forward port 8082 to your
localhost, for example:

```bash
[bash] ssh -L 8082:localhost:8082 <remote server>
```
