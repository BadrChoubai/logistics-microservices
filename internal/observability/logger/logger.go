// Package logger
package logger

import (
	"io"
	"log/slog"

	"github.com/BadrChoubai/logistics-microservices/internal/config"
)

var _ StructuredLogger = (*Logger)(nil)

type StructuredLogger interface {
	Info(msg string, args ...any)
	Error(whatWasHappening string, args ...any)
	Warn(msg string, args ...any)
	Debug(msg string, args ...any)
}

type Logger struct {
	log *slog.Logger
}

func NewLogger(out io.Writer, level config.LogLevel) *Logger {
	handler := createHandler(out, level.SlogLevel())
	return &Logger{log: slog.New(handler)}
}

func createHandler(out io.Writer, level slog.Level) *slog.JSONHandler {
	return slog.NewJSONHandler(out, &slog.HandlerOptions{
		Level: level,
	})
}

func (l *Logger) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
}

func (l *Logger) Error(whatWasHappening string, args ...any) {
	l.log.Error(whatWasHappening, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.log.Warn(msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	l.log.Debug(msg, args...)
}
