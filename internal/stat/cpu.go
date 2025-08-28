package stat

import "github.com/shirou/gopsutil/cpu"

func Usage(t1, t2 *cpu.TimesStat) float64 {
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
