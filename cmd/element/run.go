package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
	"github.com/ehazlett/element/config"
	"github.com/ehazlett/element/server"
	"github.com/ehazlett/element/version"
	"github.com/sirupsen/logrus"
)

func runAction(c *cli.Context) error {
	logrus.Infof("element %s", version.FullVersion())

	var data string
	if configPath := c.GlobalString("config"); configPath != "" && data == "" {
		logrus.Debugf("loading config: file=%s", configPath)

		d, err := ioutil.ReadFile(configPath)
		switch {
		case os.IsNotExist(err):
			return fmt.Errorf("config not found: file=%s", configPath)
		case err == nil:
			data = string(d)
		default:
			return err
		}
	}

	if data == "" {
		return fmt.Errorf("You must specify a config from a file or environment variable")
	}

	config, err := config.ParseConfig(data)
	if err != nil {
		return err
	}

	srv, err := server.NewServer(config)
	if err != nil {
		return err
	}

	if err := srv.Run(); err != nil {
		return err
	}

	return nil
}
