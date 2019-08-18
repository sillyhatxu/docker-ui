package dto

type DockerStatsDTO struct {
	ContainerId   string `json:"containerId"`
	Name          string `json:"name"`
	Cpu           string `json:"cpu"`
	MemUsageLimit string `json:"memUsageLimit"`
	Mem           string `json:"mem"`
	NetIO         string `json:"netIO"`
	BlockIO       string `json:"blockIO"`
	Pids          int    `json:"pids"`
}

type ServiceDTO struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Mode     string `json:"mode"`
	Replicas string `json:"replicas"`
	Image    string `json:"image"`
	Ports    string `json:"ports"`
}
