// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/jsenon/compagnyhelper/configs"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var version bool
var loglevel bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "compagnyhelper",
	Short:   "A swiss knife helper compagny",
	Long:    `An helper that can retrieve usefull compagny links`,
	Version: fmt.Sprintf("Version: %v, build from: %v, on: %v\n", configs.Version, configs.GitCommit, configs.BuildDate),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.SetVersionTemplate("trst")

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.compagnyhelper.yaml)")
	rootCmd.PersistentFlags().BoolVar(&loglevel, "debug", false, "Set log level to Debug")

	err := viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("debug"))
	if err != nil {
		log.Error().Msgf("Error binding loglevel value: %v", err.Error())
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".compagnyhelper" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".compagnyhelper")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
