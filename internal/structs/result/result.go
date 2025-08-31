package result

import "github.com/shirou/gopsutil/load"

type Memory struct {
	Total     uint64 `json:"total"`
	Used      uint64 `json:"used"`
	Free      uint64 `json:"free"`
	Available uint64 `json:"available"`
	Cache     uint64 `json:"cache"`
}

type Swap struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
}

type Uptime struct {
	Hours   uint64
	Minutes uint8
	Seconds uint8
}

type CPU struct {
	Cores        int       `json:"cores"`
	UsagePerCore []float64 `json:"usagePerCore"`
	AvgUsage     float64   `json:"avgUsage"`
	Idle         float64   `json:"idle"`
	IoWait       float64   `json:"iowait"`
}

type Result struct {
	Swap    *Swap         `json:"Swap"`
	Memory  *Memory       `json:"Memory"`
	Cpu     *CPU          `json:"Cpu"`
	LoadAvg *load.AvgStat `json:"LoadAVG"`
	Uptime  *Uptime       `json:"Uptime"`
}
