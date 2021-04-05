package cmd

import (
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
		"CONFIG FILE PATH",
		"config file (default is $HOME/.cobra.yaml)",
	)
}

// Get 获取 cobra root 实例
func Get() *cobra.Command {
	return cmd
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
