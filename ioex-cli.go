package main

import (
	"os"
	"sort"

	"github.com/ioeXNetwork/ioeX.Client/cli/info"
	cliLog "github.com/ioeXNetwork/ioeX.Client/cli/log"
	"github.com/ioeXNetwork/ioeX.Client/cli/wallet"
	"github.com/ioeXNetwork/ioeX.Client/log"
	"github.com/urfave/cli"
)

func init() {
	log.InitLog()
}

func main() {
	app := cli.NewApp()
	app.Name = "ioex-cli"
	app.Version = "0.1.1"
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
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))

	app.Run(os.Args)
}
