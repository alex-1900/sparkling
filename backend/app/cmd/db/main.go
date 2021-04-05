package db

import (
	"github.com/alex-1900/sparkling/app/define"
	"github.com/alex-1900/sparkling/app/service"
	"github.com/spf13/cobra"
)

type Command struct {
	cmd *cobra.Command
	define.CmdHandler
}

func New(use string) *Command {
	var flagAction *string
	cmd := &cobra.Command{
		Use:   use,
		Short: "migrate database",
		Long:  `create or update database`,
		Run: func(c *cobra.Command, _ []string) {
			// 刷新数据库
			if *flagAction == "refresh" {
				db := service.GetDB()
				db.Migrate()
			} else {
				c.Help()
			}
		},
	}
	flagAction = cmd.PersistentFlags().StringP(
		"refresh",
		"r",
		"",
		"refresh database",
	)
	return &Command{cmd: cmd}
}

func (c *Command) Handle(g define.CmdGeter) define.CmdHandler {
	g.GetCmd().AddCommand(c.cmd)
	return c
}
