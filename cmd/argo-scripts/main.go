package main

import (
	"os"

	"github.com/dictybase-playground/argo-scripts/internal/app/validate"
	"github.com/dictybase-playground/argo-scripts/internal/app/webhooks"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "argo-scripts"
	app.Usage = "cli for scripts related to argo workflows and events"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log-format",
			Usage: "format of the logging out, either of json or text.",
			Value: "json",
		},
		cli.StringFlag{
			Name:  "log-level",
			Usage: "log level for the application",
			Value: "debug",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "create-webhooks",
			Usage:  "creates new github webhooks based on given input yaml",
			Action: webhooks.RunCreateWebhooks,
			Before: validate.ValidateServerArgs,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input-file, i",
					Usage: "input yaml file",
					Value: "values.yaml",
				},
				cli.StringFlag{
					Name:  "output-file, o",
					Usage: "output yaml file",
					Value: "hooks.yaml",
				},
				cli.StringFlag{
					Name:  "github-access-token, g",
					Usage: "github access token",
				},
			},
		},
	}
	app.Run(os.Args)
}
