package info

import (
	"fmt"

	"github.com/ioeXNetwork/ioeX.Client/rpc"

	"bytes"
	"encoding/json"
	"github.com/urfave/cli"
)

func infoAction(c *cli.Context) error {
	if c.NumFlags() == 0 {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	if c.Bool("connections") {
		result, err := rpc.Call("getconnectioncount", nil)
		if err != nil {
			fmt.Println("error: get node connections failed,", err)
			return err
		}
		printFormat(result)
		return nil
	}

	if c.Bool("neighbor") {
		result, err := rpc.Call("getneighbors", nil)
		if err != nil {
			fmt.Println("error: get node neighbors info failed,", err)
			return err
		}
		printFormat(result)
		return nil
	}

	if c.Bool("state") {
		result, err := rpc.Call("getnodestate", nil)
		if err != nil {
			fmt.Println("error: get node state info failed,", err)
			return err
		}
		printFormat(result)
		return nil
	}

	if c.Bool("currentheight") {
		result, err := rpc.Call("getcurrentheight", nil)
		if err != nil {
			fmt.Println("error: get block count failed,", err)
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
		Name:        "info",
		Usage:       "show node information",
		Description: "With ioex-cli info, you could look up node status, query blocks, transactions, etc.",
		ArgsUsage:   "[args]",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "connections",
				Usage: "see how many peers are connected with current node",
			},
			cli.BoolFlag{
				Name:  "neighbor, nbr",
				Usage: "show neighbor nodes information",
			},
			cli.BoolFlag{
				Name:  "state",
				Usage: "show current node status",
			},
			cli.BoolFlag{
				Name:  "currentheight, height",
				Usage: "show blockchain height on current node",
			},
		},
		Action: infoAction,
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return cli.NewExitError(err, 1)
		},
	}
}
