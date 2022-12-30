package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/feo0o/dip/app"
	"github.com/urfave/cli/v2"
)

var dip *cli.App

func Run() error {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Version)
	}

	dip = cli.NewApp()
	dip.Name = strings.ToLower(app.Name)
	dip.HelpName = app.Name
	dip.Usage = "A network toolkit"
	dip.UsageText = fmt.Sprintf("%s [--version] [--help] <command> [<args>]", dip.Name)
	dip.Version = app.Version()

	dip.Commands = []*cli.Command{calcCmd}
	return dip.Run(os.Args)
}
