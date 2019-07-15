package log

import (
	"fmt"

	"github.com/ioeXNetwork/ioeX.Client/rpc"

	"bytes"
	"encoding/json"
	"github.com/urfave/cli"
)

func debugAction(c *cli.Context) error {
	if c.NumFlags() == 0 {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	if c.Bool("getlevel") {
		result, err := rpc.CallAndUnmarshal("getloglevel", nil)
		if err != nil {
			fmt.Println("error: get log level failed,", err)
			return err
		}
		fmt.Println(result)
		return nil
	}

	if level := c.Int("setlevel"); level >= 0 {
		result, err := rpc.CallAndUnmarshal("setloglevel", rpc.Param("setlevel", level))
		if err != nil {
			fmt.Println("error: set debug info failed, ", err)
			return err
		}
		fmt.Println(result)
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
		Name:  "log",
		Usage: "set node log output level",
		Description: "With ioex-cli log, you could change node log output level.\n" +
			"\t levels are 0~5, the lower level the more logs will be print out, 0 means print out everything",
		ArgsUsage: "[args]",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "getlevel",
				Usage: "get log output level",
			},
			cli.IntFlag{
				Name:  "setlevel, l",
				Usage: "set log output level",
				Value: 5,
			},
		},
		Action: debugAction,
		OnUsageError: func(c *cli.Context, err error, isSubCommand bool) error {
			return cli.NewExitError(err, 1)
		},
	}
}
