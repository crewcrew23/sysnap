package oncemode

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/crewcrew23/sysnap/internal/stat"
	"github.com/crewcrew23/sysnap/internal/structs/result"
	"github.com/shirou/gopsutil/cpu"
)

func Run(file *os.File, duration int64) error {
	log.Println("Start Once")

	usage, err := cpuLoad(duration)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(file).Encode(result.Result{Cpu: usage}); err != nil {
		return err
	}

	return nil
}

func cpuLoad(duration int64) ([]float64, error) {
	var usage []float64

	t1, err := cpu.Times(true)
	if err != nil {
		return nil, err
	}

	time.Sleep((time.Second * time.Duration(duration)))

	t2, err := cpu.Times(true)
	if err != nil {
		return nil, err
	}

	for i := range len(t1) {
		usage = append(usage, stat.Usage(&t1[i], &t2[i]))
	}

	return usage, nil
}
