// Package log ...
package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// SetDebug set log level to debug
func SetDebug() error {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	err := os.Setenv("LOGLEVEL", "debug")
	if err != nil {
		log.Error().Msgf("Error %s", err.Error())
		return err
	}

	return nil
}
