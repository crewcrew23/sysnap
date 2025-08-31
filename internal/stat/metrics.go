package stat

import (
	"fmt"
	"os"
	"time"

	"github.com/crewcrew23/sysnap/internal/structs/result"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func coresLoad(t1, t2 []cpu.TimesStat) []float64 {

	usage := make([]float64, len(t1))

	for i := range len(t1) {
		idle1 := (&t1[i]).Idle + (&t1[i]).Iowait
		idle2 := (&t2[i]).Idle + (&t2[i]).Iowait

		nonIdle1 := (&t1[i]).User + (&t1[i]).System + (&t1[i]).Nice + (&t1[i]).Irq + (&t1[i]).Softirq + (&t1[i]).Steal
		nonIdle2 := (&t2[i]).User + (&t2[i]).System + (&t2[i]).Nice + (&t2[i]).Irq + (&t2[i]).Softirq + (&t2[i]).Steal

		total1 := idle1 + nonIdle1
		total2 := idle2 + nonIdle2

		totalD := total2 - total1
		idleD := idle2 - idle1

		if totalD == 0 || idleD == 0 {
			usage[i] = 0
			continue
		}

		res := (totalD - idleD) / totalD * 100

		usage[i] = res
	}

	return usage
}

func cpuIdle(t1, t2 []cpu.TimesStat) float64 {
	slIdle := make([]float64, len(t1))

	for i := range len(t1) {
		idle1 := (&t1[i]).Idle + (&t1[i]).Iowait
		idle2 := (&t2[i]).Idle + (&t2[i]).Iowait

		nonIdle1 := (&t1[i]).User + (&t1[i]).System + (&t1[i]).Nice + (&t1[i]).Irq + (&t1[i]).Softirq + (&t1[i]).Steal
		nonIdle2 := (&t2[i]).User + (&t2[i]).System + (&t2[i]).Nice + (&t2[i]).Irq + (&t2[i]).Softirq + (&t2[i]).Steal

		total1 := idle1 + nonIdle1
		total2 := idle2 + nonIdle2

		totalD := total2 - total1
		idleD := idle2 - idle1

		if totalD == 0 || idleD == 0 {
			slIdle[i] = 0
			continue
		}

		slIdle[i] = (idleD / totalD) * 100
	}

	var sum float64

	for _, v := range slIdle {
		sum += v
	}

	return sum / float64(len(slIdle))

}

func ioWait(t1, t2 []cpu.TimesStat) float64 {
	ioWaitSl := make([]float64, len(t1))

	for i := range len(t1) {
		nonIdle1 := (&t1[i]).User + (&t1[i]).System + (&t1[i]).Nice + (&t1[i]).Irq + (&t1[i]).Softirq + (&t1[i]).Steal
		nonIdle2 := (&t2[i]).User + (&t2[i]).System + (&t2[i]).Nice + (&t2[i]).Irq + (&t2[i]).Softirq + (&t2[i]).Steal

		total1 := (&t1[i]).Iowait + nonIdle1
		total2 := (&t2[i]).Iowait + nonIdle2

		totalD := total2 - total1
		idleD := (&t2[i]).Iowait - (&t1[i]).Iowait

		if totalD == 0 || idleD == 0 {
			ioWaitSl[i] = 0
			continue
		}

		ioWaitSl[i] = (idleD / totalD) * 100
	}

	var sum float64

	for _, v := range ioWaitSl {
		sum += v
	}

	return sum / float64(len(t1))
}

func avgCpuUsage(t1, t2 []cpu.TimesStat) float64 {
	avgUsage := make([]float64, len(t1))

	for i := range len(t1) {
		idle1 := (&t1[i]).Idle + (&t1[i]).Iowait
		idle2 := (&t2[i]).Idle + (&t2[i]).Iowait

		nonIdle1 := (&t1[i]).User + (&t1[i]).System + (&t1[i]).Nice + (&t1[i]).Irq + (&t1[i]).Softirq + (&t1[i]).Steal
		nonIdle2 := (&t2[i]).User + (&t2[i]).System + (&t2[i]).Nice + (&t2[i]).Irq + (&t2[i]).Softirq + (&t2[i]).Steal

		total1 := idle1 + nonIdle1
		total2 := idle2 + nonIdle2

		totalD := total2 - total1
		idleD := idle2 - idle1

		if totalD == 0 || idleD == 0 {
			avgUsage[i] = 0
			continue
		}

		res := (totalD - idleD) / totalD * 100

		avgUsage[i] = res
	}

	var sum float64

	for i := range avgUsage {
		sum += avgUsage[i]
	}

	return sum / float64(len(t1))
}

func CpuLoad(duration int64) (*result.CPU, error) {
	t1, err := cpu.Times(true)
	if err != nil {
		return nil, err
	}

	time.Sleep((time.Second * time.Duration(duration)))

	t2, err := cpu.Times(true)
	if err != nil {
		return nil, err
	}

	cores, err := cpu.Counts(true)
	if err != nil {
		return nil, err
	}

	usage := coresLoad(t1, t2)
	cpuIdle := cpuIdle(t1, t2)
	ioWait := ioWait(t1, t2)
	avgUsage := avgCpuUsage(t1, t2)

	cpuL := &result.CPU{
		Cores:        cores,
		UsagePerCore: usage,
		Idle:         cpuIdle,
		IoWait:       ioWait,
		AvgUsage:     avgUsage,
	}

	return cpuL, nil
}

func MemLoad() (*result.Memory, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	memStat := result.Memory{
		Total:     memory.Total / 1024 / 1024,
		Used:      memory.Used / 1024 / 1024,
		Free:      memory.Free / 1024 / 1024,
		Available: memory.Available / 1024 / 1024,
		Cache:     memory.Cached / 1024 / 1024,
	}

	return &memStat, nil
}

func SwapLoad() (*result.Swap, error) {
	swap, err := mem.SwapMemory()
	if err != nil {
		return nil, err
	}

	swapStat := result.Swap{
		Total: swap.Total / 1024 / 1024,
		Used:  swap.Used / 1024 / 1024,
		Free:  swap.Free / 1024 / 1024,
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

func diskUsage(disks []string) ([]result.DiskWrapper, error) {
	if disks[0] == "" {
		return []result.DiskWrapper{}, nil
	}

	disksWrp := make([]result.DiskWrapper, 0, len(disks))
	var scopeErr error

	for _, v := range disks {
		disk, err := disk.Usage(v)
		if err != nil {
			scopeErr = fmt.Errorf("%s", err.Error())
			break
		}

		disksWrp = append(disksWrp, result.DiskWrapper{
			Path: v,
			Data: &result.Disk{
				Total:        disk.Total / 1024 / 1024,
				Usge:         disk.Used / 1024 / 1024,
				UsagePercent: disk.UsedPercent,
				Free:         disk.Free / 1024 / 1024,
			},
		})
	}

	if scopeErr != nil {
		return nil, scopeErr
	}

	return disksWrp, nil
}

func GatherAll(file *os.File, disks []string, duration int64) (*result.Result, error) {
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

	disk, err := diskUsage(disks)
	if err != nil {
		return nil, err
	}

	result := result.Result{
		Cpu:     cpuUsage,
		Memory:  memUsage,
		Swap:    swapUsage,
		LoadAvg: loadAVG,
		Uptime:  uptime,
		Disks:   disk,
	}

	return &result, nil
}
