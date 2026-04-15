package config

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

// Validator is implemented by config types that can validate themselves.
type Validator interface {
	Validate() error
}

// Environment represents the application's environment
type Environment string

// LogLevel represents the structured log verbosity for a service.
type LogLevel string

type Base struct {
	Environment Environment `json:"environment"`
	LogLevel    LogLevel    `json:"log_level"`
	Port        int         `json:"http_port"`
}

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"

	EnvironmentDevelopment Environment = "development"
	EnvironmentStaging     Environment = "staging"
	EnvironmentProduction  Environment = "production"
)

// Load reads cfg from a JSON file at path. Finally it calls Validate() if cfg
// implements Validator.
//
// If the file does not exist, Load returns an error.
func Load[T Validator](path string) (T, error) {
	var cfg T

	if path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			return cfg, fmt.Errorf("config: read %q: %w", path, err)
		}
		if err := json.Unmarshal(data, &cfg); err != nil {
			return cfg, fmt.Errorf("config: parse %q: %w", path, err)
		}
	}

	if err := cfg.Validate(); err != nil {
		return cfg, fmt.Errorf("config: validation failed: %w", err)
	}

	return cfg, nil
}

func (e *Environment) UnmarshalJSON(data []byte) error {
	s := Environment(strings.ToLower(strings.Trim(string(data), `"`)))
	switch s {
	case EnvironmentDevelopment, EnvironmentStaging, EnvironmentProduction:
		*e = s
		return nil
	default:
		return fmt.Errorf("invalid app_env %q: must be development | staging | production", s)
	}
}

func (l *LogLevel) UnmarshalJSON(data []byte) error {
	s := LogLevel(strings.ToLower(strings.Trim(string(data), `"`)))
	switch s {
	case LogLevelDebug, LogLevelInfo, LogLevelWarn, LogLevelError:
		*l = s
		return nil
	case "warning":
		*l = LogLevelWarn
		return nil
	default:
		return fmt.Errorf("invalid log_level %q: must be debug | info | warn | error", s)
	}
}

func (l LogLevel) SlogLevel() slog.Level {
	switch l {
	case LogLevelDebug:
		return slog.LevelDebug
	case LogLevelWarn:
		return slog.LevelWarn
	case LogLevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// validate checks Base fields for obvious misconfiguration.
func (b *Base) validate() error {
	if b.Port < 1 || b.Port > 65535 {
		return fmt.Errorf("port must be 1–65535, got %d", b.Port)
	}

	return nil
}
