package config

import (
	"os"

	"github.com/pengsrc/go-shared/reopen"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger sotres the config for logger.
type Logger struct {
	Level  string `yaml:"level"`
	Output string `yaml:"output"`
}

// Prepare will create logger instance.
func (l *Logger) Prepare() error {
	switch l.Output {
	case "stdout":
		log.Logger = log.
			Output(os.Stdout).
			Level(parseLevel(l.Level))
	default:
		f, err := reopen.NewFileWriter(l.Output)
		if err != nil {
			log.Fatal().Msgf("Failed to open logger file, %v.", err)
			return err
		}

		log.Logger = log.Output(f).Level(parseLevel(l.Level))
	}

	return nil
}

// parseLevel will parse the log level.
func parseLevel(l string) zerolog.Level {
	switch l {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	case "disable":
		return zerolog.Disabled
	}

	log.Panic().Msgf("Parse logger level failed.")
	return zerolog.Disabled
}
