package model

import (
	"fmt"
	"path"
)

type DesiredAgent struct {
	Name            string `json:"name"`
	Version         string `json:"version"`
	Tarball         string `json:"tarball"`
	Md5             string `json:"md5"`
	Cmd             string `json:"cmd"`
	AgentDir        string `json:"-"`
	AgentVersionDir string `json:"-"`
	TarballFilename string `json:"-"`
	Md5Filename     string `json:"-"`
	TarballFilepath string `json:"-"`
	Md5Filepath     string `json:"-"`
	ControlFilepath string `json:"-"`
	TarballUrl      string `json:"-"`
	Md5Url          string `json:"-"`
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

func (this *DesiredAgent) FillAttrs(workdir string) {
	this.AgentDir = path.Join(workdir, this.Name)
	this.AgentVersionDir = path.Join(this.AgentDir, this.Version)
	this.TarballFilename = fmt.Sprintf("%s-%s.tar.gz", this.Name, this.Version)
	this.Md5Filename = fmt.Sprintf("%s.md5", this.TarballFilename)
	this.TarballFilepath = path.Join(this.AgentVersionDir, this.TarballFilename)
	this.Md5Filepath = path.Join(this.AgentVersionDir, this.Md5Filename)
	this.ControlFilepath = path.Join(this.AgentVersionDir, "control")

	if this.Md5 == "" {
		this.Md5 = this.Tarball
	}

	this.TarballUrl = fmt.Sprintf("%s/%s", this.Tarball, this.TarballFilename)
	this.Md5Url = fmt.Sprintf("%s/%s", this.Md5, this.Md5Filename)
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
