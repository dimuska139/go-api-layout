package logging

//go:generate mockgen -source=logger.go -destination=./logger_mock.go -package=logging

import (
	"fmt"
	"github.com/dimuska139/urlshortener/internal/config"
	"github.com/dimuska139/urlshortener/internal/constants"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/rs/zerolog"
	"os"
	"time"
)

type Loggerer interface {
	Debug(msg string, err error, extra map[string]interface{})
	Info(msg string, err error, extra map[string]interface{})
	Warn(msg string, err error, extra map[string]interface{})
	Error(msg string, err error, extra map[string]interface{})
	Panic(msg string, err error, extra map[string]interface{})
	Printf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Println(v ...interface{})
	Print(v ...interface{})
	NewPgxLogger() pgx.Logger
}

type Logger struct {
	logger zerolog.Logger
}

func NewLogger(cfg *config.Config) Loggerer {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Logger().
		Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	if cfg != nil && cfg.Loglevel != "" {
		switch cfg.Loglevel {
		case constants.LogLevelDebug:
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case constants.LogLevelInfo:
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case constants.LogLevelError:
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		default:
			zerolog.SetGlobalLevel(constants.DefaultLogLevel)
		}
	}
	return &Logger{logger}
}

func (z Logger) Debug(msg string, err error, extra map[string]interface{}) {
	z.logger.Debug().Timestamp().Err(err).Fields(extra).Msg(msg)
}

func (z Logger) Info(msg string, err error, extra map[string]interface{}) {
	z.logger.Info().Timestamp().Err(err).Fields(extra).Msg(msg)
}

func (z Logger) Warn(msg string, err error, extra map[string]interface{}) {
	z.logger.Warn().Timestamp().Err(err).Fields(extra).Msg(msg)
}

func (z Logger) Error(msg string, err error, extra map[string]interface{}) {
	z.logger.Error().Timestamp().Err(err).Fields(extra).Msg(msg)
}

func (z Logger) Panic(msg string, err error, extra map[string]interface{}) {
	z.logger.Panic().Timestamp().Err(err).Fields(extra).Msg(msg)
}

func (z Logger) Printf(format string, v ...interface{}) {
	z.logger.Printf(format, v...)
}

func (z Logger) Fatal(v ...interface{}) {
	z.logger.Fatal().Timestamp().Msg(fmt.Sprint(v...))
}

func (z Logger) Fatalf(format string, v ...interface{}) {
	z.logger.Fatal().Timestamp().Msgf(format, v...)
}

func (z Logger) Println(v ...interface{}) {
	z.logger.Info().Timestamp().Msgf("%v\r\n", v...)
}

func (z Logger) Print(v ...interface{}) {
	z.logger.Print(v...)
}

func (z Logger) NewPgxLogger() pgx.Logger {
	return zerologadapter.NewLogger(z.logger)
}
