package main

import (
	"fmt"
	"os"

	"github.com/leodotcloud/log"
	logserver "github.com/leodotcloud/log/server"
	"github.com/leodotcloud/swiss-army-knife/server"
	"github.com/urfave/cli"
)

// VERSION of the application, that can defined during build time
var VERSION = "v0.0.0-dev"

const (
	appName     = "swiss-army-knife"
	portArg     = "port"
	alphabetArg = "alphabet"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Version = VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   portArg,
			Value:  server.DefaultServerPort,
			EnvVar: "PORT",
		},
		cli.StringFlag{
			Name:   alphabetArg,
			Usage:  "Run the web server with the given alphabet",
			EnvVar: "ALPHABET",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Turn on debug logging",
			EnvVar: "DEBUG",
		},
	}

	app.Action = run
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	logserver.StartServerWithDefaults()
	if c.Bool("debug") {
		if err := log.SetLevelString("debug"); err != nil {
			return err
		}
	}

	if c.Bool("version") {
		fmt.Println(c.App.Version)
		return nil
	}
	s, err := server.NewServer(
		c.String(portArg),
		c.String(alphabetArg),
	)
	if err != nil {
		log.Errorf("Error creating new server")
		return err
	}

	if err := s.Run(); err != nil {
		log.Errorf("Failed to start: %v", err)
	}

	<-s.GetExitChannel()
	log.Infof("Program exiting")
	return nil
}
