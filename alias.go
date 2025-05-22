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
)

// Type aliases for slog types
type (
	// Attr is an alias for slog.Attr, representing a key-value pair
	// in structured logging
	Attr = slog.Attr
)
