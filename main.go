package main

import (
	"os"

	"github.com/leodotcloud/log"
	logserver "github.com/leodotcloud/log/server"
	"github.com/leodotcloud/swiss-army-knife/server"
	"github.com/urfave/cli"
)

// VERSION of the application, that can defined during build time
var VERSION = "v0.0.0-dev"

const (
	appName            = "swiss-army-knife"
	portArg            = "port"
	useMetadataArg     = "use-rancher-metadata"
	metadataAddressArg = "metadata-address"
	alphabetArg        = "alphabet"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Version = VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   metadataAddressArg,
			Value:  server.DefaultMetadataAddress,
			EnvVar: "RANCHER_METADATA_ADDRESS",
		},
		cli.StringFlag{
			Name:   portArg,
			Value:  server.DefaultServerPort,
			EnvVar: "PORT",
		},
		cli.StringFlag{
			Name:   alphabetArg,
			Usage:  "Run the web server with the given alphabet",
			EnvVar: "NATO_ALPHABET",
		},
		cli.BoolFlag{
			Name:   useMetadataArg,
			Usage:  "Use Rancher metadata for querying information about self",
			EnvVar: "USE_RANCHER_METADATA",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Turn on debug logging",
			EnvVar: "DEBUG",
		},
	}

	app.Action = run
	app.Run(os.Args)
}

func run(c *cli.Context) error {
	logserver.StartServerWithDefaults()
	if c.Bool("debug") {
		log.SetLevelString("debug")
	}

	s, err := server.NewServer(
		c.String(portArg),
		c.String(metadataAddressArg),
		c.String(alphabetArg),
		c.Bool(useMetadataArg),
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
