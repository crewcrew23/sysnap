package oncemode

import (
	"encoding/json"
	"log"
	"os"

	"github.com/crewcrew23/sysnap/internal/stat"
	"github.com/crewcrew23/sysnap/internal/structs/result"
)

func Run(file *os.File, duration int64) error {
	log.Println("Start Once")

	cpuUsage, err := stat.CpuLoad(duration)
	if err != nil {
		return err
	}

	memUsage, err := stat.MemLoad()
	if err != nil {
		return err
	}

	swapUsage, err := stat.SwapLoad()
	if err != nil {
		return err
	}

	if err := json.NewEncoder(file).Encode(result.Result{
		Cpu:    cpuUsage,
		Memory: memUsage,
		Swap:   swapUsage,
	}); err != nil {
		return err
	}

	return nil
}
