package haproxy

import ()

type BackendHostStatus string

const (
	StatusDown    BackendHostStatus = "DOWN"
	StatusNOLB    BackendHostStatus = "NOLB"
	StatusUp      BackendHostStatus = "UP"
	StatusNoCheck BackendHostStatus = "no check"
	StatusUnknown BackendHostStatus = "unkown"
)

func ToBackendHostStatus(status string) BackendHostStatus {
	switch status {
	case "DOWN":
		return StatusDown
	case "NOLB":
		return StatusNOLB
	case "UP":
		return StatusUp
	case "no check":
		return StatusNoCheck
	default:
		return StatusUnknown
	}
}

type BackendHost struct {
	Hostname string
	Status   BackendHostStatus
}

func NewBackendHost(hostname string, status BackendHostStatus) *BackendHost {
	return &BackendHost{Hostname: hostname, Status: status}
}
