package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/crewcrew23/sysnap/internal/startup"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name: "sysnap",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "once",
				Usage:   "`take one spanshot",
				Aliases: []string{"on"},
			},
			&cli.Int64Flag{
				Name:    "interval",
				Aliases: []string{"i"},
				Usage:   "interval with that will be takes snapshots",
				Value:   30,
			},
			&cli.Int64Flag{
				Name:    "work-time",
				Aliases: []string{"wt"},
				Usage:   "runtime in seconds by default infinite",
				Value:   0,
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"out"},
				Usage:   "path for output file",
				Value:   "sysnap-result.json",
			},
			&cli.StringSliceFlag{
				Name:    "disk",
				Aliases: []string{"d"},
				Usage:   "array of disks path",
				Value:   []string{""},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// once := cmd.Bool("once")
			interval := cmd.Int64("interval")
			// workTime := cmd.Int64("work-time")
			output := cmd.String("output")
			disks := cmd.StringSlice("disk")

			fmt.Println(disks)
			err := startup.RunOnce(output, disks, interval)
			return err
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
