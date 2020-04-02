// Package cmd is part of the cli
package cmd

import (
	"fmt"

	"github.com/jsenon/compagnyhelper/configs"
	mylog "github.com/jsenon/compagnyhelper/internal/log"
	"github.com/jsenon/compagnyhelper/internal/web"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var inputdir string
var disabletrace bool
var jaegerurl string

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Launch server side",
	Version: fmt.Sprintf("Version: %v, build from: %v, on: %v\n", configs.Version, configs.GitCommit, configs.BuildDate),

	Run: func(cmd *cobra.Command, args []string) {
		log.Logger = log.With().Str("Service", configs.Service).Logger()
		log.Logger = log.With().Str("Version", configs.Version).Logger()

		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		if loglevel {
			err := mylog.SetDebug()
			if err != nil {
				log.Error().Msgf("Could not set loglevel to debug: %v", err)
			}
			log.Debug().Msg("Log level set to Debug")
		}
		web.Serve()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	serverCmd.PersistentFlags().StringVar(&inputdir, "inputdir", "",
		"Set folder where input file are located")

	err := viper.BindPFlag("inputdir", serverCmd.PersistentFlags().Lookup("env"))
	if err != nil {
		log.Error().Msgf("Error binding env value: %v", err.Error())
	}

	serverCmd.PersistentFlags().BoolVar(&disabletrace, "disabletrace", false, "Disable the trace")

	err = viper.BindPFlag("disabletrace", serverCmd.PersistentFlags().Lookup("disabletrace"))
	if err != nil {
		log.Error().Msgf("Error binding disabletrace value: %v", err.Error())
	}

	serverCmd.PersistentFlags().StringVar(&jaegerurl, "jaegerurl", "",
		"Set jaegger agent endpoint (without port, without http://)")

	err = viper.BindPFlag("jaegerurl", serverCmd.PersistentFlags().Lookup("jaegerurl"))
	if err != nil {
		log.Error().Msgf("Error binding jaegerurl value: %v", err.Error())
	}

	viper.SetDefault("jaegerurl", "")
}

func init() {
	rootCmd.AddCommand(serverCmd)
	cobra.OnInitialize(initConfig)
}
