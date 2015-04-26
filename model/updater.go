package model

type DesiredAgent struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Tarball string `json:"tarball"`
	Md5     string `json:"md5"`
	Cmd     string `json:"cmd"`
}

type RealAgent struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

type HeartbeatRequest struct {
	Hostname   string       `json:"hostname"`
	RealAgents []*RealAgent `json:"realAgents"`
}

type HeartbeatResponse struct {
	ErrorMessage  string          `json:"errorMessage"`
	DesiredAgents []*DesiredAgent `json:"desiredAgents"`
}
