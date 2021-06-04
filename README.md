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
  * [Task]
  * [TaskReport]
  * [Worker](https://github.com/thevfxcoop/go-deadline-api/blob/main/pkg/schema/worker.go)
  * [WorkerReport](https://github.com/thevfxcoop/go-deadline-api/blob/main/pkg/schema/workerreport.go)
  * [User & UserGroup](https://github.com/thevfxcoop/go-deadline-api/blob/main/pkg/schema/user.go)

## Jobs and Job Reports

The following methods either return job information or operate on jobs:

```jobs, err := client.GetJobs(<option>,<option>,...)```

Returns all jobs, or some subset based on job status, when adding OptJobState(...)
as an option.

```job,err := client.GetJobWithId(<string>)```

Returns a job, or an ErrNotFound error if no job was found.

## Tasks and Task Reports

The following methods either return task information or operate on tasks:

TODO

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

## Appendix: Jobs JSON Format

```json
    [
        {
            "Props": {
                "Name": "sample",
                "Batch": "",
                "User": "djt",
                "Region": "none",
                "Cmmt": "",
                "Dept": "",
                "Frames": "1-400",
                "Chunk": 1,
                "Tasks": 400,
                "Grp": "none",
                "Pool": "none",
                "SecPool": "",
                "Pri": 50,
                "ReqAss": [],
                "ScrDep": [],
                "Conc": 1,
                "ConcLimt": true,
                "AuxSync": false,
                "Int": false,
                "IntPer": 100,
                "RemTmT": 0,
                "Seq": false,
                "Reload": false,
                "NoEvnt": false,
                "OnComp": 2,
                "Protect": false,
                "PathMap": [],
                "AutoTime": false,
                "TimeScrpt": false,
                "MinTime": 0,
                "MaxTime": 0,
                "Timeout": 1,
                "FrameTimeout": false,
                "StartTime": 0,
                "InitializePluginTime": 0,
                "Dep": [],
                "DepFrame": false,
                "DepFrameStart": 0,
                "DepFrameEnd": 0,
                "DepComp": true,
                "DepDel": false,
                "DepFail": false,
                "DepPer": -1.0,
                "NoBad": false,
                "OverAutoClean": false,
                "OverClean": false,
                "OverCleanDays": 0,
                "OverCleanType": 1,
                "JobFailOvr": false,
                "JobFailErr": 0,
                "TskFailOvr": false,
                "TskFailErr": 0,
                "SndWarn": true,
                "NotOvr": false,
                "SndEmail": false,
                "SndPopup": false,
                "NotEmail": [],
                "NotUser": [
                    "djt"
                ],
                "NotNote": "",
                "Limits": [],
                "ListedSlaves": [],
                "White": false,
                "MachLmt": 0,
                "MachLmtProg": -1.0,
                "PrJobScrp": "",
                "PoJobScrp": "",
                "PrTskScrp": "",
                "PoTskScrp": "",
                "Schd": 0,
                "SchdDays": 1,
                "SchdDate": "2021-05-26T18:35:51.455+00:00",
                "SchdStop": "0001-01-01T00:00:00Z",
                "MonStart": "-10675199.02:48:05.4775808",
                "MonStop": "-10675199.02:48:05.4775808",
                "TueStart": "-10675199.02:48:05.4775808",
                "TueStop": "-10675199.02:48:05.4775808",
                "WedStart": "-10675199.02:48:05.4775808",
                "WedStop": "-10675199.02:48:05.4775808",
                "ThuStart": "-10675199.02:48:05.4775808",
                "ThuStop": "-10675199.02:48:05.4775808",
                "FriStart": "-10675199.02:48:05.4775808",
                "FriStop": "-10675199.02:48:05.4775808",
                "SatStart": "-10675199.02:48:05.4775808",
                "SatStop": "-10675199.02:48:05.4775808",
                "SunStart": "-10675199.02:48:05.4775808",
                "SunStop": "-10675199.02:48:05.4775808",
                "PlugInfo": {
                    "SceneFile": "/mnt/home/djt/Documents/Blender/sample.blend",
                    "OutputFile": "/tmp/####.png",
                    "Threads": "0",
                    "Build": "None"
                },
                "Env": {},
                "EnvOnly": false,
                "PlugDir": "",
                "EventDir": "",
                "OptIns": {},
                "EventOI": [],
                "AWSPortalAssets": [],
                "AWSPortalAssetFileWhitelist": [],
                "Ex0": "",
                "Ex1": "",
                "Ex2": "",
                "Ex3": "",
                "Ex4": "",
                "Ex5": "",
                "Ex6": "",
                "Ex7": "",
                "Ex8": "",
                "Ex9": "",
                "ExDic": {},
                "OvrTaskEINames": false,
                "TaskEx0": "",
                "TaskEx1": "",
                "TaskEx2": "",
                "TaskEx3": "",
                "TaskEx4": "",
                "TaskEx5": "",
                "TaskEx6": "",
                "TaskEx7": "",
                "TaskEx8": "",
                "TaskEx9": ""
            },
            "ComFra": 0,
            "IsSub": true,
            "Purged": false,
            "Mach": "ws-dcv-a-00",
            "Date": "2021-05-26T18:35:51.53+00:00",
            "DateStart": "2021-05-26T18:35:53.2+00:00",
            "DateComp": "0001-01-01T00:00:00Z",
            "Plug": "Blender",
            "OutDir": [
                "/tmp"
            ],
            "OutFile": [
                "####.png"
            ],
            "TileFile": [],
            "Main": false,
            "MainStart": 0,
            "MainEnd": 0,
            "Tile": false,
            "TileFrame": 0,
            "TileCount": 0,
            "TileX": 0,
            "TileY": 0,
            "Stat": 4,
            "Aux": [],
            "Bad": [],
            "CompletedChunks": 0,
            "QueuedChunks": 0,
            "SuspendedChunks": 0,
            "RenderingChunks": 0,
            "FailedChunks": 400,
            "PendingChunks": 0,
            "SnglTskPrg": "0 %",
            "Errs": 51,
            "DataSize": -1,
            "_id": "60ae9508f69eb6110034df45",
            "ExtraElements": null
        }
    ]
```

## Appendix: Task JSON format

```json
{
    "JobID": "60ae9508f69eb6110034df45",
    "TaskID": 20,
    "Frames": "21-21",
    "Slave": "ws-dcv-a-00",
    "Stat": 6,
    "Prog": "0 %",
    "RndStat": "",
    "Errs": 0,
    "Start": "0001-01-01T00:00:00Z",
    "StartRen": "0001-01-01T00:00:00Z",
    "Comp": "0001-01-01T00:00:00Z",
    "WtgStrt": false,
    "NormMult": 1.0,
    "Size": 0,
    "RamPeak": 0,
    "RamPeakPer": 0,
    "Cpu": 0,
    "CpuPer": 0,
    "RamAvg": 0,
    "RamAvgPer": 0,
    "SwapAvg": 0,
    "SwapPeak": 0,
    "UsedClock": 0,
    "TotalClock": 0,
    "Props": {
        "Ex0": "",
        "Ex1": "",
        "Ex2": "",
        "Ex3": "",
        "Ex4": "",
        "Ex5": "",
        "Ex6": "",
        "Ex7": "",
        "Ex8": "",
        "Ex9": "",
        "ExDic": {}
    },
    "_id": "60ae9508f69eb6110034df45_20",
    "ExtraElements": null
}
```

## Appendix: Task Report JSON Format

```json
{
    "Job": "60ae9508f69eb6110034df45",
    "Type": 1,
    "Date": "2021-05-26T18:36:05.437+00:00",
    "Slave": "ws-dcv-a-00",
    "Task": 0,
    "Title": "Error: FailRenderException : Blender render executable was not found in the semicolon separated list \"C:\\Program Files\\Blender Foundation\\Blender\\blender.exe;C:\\Program Files (x86)\\Blender Foundation\\Blender\\blender.exe;/Applications/Blender/blender.app/Contents/MacOS/blender;/usr/local/Blender/blender\". The path to the render executable can be configured from the Plugin Configuration in the Deadline Monitor.\n   at Deadline.Plugins.DeadlinePlugin.FailRender(String message) (Python.Runtime.PythonException)\n  File \"/mnt/site/config/deadline/Deadline10/workers/ws-dcv-a-00/plugins/60ae9508f69eb6110034df45/Blender.py\", line 72, in RenderExecutable\n    self.FailRender( \"Blender render executable was not found in the semicolon separated list \\\"\" + executableList + \"\\\". The path to the render executable can be configured from the Plugin Configuration in the Deadline Monitor.\" )\n   at Python.Runtime.Dispatcher.Dispatch(ArrayList args)\n   at __FranticX_GenericDelegate0`1\\[\\[System_String\\, System_Private_CoreLib\\, Version=4_0_0_0\\, Culture=neutral\\, PublicKeyToken=7cec85d7bea7798e\\]\\]Dispatcher.Invoke()\n   at FranticX.Processes.ManagedProcess.RenderExecutable()\n   at Deadline.Plugins.DeadlinePlugin.RenderExecutable()\n   at FranticX.Processes.ManagedProcess.Execute(Boolean waitForExit)\n   at Deadline.Plugins.PluginWrapper.RenderTasks(Task task, String& outMessage, AbortLevel& abortLevel)",
    "JobName": "sample",
    "JobUser": "djt",
    "Plugin": "Blender",
    "Frames": "1-1",
    "TaskStartTime": "0001-01-01T00:00:00Z",
    "TaskTime": 10,
    "LogErr": "",
    "AverageCpu": 13,
    "PeakCpu": 20,
    "AverageRam": 8305750528,
    "RamAvgPer": 51,
    "PeakRam": 8309153792,
    "RamPeakPer": 51,
    "UsedClock": 4018,
    "TotalClock": 30902,
    "_id": "60ae9515d1300e1054f5cb7d"
}
```

