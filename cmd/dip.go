package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/feo0o/dip/app"
	"github.com/urfave/cli/v2"
)

var dip *cli.App

func init() {
	dip = cli.NewApp()
	dip.Name = strings.ToLower(app.Name)
	dip.HelpName = app.Name
	dip.Usage = "A network tool for CIDR."
	dip.UsageText = fmt.Sprintf("%s [--version] [--help] <command> [<args>]", dip.Name)
	dip.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:     "version",
			Aliases:  []string{"v"},
			Value:    false,
			Usage:    "show version",
			Required: false,
		},
	}

	// functions
	dip.Action = func(c *cli.Context) error {
		if c.Bool("version") {
			fmt.Println(app.Version())
		}
		return nil
	}
}

func Run() error {
	return dip.Run(os.Args)
}
