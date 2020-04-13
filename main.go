package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "nali",
		Usage: "get ip/domain 's location info.",
		Action: func(c *cli.Context) error {
			getIpInfo(c.Args().First(), c.String("lang"))
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "zh",
				Usage:   "language for the output, zh/en available",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
