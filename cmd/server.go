// launch server part

package cmd

import (
	"fmt"

	"github.com/jsenon/compagnyhelper/configs"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Launch server side",
	Version: fmt.Sprintf("Version: %v, build from: %v, on: %v\n", configs.Version, configs.GitCommit, configs.BuildDate),

	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	cobra.OnInitialize(initConfig)

}

func init() {
	rootCmd.AddCommand(serverCmd)
	cobra.OnInitialize(initConfig)

}
