package stat

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func coreLoad(t1, t2 *cpu.TimesStat) float64 {
	idle1 := t1.Idle + t1.Iowait
	idle2 := t2.Idle + t2.Iowait

	nonIdle1 := t1.User + t1.System + t1.Nice + t1.Irq + t1.Softirq + t1.Steal
	nonIdle2 := t2.User + t2.System + t2.Nice + t2.Irq + t2.Softirq + t2.Steal

	total1 := idle1 + nonIdle1
	total2 := idle2 + nonIdle2

	totald := total2 - total1
	idled := idle2 - idle1

	return (totald - idled) / totald * 100
}

func CpuLoad(duration int64) ([]float64, error) {
	t1, err := cpu.Times(true)
	if err != nil {
		return nil, err
	}

	time.Sleep((time.Second * time.Duration(duration)))

	t2, err := cpu.Times(true)
	if err != nil {
		return nil, err
	}

	usage := make([]float64, len(t1))

	for i := range len(t1) {
		usage[i] = coreLoad(&t1[i], &t2[i])
	}

	return usage, nil
}
