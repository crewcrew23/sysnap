package startup

import (
	"os"

	"github.com/crewcrew23/sysnap/internal/modes/oncemode"
)

func RunOnce(out string, duration int64) error {
	file, err := os.Create(out)
	if err != nil {
		return err
	}

	err = oncemode.Run(file, duration)
	return err
}
