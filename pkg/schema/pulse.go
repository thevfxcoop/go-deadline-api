package schema

import (
	"encoding/json"
	"fmt"
	"time"

	// Modules
	"github.com/thevfxcoop/go-deadline-api"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Pulse struct {
	Id              string        `deadline:"_id" json:"_id,omitempty"`
	Name            string        `deadline:"Name" json:"name,omitempty"`
	Status          PulseStatus   `deadline:"Stat" json:"status,omitempty"`
	Date            time.Time     `deadline:"StatTime" json:"date,omitempty"`
	Uptime          time.Duration `deadline:"UpTime" json:"uptime,omitempty"`
	Version         string        `deadline:"Ver" json:"version,omitempty"`
	PulsePort       uint          `deadline:"Port" json:"pulse_port,omitempty"`
	ServerPort      uint          `deadline:"ServerPort" json:"server_port,omitempty"`
	PulseHost       string        `deadline:"Host" json:"pulse_host,omitempty"`
	MacAddress      string        `deadline:"MAC" json:"mac_address,omitempty"`
	IpAddress       string        `deadline:"IP" json:"ip_address,omitempty"`
	MemTotalBytes   uint          `deadline:"RAM" json:"mem_total_bytes,omitempty"`
	MemFreeBytes    uint          `deadline:"RAMFree" json:"mem_free_bytes,omitempty"`
	CpuOS           string        `deadline:"OS" json:"cpu_os,omitempty"`
	CpuArchitecture string        `deadline:"Arch" json:"cpu_architecture,omitempty"`
	CpuCores        uint          `deadline:"Procs" json:"cpu_cores,omitempty"`
	CpuPercent      uint          `deadline:"CPU" json:"cpu_percent,omitempty"`
	GpuName         string        `deadline:"Vid" json:"gpu_name,omitempty"`
}

type PulseSettings struct {
	Id          string `deadline:"_id" json:"_id,omitempty"`
	Name        string `deadline:"Name" json:"name,omitempty"`
	Primary     bool   `deadline:"Primary" json:"primary,omitempty"`
	UsePort     bool   `deadline:"UsePort" json:"use_port,omitempty"`
	UseListPort bool   `deadline:"UseListPort" json:"use_list_port,omitempty"`
	PulsePort   uint   `deadline:"Port" json:"pulse_port,omitempty"`
	ListPort    uint   `deadline:"ListPort" json:"list_port,omitempty"`
	PulseHost   string `deadline:"Host" json:"pulse_host,omitempty"`
	MacAddress  string `deadline:"Mac" json:"mac_address,omitempty"`
}

type PulseStatus uint

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	// https://docs.thinkboxsoftware.com/products/deadline/10.1/1_User%20Manual/manual/rest-pulse.html#pulse-property-values
	PulseStatusUnknown PulseStatus = 0
	PulseStatusRunning PulseStatus = 1
	PulseStatusOffline PulseStatus = 2
	PulseStatusStalled PulseStatus = 4
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewPulse(v map[string]interface{}) (*Pulse, error) {
	this := new(Pulse)

	if err := decoder.Decode(v, this); err != nil {
		return nil, err
	} else {
		return this, nil
	}
}

func NewPulseSettings(v map[string]interface{}) (*PulseSettings, error) {
	this := new(PulseSettings)

	if err := decoder.Decode(v, this); err != nil {
		return nil, err
	} else {
		return this, nil
	}
}

///////////////////////////////////////////////////////////////////////////////
// MARSHAL AND UNMARSHALL

func (s *PulseStatus) Unmarshal(in interface{}) error {
	switch v := in.(type) {
	case float64:
		*s = PulseStatus(v)
	default:
		return deadline.ErrBadParameter.With("Decode: ", in)
	}
	// return success
	return nil
}

func (s PulseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprint(s))
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (v *Pulse) String() string {
	if data, err := json.MarshalIndent(v, "", "  "); err == nil {
		return string(data)
	} else {
		return err.Error()
	}
}

func (v PulseStatus) String() string {
	switch v {
	case PulseStatusUnknown:
		return "unknown"
	case PulseStatusRunning:
		return "running"
	case PulseStatusOffline:
		return "offline"
	case PulseStatusStalled:
		return "stalled"
	default:
		return "[?? Invalid PulseStatus value]"
	}
}
