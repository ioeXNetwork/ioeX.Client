package wallet

import (
	"fmt"
	"os"

	"github.com/ioeXNetwork/ioeX.Client/rpc"

	"bytes"
	"encoding/json"
	"github.com/urfave/cli"
)

func walletAction(context *cli.Context) error {
	if context.NumFlags() == 0 {
		cli.ShowSubcommandHelp(context)
		os.Exit(0)
	}

	if address := context.String("balance"); address != "" {
		result, err := rpc.Call("getreceivedbyaddress", rpc.Param("address", address))
		if err != nil {
			fmt.Println("error: get balance failed,", err)
			return err
		}
		printFormat(result)
		return nil
	}

	return nil
}

func printFormat(data []byte) {
	buf := new(bytes.Buffer)
	json.Indent(buf, data, "", "    ")
	fmt.Println(string(buf.Bytes()))
}

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:        "wallet",
		Usage:       "wallet operations",
		Description: "With ioex-cli wallet, check account balance.",
		ArgsUsage:   "[args]",
		Flags: []cli.Flag{

			cli.StringFlag{
				Name:  "balance",
				Usage: "show account balance.",
			},
		},
		Action: walletAction,
		OnUsageError: func(c *cli.Context, err error, subCommand bool) error {
			return cli.NewExitError(err, 1)
		},
	}
}
