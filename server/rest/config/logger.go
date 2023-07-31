package config

import (
	"github.com/rs/zerolog"
)

func InitZeroLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
