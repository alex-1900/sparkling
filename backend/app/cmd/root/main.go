package root

import (
	"fmt"

	"github.com/alex-1900/sparkling/app/define"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Command struct {
	cmd *cobra.Command
	define.CmdGeter
}

var (
	cmd = &cobra.Command{
		Use:   "sparkling",
		Short: "A personal knowledge base backend",
		Long: `A personal knowledge base backend built with alex-1900 in Go.
Communicate with me by email: omytty@126.com`,
		Run: func(c *cobra.Command, _ []string) {
			c.Help()
		},
	}
	cfgFile *string
)

func init() {
	cobra.OnInitialize(initConfig)
	cfgFile = cmd.PersistentFlags().StringP(
		"config",
		"c",
		"",
		"config file (default is $HOME/.cobra.yaml)",
	)
}

func New() *Command {
	return &Command{cmd: cmd}
}

func (c *Command) GetCmd() *cobra.Command {
	return c.cmd
}

// https://github.com/spf13/cobra#readme
func initConfig() {
	if *cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(*cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
