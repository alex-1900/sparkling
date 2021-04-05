package server

import (
	"github.com/alex-1900/sparkling/app/define"
	"github.com/spf13/cobra"
)

type Command struct {
	cmd *cobra.Command
	define.CmdHandler
}

func New(use string) *Command {
	var flagAction *string
	var flagPort *string
	cmd := &cobra.Command{
		Use:   use,
		Short: "web server",
		Long:  `Start or stop the HTTP server of sparkling`,
		Run: func(c *cobra.Command, _ []string) {
			if *flagAction == "start" {
				startServer("0.0.0.0:" + *flagPort)
			}
		},
	}
	flagAction = cmd.PersistentFlags().StringP(
		"execute",
		"e",
		"start",
		"\"start\" or \"stop\"",
	)
	flagPort = cmd.PersistentFlags().StringP(
		"port",
		"p",
		"8080",
		"The HTTP server port for sparkling",
	)
	return &Command{cmd: cmd}
}

func (c *Command) Handle(g define.CmdGeter) define.CmdHandler {
	g.GetCmd().AddCommand(c.cmd)
	return c
}
