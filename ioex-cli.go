package main

import (
	"os"
	"sort"

	"github.com/ioeXNetwork/ioeX.Client/cli/info"
	cliLog "github.com/ioeXNetwork/ioeX.Client/cli/log"
	"github.com/ioeXNetwork/ioeX.Client/cli/mine"
	"github.com/ioeXNetwork/ioeX.Client/cli/wallet"
	"github.com/ioeXNetwork/ioeX.Client/log"
	"github.com/urfave/cli"
)

var Version string

func init() {
	log.InitLog()
}

func main() {
	app := cli.NewApp()
	app.Name = "ela-cli"
	app.Version = Version
	app.HelpName = "ela-cli"
	app.Usage = "command line tool for ELA blockchain"
	app.UsageText = "ela-cli [global options] command [command options] [args]"
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
