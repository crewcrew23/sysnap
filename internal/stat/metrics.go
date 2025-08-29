package stat

import (
	"os"
	"time"

	"github.com/crewcrew23/sysnap/internal/structs/result"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
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

func LoadAVG() (*load.AvgStat, error) {
	avg, err := load.Avg()
	if err != nil {
		return nil, err
	}

	return avg, nil
}

func Uptime() (*result.Uptime, error) {
	uptime, err := host.Uptime()
	if err != nil {
		return nil, err
	}

	h := uptime / 3600
	m := (uptime % 3600) / 60
	s := uptime % 60

	upt := &result.Uptime{
		Hours:   h,
		Minutes: uint8(m),
		Seconds: uint8(s),
	}

	return upt, nil
}

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

	uptime, err := Uptime()
	if err != nil {
		return nil, err
	}

	result := result.Result{
		Cpu:     cpuUsage,
		Memory:  memUsage,
		Swap:    swapUsage,
		LoadAvg: loadAVG,
		Uptime:  uptime,
	}

	return &result, nil
}
