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

	TimeKey    = slog.TimeKey
	SourceKey  = slog.SourceKey
	MessageKey = slog.MessageKey
	LevelKey   = slog.LevelKey
)

// Type aliases for slog types
type (
	// Attr is an alias for slog.Attr, representing a key-value pair
	// in structured logging
	Attr = slog.Attr

	Level = slog.Level
)

var (
	/*
		Attribute functions for common types
	*/

	Bool    = slog.Bool
	String  = slog.String
	Int     = slog.Int
	Int64   = slog.Int64
	Float64 = slog.Float64
	Any     = slog.Any
	Dict    = slog.Group

	NewLogLogger   = slog.NewLogLogger
	NewTextHandler = slog.NewTextHandler
	NewJSONHandler = slog.NewJSONHandler
	SetDefault     = slog.SetDefault
)
