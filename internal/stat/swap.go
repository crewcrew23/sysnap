package stat

import (
	"github.com/crewcrew23/sysnap/internal/structs/result"
	"github.com/shirou/gopsutil/mem"
)

func SwapLoad() (*result.Memory, error) {
	swap, err := mem.SwapMemory()
	if err != nil {
		return nil, err
	}

	swapStat := result.Memory{
		Total: swap.Total / 1024 / 1024,
		Used:  swap.Used / 1024 / 1024,
	}

	return &swapStat, nil
}
