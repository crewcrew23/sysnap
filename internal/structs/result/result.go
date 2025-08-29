package result

import "github.com/shirou/gopsutil/load"

type Memory struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
}

type Uptime struct {
	Hours   uint64
	Minutes uint8
	Seconds uint8
}

type Result struct {
	Swap    *Memory       `json:"Swap"`
	Memory  *Memory       `json:"Memory"`
	Cpu     []float64     `json:"Cpu"`
	LoadAvg *load.AvgStat `json:"LoadAVG"`
	Uptime  *Uptime       `json:"uptime"`
}
