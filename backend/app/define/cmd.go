package define

import "github.com/spf13/cobra"

type CmdGeter interface {
	GetCmd() *cobra.Command
}

type CmdHandler interface {
	Handle(CmdGeter) CmdHandler
}
