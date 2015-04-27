package model

import (
	"fmt"
)

type DesiredAgent struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Tarball string `json:"tarball"`
	Md5     string `json:"md5"`
	Cmd     string `json:"cmd"`
}

func (this *DesiredAgent) String() string {
	return fmt.Sprintf(
		"<Name:%s, Version:%s, Tarball:%s, Md5:%s, Cmd:%s>",
		this.Name,
		this.Version,
		this.Tarball,
		this.Md5,
		this.Cmd,
	)
}

type RealAgent struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

func (this *RealAgent) String() string {
	return fmt.Sprintf(
		"<Name:%s, Version:%s, Status:%s, Timestamp:%v>",
		this.Name,
		this.Version,
		this.Status,
		this.Timestamp,
	)
}

type HeartbeatRequest struct {
	Hostname   string       `json:"hostname"`
	RealAgents []*RealAgent `json:"realAgents"`
}

type HeartbeatResponse struct {
	ErrorMessage  string          `json:"errorMessage"`
	DesiredAgents []*DesiredAgent `json:"desiredAgents"`
}
