package stat

import (
	"github.com/crewcrew23/sysnap/internal/structs/result"
	"github.com/shirou/gopsutil/mem"
)

func MemLoad() (*result.Memory, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	memStat := result.Memory{
		Total: memory.Total / 1024 / 1024,
		Used:  memory.Used / 1024 / 1024,
	}

	return &memStat, nil
}
