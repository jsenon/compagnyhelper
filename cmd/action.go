// Package cmd is part of the cli
package cmd

import (
	"github.com/spf13/cobra"
)

const minArgs int = 1

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an object",
	Args:  cobra.MinimumNArgs(minArgs),
	//	Run: func(cmd *cobra.Command, args []string) { },
}

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe an object",
	Args:  cobra.MinimumNArgs(minArgs),
	//	Run: func(cmd *cobra.Command, args []string) { },
}

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open an object",
	Args:  cobra.MinimumNArgs(minArgs),
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(describeCmd)
	rootCmd.AddCommand(openCmd)
}
