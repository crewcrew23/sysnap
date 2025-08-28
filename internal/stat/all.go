package stat

import (
	"os"

	"github.com/crewcrew23/sysnap/internal/structs/result"
)

func GatherAll(file *os.File, duration int64) (*result.Result, error) {
	cpuUsage, err := CpuLoad(duration)
	if err != nil {
		return nil, err
	}

	memUsage, err := MemLoad()
	if err != nil {
		return nil, err
	}

	swapUsage, err := SwapLoad()
	if err != nil {
		return nil, err
	}

	loadAVG, err := LoadAVG()
	if err != nil {
		return nil, err
	}

	result := result.Result{
		Cpu:     cpuUsage,
		Memory:  memUsage,
		Swap:    swapUsage,
		LoadAvg: loadAVG,
	}

	return &result, nil
}
