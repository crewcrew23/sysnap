package oncemode

import (
	"encoding/json"
	"log"
	"os"

	"github.com/crewcrew23/sysnap/internal/stat"
)

func Run(file *os.File, disks []string, duration int64) error {
	log.Println("Start Once")

	result, err := stat.GatherAll(file, disks, duration)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(file).Encode(result); err != nil {
		return err
	}

	return nil
}
