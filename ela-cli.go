package main

import (
	"os"
	"sort"

	"github.com/ioeX/ioeX.Client/cli/info"
	"github.com/ioeX/ioeX.Client/cli/wallet"
	"github.com/ioeX/ioeX.Client/cli/mine"
	"github.com/ioeX/ioeX.Client/log"
	cliLog "github.com/ioeX/ioeX.Client/cli/log"
	"github.com/urfave/cli"
)

var Version string

func init() {
	log.InitLog()
}

func main() {
	app := cli.NewApp()
	app.Name = "ioex-cli"
	app.Version = Version
	app.HelpName = "ioex-cli"
	app.Usage = "command line tool for IOEX blockchain"
	app.UsageText = "ioex-cli [global options] command [command options] [args]"
	app.HideHelp = false
	app.HideVersion = false
	//commands
	app.Commands = []cli.Command{
		*cliLog.NewCommand(),
		*info.NewCommand(),
		*wallet.NewCommand(),
		*mine.NewCommand(),
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))

	app.Run(os.Args)
}
