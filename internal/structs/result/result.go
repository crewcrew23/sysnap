package result

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type CoreStat struct {
	Number int
	Total  float64
}

type Result struct {
	Swap          *mem.SwapMemoryStat            `json:"Swap"`
	VirtialMemory *mem.VirtualMemoryStat         `json:"VirtualMemory"`
	Cpu           []float64                      `json:"Cpu"`
	PartitionStat []disk.PartitionStat           `json:"PartitionStat"`
	IOCounter     map[string]disk.IOCountersStat `json:"IOCounter"`
}
