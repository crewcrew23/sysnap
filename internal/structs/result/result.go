package result

import "github.com/shirou/gopsutil/load"

type Memory struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
}

type Result struct {
	Swap    *Memory       `json:"Swap"`
	Memory  *Memory       `json:"Memory"`
	Cpu     []float64     `json:"Cpu"`
	LoadAvg *load.AvgStat `json:"LoadAVG"`
}
