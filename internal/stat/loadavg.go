package stat

import "github.com/shirou/gopsutil/load"

func LoadAVG() (*load.AvgStat, error) {
	avg, err := load.Avg()
	if err != nil {
		return nil, err
	}

	return avg, nil
}
