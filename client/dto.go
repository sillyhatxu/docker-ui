package client

import "fmt"

type ContainerStats struct {
	OSType string     `json:"ostype"`
	Body   *StatsBody `json:"body"`
}

func (cs ContainerStats) GetId() string {
	if cs.Body != nil && len(cs.Body.Id) > 12 {
		return cs.Body.Id[0:12]
	}
	return ""
}

func (cs ContainerStats) GetName() string {
	if cs.Body != nil && len(cs.Body.Name) > 0 {
		return cs.Body.Name[1:len(cs.Body.Name)]
	}
	return ""
}

func (cs ContainerStats) GetCPUPercent() string {
	if cs.Body != nil && cs.Body.CPUStats != nil && cs.Body.CPUStats.CPUUsage != nil {
		cpuPercent := fmt.Sprintf("%.3f", float64(cs.Body.CPUStats.CPUUsage.TotalUsage*int64(100))/float64(cs.Body.CPUStats.SystemCPUUsage))
		return cpuPercent + "%"
	}
	return "0%"
}

func (cs ContainerStats) GetPids() int {
	if cs.Body != nil && cs.Body.PidsStats != nil {
		return cs.Body.PidsStats.Current
	}
	return 0
}

type StatsBodyPidsStats struct {
	Current int `json:"current"`
}

type StatsBodyCPUStats struct {
	OnlineCPUs     int                        `json:"online_cpus"`
	SystemCPUUsage int64                      `json:"system_cpu_usage"`
	CPUUsage       *StatsBodyCPUStatsCPUUsage `json:"cpu_usage"`
}

type StatsBodyCPUStatsCPUUsage struct {
	TotalUsage int64 `json:"total_usage"`
}

type StatsBody struct {
	Id        string              `json:"id"`
	Name      string              `json:"name"`
	Read      string              `json:"read"`
	Preread   string              `json:"preread"`
	NumProcs  int                 `json:"num_procs"`
	PidsStats *StatsBodyPidsStats `json:"pids_stats"`
	CPUStats  *StatsBodyCPUStats  `json:"cpu_stats"`
	//xxxxxx    string              `json:"xxxxxx"`
}
