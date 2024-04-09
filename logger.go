package f_logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ZeroLoggerOptions struct {
	Level    zerolog.Level
	Out      *os.File
	ErrorOut *os.File
}

func NewZeroLogger(options ZeroLoggerOptions) (*zerolog.Logger, error) {
	zerolog.SetGlobalLevel(options.Level)

	var output zerolog.ConsoleWriter
	if options.Out != nil {
		output = zerolog.ConsoleWriter{Out: options.Out, TimeFormat: "2006-01-02 15:04:05"}
	} else {
		output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	}

	logger := log.Output(output)

	return &logger, nil
}
