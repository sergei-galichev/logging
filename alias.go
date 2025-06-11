// Package logging provides a simplified interface for structured logging
// using Go's standard log/slog package. It exposes common log levels
// and attribute types while maintaining compatibility with slog.
package logging

import (
	"log/slog"
)

// Log level constants matching slog's levels
const (
	// LevelDebug is the debug log level (lowest level)
	LevelDebug = slog.LevelDebug

	// LevelInfo is the info log level
	LevelInfo = slog.LevelInfo

	// LevelWarn is the warning log level
	LevelWarn = slog.LevelWarn

	// LevelError is the error log level (highest level)
	LevelError = slog.LevelError

	// TimeKey is the key used for timestamps in log records
	TimeKey = slog.TimeKey

	// SourceKey is the key used for source location information
	SourceKey = slog.SourceKey

	// MessageKey is the key used for log messages
	MessageKey = slog.MessageKey

	// LevelKey is the key used for log levels
	LevelKey = slog.LevelKey
)

// Level is an alias for slog.Level representing log level severity.
// It follows the same values as the level constants (Debug, Info, Warn, Error).
type (
	Level = slog.Level

	Logger = slog.Logger
)

// Commonly used attribute constructors that mirror slog's functions
var (
	// Bool creates a boolean-valued attribute
	Bool = slog.Bool

	// String creates a string-valued attribute
	String = slog.String

	// Int creates an integer-valued attribute
	Int = slog.Int

	// Int64 creates a 64-bit integer-valued attribute
	Int64 = slog.Int64

	// Time creates a time.Time-valued attribute
	Time = slog.Time

	// Duration creates a time.Duration-valued attribute
	Duration = slog.Duration

	// Float64 creates a 64-bit floating-point-valued attribute
	Float64 = slog.Float64

	// Any creates an attribute of any type
	Any = slog.Any

	// Dict creates a group of attributes (alias for slog.Group)
	Dict = slog.Group

	// Default is the default logger
	Default = slog.Default
)
