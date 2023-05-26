package logging

import (
	"github.com/orkarstoft/kscale/pkg/config"
	"github.com/rs/zerolog"
)

// Initialize rs/zerolog logging
func Init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if config.Config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		return
	}
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
