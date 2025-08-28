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
			&cli.Int64Flag{
				Name:    "duration",
				Aliases: []string{"d"},
				Usage:   "duration for once snapshot",
				Value:   5,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			once := cmd.Bool("once")
			interval := cmd.Int64("interval")
			duration := cmd.Int64("duration")
			workTime := cmd.Int64("work-time")
			output := cmd.String("output")

			fmt.Println("Once", once)
			fmt.Println("interval", interval)
			fmt.Println("duration", duration)
			fmt.Println("workTime", workTime)
			fmt.Println("output", output)

			err := startup.RunOnce(output, duration)
			return err
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(cmd)
	}
}
