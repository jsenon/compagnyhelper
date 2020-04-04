// Package cmd is part of the cli
package cmd

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/jsenon/compagnyhelper/internal/link"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var env string
var app string

const nbrArgs int = 1

var linkGetCmd = &cobra.Command{
	Use:   "link",
	Short: "Object link",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < nbrArgs {
			return nil
		}
		app = args[0]
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		mystringurl := viper.GetString("URL_SERVER")
		myurl, err := url.Parse(mystringurl)
		if err != nil {
			fmt.Printf("Error getting the server URL: %v", err)
			return
		}
		if app == "" {
			link.Get(ctx, myurl.String())
		} else {
			link.GetDetail(ctx, env, app, myurl.String())
		}
	},
}
var linkDescCmd = &cobra.Command{ //nolint: dupl
	Use:   "link",
	Short: "Object link",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < nbrArgs {
			return errors.New("requires a link argument")
		}
		app = args[0]
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		mystringurl := viper.GetString("URL_SERVER")
		myurl, err := url.Parse(mystringurl)

		if err != nil {
			fmt.Printf("Error getting the server URL: %v", err)
			return
		}
		link.Describe(ctx, env, app, myurl.String())

	},
}
var linkOpCmd = &cobra.Command{ //nolint: dupl
	Use:   "link",
	Short: "Object link",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != nbrArgs {
			return errors.New("requires a link argument")
		}
		app = args[0]
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		mystringurl := viper.GetString("URL_SERVER")
		myurl, err := url.Parse(mystringurl)

		if err != nil {
			fmt.Printf("Error getting the server URL: %v", err)
			return
		}
		link.OpenLink(ctx, env, app, myurl.String())

	},
}

func init() {
	cobra.OnInitialize(initConfig)

	linkGetCmd.PersistentFlags().StringVar(&env, "n", "all",
		"Set environment variable where retrieve object")

	err := viper.BindPFlag("env", linkGetCmd.PersistentFlags().Lookup("env"))
	if err != nil {
		log.Error().Msgf("Error binding env value: %v", err.Error())
	}

	linkDescCmd.PersistentFlags().StringVar(&env, "n", "all",
		"Set environment variable where retrieve object")

	err = viper.BindPFlag("env", linkDescCmd.PersistentFlags().Lookup("env"))
	if err != nil {
		log.Error().Msgf("Error binding env value: %v", err.Error())
	}

	linkOpCmd.PersistentFlags().StringVar(&env, "n", "all",
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
