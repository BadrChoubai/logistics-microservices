package config

import (
	"fmt"
	"strings"
	"time"
)

var _ ServiceConfig = (*Service)(nil)

type ServiceConfig interface {
	Validator
}

// Service is the configuration for API services: inventory, shipment, and telemetry.
// It embeds Base for all shared fields and adds a database connection.
type Service struct {
	Base

	// DB holds the database connection settings for this service.
	DB DatabaseConfig `json:"db"`
}

// DatabaseConfig holds PostgreSQL connection parameters.
type DatabaseConfig struct {
	// Host is the database server hostname or IP.
	Host string `json:"host"`

	// Port is the database server port (default 5432).
	Port int `json:"port"`

	// Name is the database/schema name.
	Name string `json:"name"`

	// User is the database login user.
	User string `json:"user"`

	// Password is the database login password.
	// Prefer injecting via env var DB_PASSWORD over committing to JSON.
	Password string `json:"password"`

	// SSLMode controls TLS: disable | require | verify-full
	SSLMode string `json:"ssl_mode"`

	// MaxOpenConns caps the connection pool size.
	MaxOpenConns int `json:"max_open_conns"`

	// MaxIdleConns caps idle connections in the pool.
	MaxIdleConns int `json:"max_idle_conns"`

	// ConnMaxLifetime is the maximum time a connection may be reused.
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`
}

// DSN returns a PostgreSQL connection string for database/sql.
func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		d.Host, d.Port, d.Name, d.User, d.Password, d.SSLMode,
	)
}

// Validate checks that the Service config is complete and sensible.
func (s *Service) Validate() error {
	if err := s.Base.validate(); err != nil {
		return fmt.Errorf("base: %w", err)
	}
	if err := s.DB.validate(); err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (d *DatabaseConfig) validate() error {
	if strings.TrimSpace(d.Host) == "" {
		return fmt.Errorf("host is required")
	}
	if d.Port < 1 || d.Port > 65535 {
		return fmt.Errorf("port must be 1–65535, got %d", d.Port)
	}
	if strings.TrimSpace(d.Name) == "" {
		return fmt.Errorf("name is required")
	}
	if strings.TrimSpace(d.User) == "" {
		return fmt.Errorf("user is required")
	}

	validSSL := map[string]bool{"disable": true, "require": true, "verify-full": true}
	if !validSSL[d.SSLMode] {
		return fmt.Errorf("ssl_mode must be disable | require | verify-full, got %q", d.SSLMode)
	}
	return nil
}
