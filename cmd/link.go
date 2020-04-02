// Object link

package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var env string

var linkGetCmd = &cobra.Command{
	Use:   "link",
	Short: "Object link",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var linkDescCmd = &cobra.Command{
	Use:   "link",
	Short: "Object link",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var linkOpCmd = &cobra.Command{
	Use:   "link",
	Short: "Object link",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	cobra.OnInitialize(initConfig)
	linkGetCmd.PersistentFlags().StringVar(&env, "environment", "dev",
		"Set environment variable where retrieve object")

	err := viper.BindPFlag("env", linkGetCmd.PersistentFlags().Lookup("env"))
	if err != nil {
		log.Error().Msgf("Error binding env value: %v", err.Error())
	}
	linkDescCmd.PersistentFlags().StringVar(&env, "environment", "dev",
		"Set environment variable where retrieve object")

	err = viper.BindPFlag("env", linkDescCmd.PersistentFlags().Lookup("env"))
	if err != nil {
		log.Error().Msgf("Error binding env value: %v", err.Error())
	}
	linkOpCmd.PersistentFlags().StringVar(&env, "environment", "dev",
		"Set environment variable where retrieve object")

	err = viper.BindPFlag("env", linkOpCmd.PersistentFlags().Lookup("env"))
	if err != nil {
		log.Error().Msgf("Error binding env value: %v", err.Error())
	}

}

func init() {
	getCmd.AddCommand(linkGetCmd)
	describeCmd.AddCommand(linkDescCmd)
	openCmd.AddCommand(linkOpCmd)
	cobra.OnInitialize(initConfig)
}
