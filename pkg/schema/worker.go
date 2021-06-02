package schema

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/thevfxcoop/go-deadline-api"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type WorkerInfo struct {
	Id                  string        `deadline:"Info/_id" json:"_id"`
	Name                string        `deadline:"Info/Name" json:"name,omitempty"`
	User                string        `deadline:"Info/User" json:"user,omitempty"`
	Status              WorkerStatus  `deadline:"Info/Stat" json:"status"`
	Job                 string        `deadline:"Info/JobId" json:"job,omitempty"`
	JobName             string        `deadline:"Info/JobName" json:"job_name,omitempty"`
	JobGroup            string        `deadline:"Info/JobGrp" json:"job_group,omitempty"`
	JobUser             string        `deadline:"Info/JobUser" json:"job_user,omitempty"`
	JobPool             string        `deadline:"Info/JobPool" json:"job_pool,omitempty"`
	Region              string        `deadline:"Info/Reg" json:"region,omitempty"`
	Hostname            string        `deadline:"Info/Host" json:"hostname,omitempty"`
	Uptime              time.Duration `deadline:"Info/UpTime" json:"uptime,omitempty"`
	MacAddress          string        `deadline:"Info/MAC" json:"mac_address,omitempty"`
	IpAddress           string        `deadline:"Info/IP" json:"ip_address,omitempty"`
	MemTotalBytes       uint          `deadline:"Info/RAM" json:"mem_total_bytes,omitempty"`
	MemFreeBytes        uint          `deadline:"Info/RAMFree" json:"mem_free_bytes,omitempty"`
	CpuOS               string        `deadline:"Info/OS" json:"cpu_os,omitempty"`
	CpuArchitecture     string        `deadline:"Info/Arch" json:"cpu_architecture,omitempty"`
	CpuCores            uint          `deadline:"Info/Procs" json:"cpu_cores,omitempty"`
	GpuName             string        `deadline:"Info/Vid" json:"gpu_name,omitempty"`
	Groups              string        `deadline:"Info/Grps" json:"groups,omitempty"`
	CompletedTasks      uint          `deadline:"Info/CompletedTasks" json:"completed_tasks,omitempty"`
	AWSInstanceId       string        `deadline:"Info/AWSInfo/AWSInstanceId" json:"aws_instance_id,omitempty"`
	AWSInstanceType     string        `deadline:"Info/AWSInfo/AWSInstanceType" json:"aws_instance_type,omitempty"`
	AWSImageId          string        `deadline:"Info/AWSInfo/AWSImageId" json:"aws_image_id,omitempty"`
	AWSArchitecture     string        `deadline:"Info/AWSInfo/AWSArchitecture" json:"aws_instance_arch,omitempty"`
	AWSRegion           string        `deadline:"Info/AWSInfo/AWSRegion" json:"aws_region,omitempty"`
	AWSAvailabilityZone string        `deadline:"Info/AWSInfo/AWSAvailabilityZone" json:"aws_availability_zone,omitempty"`
}

type WorkerStatus uint

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	// https://docs.thinkboxsoftware.com/products/deadline/10.1/1_User%20Manual/manual/rest-workers.html#worker-property-values
	WorkerStatusUnknown     WorkerStatus = 0
	WorkerStatusRendering   WorkerStatus = 1
	WorkerStatusIdle        WorkerStatus = 2
	WorkerStatusOffline     WorkerStatus = 3
	WorkerStatusStalled     WorkerStatus = 4
	WorkerStatusStartingJob WorkerStatus = 8
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewWorkerInfo(v map[string]interface{}) (*WorkerInfo, error) {
	this := new(WorkerInfo)

	if err := decoder.Decode(v, this); err != nil {
		return nil, err
	} else {
		return this, nil
	}
}

///////////////////////////////////////////////////////////////////////////////
// MARSHAL AND UNMARSHALL

func (s *WorkerStatus) Unmarshal(in interface{}) error {
	switch v := in.(type) {
	case float64:
		*s = WorkerStatus(v)
	default:
		return deadline.ErrBadParameter.With("Decode: ", in)
	}
	// return success
	return nil
}

func (s WorkerStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprint(s))
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (v *WorkerInfo) String() string {
	if data, err := json.MarshalIndent(v, "", "  "); err == nil {
		return string(data)
	} else {
		return err.Error()
	}
}

func (v WorkerStatus) String() string {
	switch v {
	case WorkerStatusUnknown:
		return "unknown"
	case WorkerStatusRendering:
		return "rendering"
	case WorkerStatusIdle:
		return "idle"
	case WorkerStatusOffline:
		return "offline"
	case WorkerStatusStalled:
		return "stalled"
	case WorkerStatusStartingJob:
		return "startingjob"
	default:
		return "[?? Invalid WorkerStatus value]"
	}
}
