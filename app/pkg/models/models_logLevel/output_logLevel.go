package models_logLevel

import (
	"log/slog"
)

type LogLevel string

const (
	Undefined LogLevel = ""
	Error     LogLevel = "error"
	Warn      LogLevel = "warn"
	Info      LogLevel = "info"
	Debug     LogLevel = "debug"
)

func (l *LogLevel) String() string {
	return string(*l)
}

func (l *LogLevel) Set(v string) error {
	*l = LogLevel(v)
	return nil
}

func (l *LogLevel) Type() string {
	return "LogLevel"
}

func GetAllSupportedValues() string {
	return "[error|warn|info|debug]"
}

func SlogLevel(v LogLevel) slog.Level {
	switch v {
	case Error:
		return slog.LevelError
	case Warn:
		return slog.LevelWarn
	case Info:
		return slog.LevelInfo
	case Debug:
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}
