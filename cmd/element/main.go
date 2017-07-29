package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/ehazlett/element/version"
	"github.com/sirupsen/logrus"
)

func main() {
	app := cli.NewApp()
	app.Name = "element"
	app.Version = version.FullVersion()
	app.Author = "@ehazlett"
	app.Email = ""
	app.Usage = "container execution platform"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug logging",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "path to configuration file",
		},
	}
	app.Action = runAction
	app.Before = func(c *cli.Context) error {
		if c.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
