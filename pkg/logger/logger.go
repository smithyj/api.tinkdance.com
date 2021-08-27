package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logs = map[zerolog.Level]zerolog.Logger{}

func Debug() *zerolog.Event {
	l, ok := logs[zerolog.DebugLevel]
	if !ok {
		return log.Debug()
	}
	l = l.With().Timestamp().Caller().Stack().Logger()
	return l.Debug()
}

func Info() *zerolog.Event {
	l, ok := logs[zerolog.InfoLevel]
	if !ok {
		return log.Debug()
	}
	l = l.With().Timestamp().Logger()
	return l.Info()
}

func Warn() *zerolog.Event {
	l, ok := logs[zerolog.WarnLevel]
	if !ok {
		return log.Debug()
	}
	l = l.With().Timestamp().Caller().Logger()
	return l.Warn()
}

func Error() *zerolog.Event {
	l, ok := logs[zerolog.ErrorLevel]
	if !ok {
		return log.Debug()
	}
	l = l.With().Timestamp().Caller().Stack().Logger()
	return l.Error()
}

func Fatal() *zerolog.Event {
	l, ok := logs[zerolog.FatalLevel]
	if !ok {
		return log.Debug()
	}
	l = l.With().Timestamp().Caller().Stack().Logger()
	return l.Fatal()
}

func Panic() *zerolog.Event {
	l, ok := logs[zerolog.PanicLevel]
	if !ok {
		return log.Debug()
	}
	l = l.With().Timestamp().Caller().Stack().Logger()
	return l.Panic()
}

func Trace() *zerolog.Event {
	l, ok := logs[zerolog.TraceLevel]
	if !ok {
		return log.Debug()
	}
	l = l.With().Timestamp().Logger()
	return l.Trace()
}

func Setup(config Config) {
	zerolog.TimestampFieldName = "time"
	if config.Path == "" {
		config.Path = "./logs"
	}
	levels := []zerolog.Level{
		zerolog.DebugLevel,
		zerolog.InfoLevel,
		zerolog.WarnLevel,
		zerolog.ErrorLevel,
		zerolog.FatalLevel,
		zerolog.PanicLevel,
		zerolog.TraceLevel,
	}

	for _, level := range levels {
		if !config.Debug && level == zerolog.DebugLevel {
			level = zerolog.Disabled
		}
		hook := lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s.log", config.Path, zerolog.LevelFieldMarshalFunc(level)),
			MaxSize:    config.MaxSize,
			MaxAge:     config.MaxAge,
			MaxBackups: config.MaxBackups,
			Compress:   config.Compress,
		}
		writer := zerolog.MultiLevelWriter(
			os.Stdout,
			&hook,
		)
		logs[level] = zerolog.New(writer)
	}
}
